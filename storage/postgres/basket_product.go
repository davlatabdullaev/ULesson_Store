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
)

type basketProductRepo struct {
	db *pgxpool.Pool
}

func NewBasketProductRepo(db *pgxpool.Pool) storage.IBasketProductStorage {
	return &basketProductRepo{db: db}
}

func (b *basketProductRepo) Create(ctx context.Context, product models.CreateBasketProduct) (string, error) {
	id := uuid.New()
	query := `insert into basket_products(id, basket_id, product_id, quantity) 
					values($1, $2, $3, $4)`

	fmt.Println("id", id)
	if _, err := b.db.Exec(ctx, query,
		id,
		product.BasketID,
		product.ProductID,
		product.Quantity); err != nil {
		fmt.Println("error is while insert", err.Error())
		return "", err
	}
	return id.String(), nil
}

func (b *basketProductRepo) GetByID(ctx context.Context, key models.PrimaryKey) (models.BasketProduct, error) {
	var createdAt, updatedAt = sql.NullString{}, sql.NullString{}
	product := models.BasketProduct{}
	query := `select id, basket_id, product_id, quantity, created_at, updated_at from basket_products where id = $1 and deleted_at = 0`

	if err := b.db.QueryRow(ctx, query, key.ID).Scan(
		&product.ID,
		&product.BasketID,
		&product.ProductID,
		&product.Quantity,
		&createdAt,
		&updatedAt,
	); err != nil {
		fmt.Println("error is while selecting by id", err.Error())
		return models.BasketProduct{}, err
	}

	if createdAt.Valid {
		product.CreatedAt = createdAt.String
	}

	if updatedAt.Valid {
		product.UpdatedAt = updatedAt.String
	}

	return product, nil
}

func (b *basketProductRepo) GetList(ctx context.Context, request models.GetListRequest) (models.BasketProductResponse, error) {
	var (
		count                = 0
		basketProducts       = []models.BasketProduct{}
		query, countQuery    string
		page                 = request.Page
		offset               = (page - 1) * request.Limit
		search               = request.Search
		createdAt, updatedAt = sql.NullString{}, sql.NullString{}
	)

	countQuery = `select count(1) from basket_products where deleted_at = 0 `
	if search != "" {
		countQuery += fmt.Sprintf(` and CAST(quantity AS TEXT) = '%s'`, search)
	}

	if request.BasketID != "" {
		countQuery += fmt.Sprintf(" and basket_id = '%s'", request.BasketID)
	}

	if err := b.db.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while scanning count", err.Error())
		return models.BasketProductResponse{}, err
	}

	query = `select id, basket_id, product_id, quantity, created_at, updated_at from basket_products where deleted_at = 0`
	if search != "" {
		query += fmt.Sprintf(` and CAST(quantity AS TEXT) = '%s'`, search)
	}

	if request.BasketID != "" {
		query += fmt.Sprintf(" and basket_id = '%s'", request.BasketID)
	}

	query += ` order by created_at desc LIMIT $1 OFFSET $2`

	rows, err := b.db.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting basket products", err.Error())
		return models.BasketProductResponse{}, err
	}

	for rows.Next() {
		basketProd := models.BasketProduct{}
		if err = rows.Scan(
			&basketProd.ID,
			&basketProd.BasketID,
			&basketProd.ProductID,
			&basketProd.Quantity,
			&createdAt,
			&updatedAt,
		); err != nil {
			fmt.Println("error is while scanning basket products", err.Error())
			return models.BasketProductResponse{}, err
		}
		if createdAt.Valid {
			basketProd.CreatedAt = createdAt.String
		}

		if updatedAt.Valid {
			basketProd.UpdatedAt = updatedAt.String
		}
		basketProducts = append(basketProducts, basketProd)
	}

	return models.BasketProductResponse{
		BasketProducts: basketProducts,
		Count:          count,
	}, err
}

func (b *basketProductRepo) Update(ctx context.Context, product models.UpdateBasketProduct) (string, error) {
	query := `update basket_products set product_id = $1, quantity = $2, updated_at = now() where id = $3 `
	if rowsAffected, err := b.db.Exec(ctx, query,
		&product.ProductID,
		&product.Quantity,
		&product.ID); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			fmt.Println("error is in rows affected", err.Error())
			return "", err
		}
		fmt.Println("error is while updating basket_products", err.Error())
		return "", err
	}

	return product.ID, nil
}

func (b *basketProductRepo) Delete(ctx context.Context, key models.PrimaryKey) error {
	query := `update basket_products set deleted_at = extract(epoch from current_timestamp) where id = $1`

	if rowsAffected, err := b.db.Exec(ctx, query, key.ID); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			fmt.Println("error is in rows affected", err.Error())
			return err
		}
		fmt.Println("error is while deleting basket products", err.Error())
		return err
	}
	return nil
}

func (b *basketProductRepo) AddProducts(ctx context.Context, basketID string, products map[string]int) error {
	var (
		insertStatements []string
	)
	query := `
		DO $$
		BEGIN 
           %s
		END $$
`
	for productID, quantity := range products {
		insertStatements = append(insertStatements, fmt.Sprintf(`insert into basket_products (id, basket_id, product_id, quantity)
                      values ('%s', '%s', '%s', %d) ;`, uuid.New(), basketID, productID, quantity))
	}

	finalQuery := fmt.Sprintf(query, strings.Join(insertStatements, "\n"))

	if _, err := b.db.Exec(ctx, finalQuery); err != nil {
		fmt.Println("error is while inserting to basket products", err.Error())
		return err
	}

	return nil
}

//func (b basketProductRepo) AddProducts(basketID string, products map[string]int) error {
//	query := `
//			insert into basket_products
//			    (id, basket_id, product_id, quantity)
//					values ($1, $2, $3, $4)
//`
//
//	for productID, quantity := range products {
//		if _, err := b.db.Exec(context.Background(), query, uuid.New(), basketID, productID, quantity); err != nil {
//			fmt.Println("Error while adding product to basket_products table", err.Error())
//			return err
//		}
//	}
//
//	return nil
//}
