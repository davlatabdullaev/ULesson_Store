package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"test/api/models"
	"test/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) storage.IProductStorage {
	return &productRepo{db: db}
}

func (p *productRepo) Create(ctx context.Context, product models.CreateProduct) (string, error) {
	id := uuid.New()
	query := `insert into products(id, name, price, original_price, quantity, category_id, branch_id) 
						values($1, $2, $3, $4, $5, $6, $7)`

	if rowsAffected, err := p.db.Exec(ctx, query,
		id,
		product.Name,
		product.Price,
		product.OriginalPrice,
		product.Quantity,
		product.CategoryID,
		product.BranchID); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			fmt.Println("error is in rows affected", err.Error())
			return "", err
		}
		fmt.Println("error is while inserting product", err.Error())
		return "", err
	}

	return id.String(), nil
}

func (p *productRepo) GetByID(ctx context.Context, key models.PrimaryKey) (models.Product, error) {
	var createdAt, updatedAt = sql.NullString{}, sql.NullString{}
	product := models.Product{}
	query := `select id, name, price, original_price, quantity, category_id, branch_id, created_at, updated_at
							from products where id = $1 and deleted_at = 0`
	if err := p.db.QueryRow(ctx, query, key.ID).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.OriginalPrice,
		&product.Quantity,
		&product.CategoryID,
		&product.BranchID,
		&createdAt,
		&updatedAt); err != nil {
		fmt.Println("error is while selecting product by id", err.Error())
		return models.Product{}, err
	}

	if createdAt.Valid {
		product.CreatedAt = createdAt.String
	}

	if updatedAt.Valid {
		product.UpdatedAt = updatedAt.String
	}
	return product, nil
}

func (p *productRepo) GetList(ctx context.Context, request models.GetListRequest) (models.ProductResponse, error) {
	var (
		products             = []models.Product{}
		page                 = request.Page
		offset               = (page - 1) * request.Limit
		search               = request.Search
		query, countQuery    string
		count                = 0
		createdAt, updatedAt = sql.NullString{}, sql.NullString{}
	)

	countQuery = `select count(1) from products where deleted_at = 0 `

	if search != "" {
		countQuery += fmt.Sprintf(` and (name ilike '%%%s%%' or 
			CAST(price AS TEXT) ilike '%s' or CAST(quantity AS TEXT) ilike '%s')`, search, search, search)
	}

	if err := p.db.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while scanning count", err.Error())
		return models.ProductResponse{}, err
	}

	query = `select id, name, price, original_price, quantity, category_id, branch_id, created_at, updated_at
								from products where deleted_at = 0`

	if search != "" {
		query += fmt.Sprintf(` and (name ilike '%%%s%%' or 
			CAST(price AS TEXT) ilike '%s' or CAST(quantity AS TEXT) ilike '%s')`, search, search, search)
	}

	query += ` order by created_at desc LIMIT $1 OFFSET $2`

	rows, err := p.db.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting products", err.Error())
		return models.ProductResponse{}, err
	}

	for rows.Next() {
		product := models.Product{}
		if err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.OriginalPrice,
			&product.Quantity,
			&product.CategoryID,
			&product.BranchID,
			&createdAt,
			&updatedAt); err != nil {
			fmt.Println("error is while scanning products", err.Error())
			return models.ProductResponse{}, err
		}
		if createdAt.Valid {
			product.CreatedAt = createdAt.String
		}

		if updatedAt.Valid {
			product.UpdatedAt = updatedAt.String
		}
		products = append(products, product)
	}
	return models.ProductResponse{
		Product: products,
		Count:   count,
	}, err
}

func (p *productRepo) Update(ctx context.Context, product models.UpdateProduct) (string, error) {
	query := `update products set name = $1, price = $2, original_price = $3, quantity = $4, 
                    category_id = $5, updated_at = now()  where id = $6`

	if _, err := p.db.Exec(ctx, query,
		&product.Name,
		&product.Price,
		&product.OriginalPrice,
		&product.Quantity,
		&product.CategoryID,
		&product.ID); err != nil {
		fmt.Println("error is while updating product", err.Error())
		return "", err
	}

	return product.ID, nil
}

