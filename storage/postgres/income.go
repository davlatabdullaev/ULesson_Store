package postgres

import (
	"errors"
	"fmt"
	"log"
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

	query = `insert into incomes (id, external_id, total_sum, created_at) values ($1, $2, $3, now()) returning id, external_id`

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

func (i *incomeRepo) GetByID(ctx context.Context, key models.PrimaryKey) (models.Income, error) {

	income := models.Income{}

	query := `select id, external_id, total_sum from incomes where id = $1`

	if err := i.db.QueryRow(ctx, query, income.ID).Scan(
		&income.ID,
		&income.ExternalID,
		&income.TotalSum,
	); err != nil {
		log.Println("error while getting income by id")

		return models.Income{}, err

	}

	return income, nil
}

func (i *incomeRepo) GetList(ctx context.Context, request models.GetListRequest) (models.IncomesResponse, error) {

	var (
		query, countQuery string
		count             = 0
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
		incomes           = []models.Income{}
	)

	countQuery = `select count(1) from incomes `

	if search != "" {
		countQuery += fmt.Sprintf(` and external_id ilike '%%%s%%'`, search)
	}

	if err := i.db.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		log.Println("error is while scanning count", err.Error())
		return models.IncomesResponse{}, err
	}

	query = `select id, external_id, total_sum from incomes `

	if search != "" {
		query += fmt.Sprintf(` and external_id ilike '%%%s%%'`, search)
	}

	query += ` LIMIT $1 OFFSET $2 `

	rows, err := i.db.Query(ctx, query, request.Limit, offset)
	if err != nil {
		log.Println("error while selecting incomes ", err.Error())
		return models.IncomesResponse{}, err
	}

	for rows.Next() {
		income := models.Income{}
		if err = rows.Scan(
			&income.ID,
			&income.ExternalID,
			&income.TotalSum,
		); err != nil {
			log.Println("error while scanning income data...", err.Error())
			return models.IncomesResponse{}, err
		}

		incomes = append(incomes, income)
	}

	return models.IncomesResponse{
		Incomes: incomes,
		Count:   count,
	}, nil
}

func (i *incomeRepo) Delete(ctx context.Context, key models.PrimaryKey) error {

	query := `update incomes set deleted_at = extract(epoch from current_timestamp) where id = $1`

	rowAffected, err := i.db.Exec(ctx, query, key.ID)

	r := rowAffected.RowsAffected()

	if r == 0 {
		log.Println("error while rows affected")
	}

	if err != nil {
		log.Println("error while delteing incomes")
	}

	return nil
}
