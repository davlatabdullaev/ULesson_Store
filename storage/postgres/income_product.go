package postgres

import (
	"context"
	"fmt"
	"log"
	"strings"
	"test/api/models"
	"test/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type incomeProductRepo struct {
	db *pgxpool.Pool
}

func NewIncomeProductRepo(db *pgxpool.Pool) storage.IIncomeProductStorage {
	return &incomeProductRepo{db: db}
}

func (i *incomeProductRepo) CreateMultiple(ctx context.Context, request models.CreateIncomeProducts) error {
	query := `insert into income_products (id, income_id, product_id, quantity, price) values `

	for _, incomeProduct := range request.IncomeProducts {
		query += fmt.Sprintf(`('%s', '%s', '%s', %d, %d), `, uuid.New().String(), incomeProduct.IncomeID, incomeProduct.ProductID, incomeProduct.Quantity, incomeProduct.Price)
	}
	query = query[:len(query)-2]

	if _, err := i.db.Exec(ctx, query); err != nil {
		fmt.Println("error while inserting income products ", err.Error())
		return err
	}

	return nil
}

func (i *incomeProductRepo) GetList(ctx context.Context, request models.GetListRequest) (models.IncomeProductsResponse, error) {
	var (
		page              = request.Page
		offset            = (page - 1) * request.Limit
		count             = 0
		query, countQuery string
		incomeProducts    = []models.IncomeProduct{}
	)

	countQuery = `select count(1) from income_products where deleted_at = 0`
	if request.Search != "" {
		countQuery += fmt.Sprintf(` and income_id = '%s'`, request.Search)

	}
	if err := i.db.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while scanning count from income products", err.Error())
		return models.IncomeProductsResponse{}, err
	}

	query = ` select id, income_id, product_id, quantity, price from income_products where deleted_at = 0`
	if request.Search != "" {
		query += fmt.Sprintf(` and income_id = '%s'`, request.Search)
	}

	query += ` limit $1 offset $2`
	rows, err := i.db.Query(ctx, query, request.Limit, offset)
	if err != nil {
		log.Println("error is while selecting all from income products", err.Error())
		return models.IncomeProductsResponse{}, err
	}

	for rows.Next() {
		incomeProduct := models.IncomeProduct{}
		if err = rows.Scan(
			&incomeProduct.ID,
			&incomeProduct.IncomeID,
			&incomeProduct.ProductID,
			&incomeProduct.Quantity,
			&incomeProduct.Price,
		); err != nil {
			log.Println("error is while scanning all from income products", err.Error())
			return models.IncomeProductsResponse{}, err
		}
		incomeProducts = append(incomeProducts, incomeProduct)
	}

	return models.IncomeProductsResponse{
		IncomeProducts: incomeProducts,
		Count:          count,
	}, err

}

func (i *incomeProductRepo) UpdateMultiple(ctx context.Context, request models.IncomeProducts) error {
	var (
		updateStatements []string
	)
	query := `do $$ begin %s end $$`
	for _, incomeProducts := range request.IncomeProducts {
		updateStatements = append(updateStatements, fmt.Sprintf(`update income_products set income_id = '%s', product_id = '%s', quantity = %d, price = %d where id = '%s' ;`,
			incomeProducts.IncomeID,
			incomeProducts.ProductID,
			incomeProducts.Quantity,
			incomeProducts.Price,
			incomeProducts.ID,
		))
	}

	finalQuery := fmt.Sprintf(query, strings.Join(updateStatements, "\n"))
	if rowsAffected, err := i.db.Exec(ctx, finalQuery); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			log.Println("error is while rows affected", err.Error())
			return err
		}
		log.Println("error is while updating income products", err.Error())
		return err
	}
	return nil
}

func (i *incomeProductRepo) DeleteMultiple(ctx context.Context, request models.DeleteIncomeProducts) error {
	var deleteStatements []string

	query := `do $$ begin %s end $$`
	for _, value := range request.IDs {
		deleteStatements = append(deleteStatements, fmt.Sprintf(`update income_products set deleted_at = extract(epoch from current_timestamp) where id = '%s'`, value.ID))
	}

	finalQuery := fmt.Sprintf(query, strings.Join(deleteStatements, "\n"))
	if rowsAffected, err := i.db.Exec(ctx, finalQuery); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			log.Println("error while rows affected", err.Error())
			return err
		}
		log.Println("error is while deleting income products", err.Error())
		return err
	}

	return nil

}
