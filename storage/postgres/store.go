package postgres

import (
	"context"
	"fmt"
	"test/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type storeRepo struct {
	db *pgxpool.Pool
}

func NewStoreRepo(db *pgxpool.Pool) storage.IStoreStorage {
	return &storeRepo{
		db: db,
	}
}

func (s *storeRepo) AddProfit(ctx context.Context, profit float32, branchID string) error {
	rowsAffected, err := s.db.Exec(ctx, `update store set profit = profit + $1, updated_at = now() where branch_id = $2`, profit, branchID)
	if err != nil {
		fmt.Println("Error while adding profit to store", err.Error())
		return err
	}

	if n := rowsAffected.RowsAffected(); n == 0 {
		fmt.Println("Error in rows affected", err.Error())
		return err
	}

	return err
}

func (s *storeRepo) GetStoreBudget(ctx context.Context, branchID string) (float32, error) {
	var budget float32
	query := `select budget from store where branch_id = $1`
	if err := s.db.QueryRow(ctx, query, branchID).Scan(&budget); err != nil {
		fmt.Println("error is while getting store budget", err.Error())
		return 0, err
	}

	return budget, nil
}

func (s *storeRepo) WithdrawalDeliveredSum(ctx context.Context, totalSum float32, branchID string) error {
	query := `update store set budget = budget - $1 where branch_id = $2 `
	if rowsAffected, err := s.db.Exec(ctx, query, &totalSum, &branchID); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			fmt.Println("error is while rows affected", err.Error())
			return err
		}
		fmt.Println("error is while updating budget", err.Error())
		return err
	}

	return nil
}
