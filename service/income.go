package service

import (
	"context"
	"fmt"
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