func (p *productRepo) Delete(ctx context.Context, key models.PrimaryKey) error {
	query := `update products set deleted_at = extract(epoch from current_timestamp) where id = $1`

	if rowsAffected, err := p.db.Exec(ctx, query, key.ID); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			fmt.Println("error is in rows affected", err.Error())
			return err
		}
		fmt.Println("error is while deleting product", err.Error())
		return err
	}
	return nil
}

func (p *productRepo) Search(ctx context.Context, customerProductIDs map[string]int) (models.ProductSell, error) {
	var (
		selectedProducts = models.SellRequest{
			Products: map[string]int{},
		}
		products               = make([]string, len(customerProductIDs))
		selectedProductPrices  = make(map[string]int, 0)
		notEnoughProducts      = make(map[string]int)
		productsBranchID       string
		notEnoughProductPrices = make(map[string]int)
	)

	for key := range customerProductIDs {
		products = append(products, key)
	}

	query := `
				select id, quantity, price, original_price, branch_id from products where id::varchar = ANY($1)
	`

	rows, err := p.db.Query(ctx, query, pq.Array(products)) // [a, b, c]
	if err != nil {
		fmt.Println("Error while getting products by product ids", err.Error())
		return models.ProductSell{}, err
	}

	for rows.Next() {
		var (
			quantity, price, originalPrice int
			productID, branchID            string
		)
		if err = rows.Scan(
			&productID,
			&quantity,
			&price,
			&originalPrice,
			&branchID,
		); err != nil {
			fmt.Println("Error while scanning rows one by one", err.Error())
			return models.ProductSell{}, err
		}

		productsBranchID = branchID

		if customerProductIDs[productID] <= quantity {
			selectedProducts.Products[productID] = price
			selectedProductPrices[productID] = originalPrice
		} else if customerProductIDs[productID] > quantity || quantity == 0 {
			notEnoughProducts[productID] = customerProductIDs[productID]
			notEnoughProductPrices[productID] = originalPrice
		}
	}

	return models.ProductSell{
		SelectedProducts:       selectedProducts,
		ProductPrices:          selectedProductPrices,
		NotEnoughProducts:      notEnoughProducts,
		NotEnoughProductPrices: notEnoughProductPrices,
		ProductsBranchID:       productsBranchID,
	}, nil
}

func (p *productRepo) TakeProducts(ctx context.Context, products map[string]int) error {
	var (
		updateStatements []string
	)
	query := `
	DO $$
	BEGIN
		%s
	END $$
`

	for productID, quantity := range products {
		updateStatements = append(updateStatements, fmt.Sprintf(`update products 
			set quantity = quantity - %d where id = '%s' ;`, quantity, productID))
	}

	finalQuery := fmt.Sprintf(query, strings.Join(updateStatements, "\n"))

	if _, err := p.db.Exec(ctx, finalQuery); err != nil {
		fmt.Println("Error while updating product quantity", err.Error())
		return err
	}
	return nil
}

func (p *productRepo) AddDeliveredProducts(ctx context.Context, products models.DeliverProducts, branchID string) error {

	var (
		updatedStatements []string
	)

	query := `
					DO $$
					BEGIN
                       %s
					END $$
`
	if products.NotEnoughProducts != nil {
		for productID, quantity := range products.NotEnoughProducts {
			updatedStatements = append(updatedStatements, fmt.Sprintf(`update products set quantity = quantity + %d where id = '%s' ;`,
				quantity, productID))
		}

		finalQuery := fmt.Sprintf(query, strings.Join(updatedStatements, "\n"))

		if rowsAffected, err := p.db.Exec(ctx, finalQuery); err != nil {
			if r := rowsAffected.RowsAffected(); r == 0 {
				fmt.Println("error is while rows affected", err.Error())
				return err
			}
			fmt.Println("error is while updating quantity of delivered products", err.Error())
			return err
		}
	}

	return nil
}
