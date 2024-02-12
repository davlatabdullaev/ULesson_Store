package service

import (
	"context"
	"fmt"
	"log"
	"test/api/models"
	"test/storage"
)

type incomeService struct {
	storage storage.IStorage
}

func NewIncomeService(storage storage.IStorage) incomeService {
	return incomeService{
		storage: storage,
	}
}

func (i incomeService) Create(ctx context.Context) (models.Income, error) {
	income, err := i.storage.Income().Create(ctx)
	if err != nil {
		fmt.Println("error while creating income ", err.Error())
		return models.Income{}, err
	}

	return income, nil
}

func (i incomeService) Get(ctx context.Context, key models.PrimaryKey) (models.Income, error) {

	income, err := i.storage.Income().GetByID(ctx, key)
	if err != nil {
		log.Println("error in service layer while getting by id", err.Error())
		return models.Income{}, err
	}
	return income, nil

}

func (i incomeService) GetList(ctx context.Context, request models.GetListRequest) (models.IncomesResponse, error) {
	incomes, err := i.storage.Income().GetList(ctx, request)
	if err != nil {
		log.Println("error in service layer getting list", err.Error())
		return models.IncomesResponse{}, err
	}
	return incomes, nil
}

func (i incomeService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := i.storage.Income().Delete(ctx, key)
	return err
}
