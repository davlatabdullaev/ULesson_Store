package postgres

import (
	"context"
	"fmt"
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
