package postgres

import (
	"errors"
	"fmt"
	"test/api/models"
	"test/pkg/helper"
	"test/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
)

type incomeRepo struct {
	db *pgxpool.Pool
}

func NewIncomeRepo(db *pgxpool.Pool) storage.IIncomeStorage {
	return &incomeRepo{
		db: db,
	}
}

func (i *incomeRepo) Create(ctx context.Context) (models.Income, error) {
	var (
		income = models.Income{}
		extID  string
	)

	query := `select external_id from incomes order by external_id desc`

	if err := i.db.QueryRow(ctx, query).Scan(
		&extID,
	); err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error while getting ext id ", err.Error())
			return models.Income{}, err
		}
		extID = "I"
	}

	if extID != "I" {
		extID = helper.GenerateExternalID(extID)
	} else {
		extID = "I-0001"
	}

	fmt.Println("ex", extID)

	query = `insert into incomes values ($1, $2, $3) returning id, external_id`

	fmt.Println("ext id ", extID)
	if err := i.db.QueryRow(ctx, query, uuid.New(), extID, 0).Scan(
		&income.ID,
		&income.ExternalID,
	); err != nil {
		fmt.Println("error while creating income ", err.Error())
		return models.Income{}, err
	}

	return income, nil
}
