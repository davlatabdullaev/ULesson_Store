package postgres

import (
	"context"
	"fmt"
	"test/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type dealerRepo struct {
	db *pgxpool.Pool
}

func NewDealerRepo(db *pgxpool.Pool) storage.IDealerStorage {
	return &dealerRepo{
		db: db,
	}
}

func (d *dealerRepo) AddSum(ctx context.Context, totalSum int) error {
	//ozini sum: ga qoshish kerak total sum -> update
	query := `update dealer set sum = sum + $1 where id = '1cfd84e6-72cb-4135-a802-85d10e4183ea'`
	if rowsAffected, err := d.db.Exec(ctx, query, &totalSum); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			fmt.Println("error is while rows affected", err.Error())
			return err
		}
		fmt.Println("error is while updating dealer sum", err.Error())
		return err
	}
	return nil
}
