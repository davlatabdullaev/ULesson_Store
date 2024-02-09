package service

import (
	"context"
	"fmt"
	"test/api/models"
	"test/storage"
)

type incomeProductService struct {
	storage storage.IStorage
}

func NewIncomeProductService(storage storage.IStorage) incomeProductService {
	return incomeProductService{
		storage: storage,
	}
}

func (i incomeProductService) CreateMultiple(ctx context.Context, request models.CreateIncomeProducts) error {
	if err := i.storage.IncomeProduct().CreateMultiple(ctx, request); err != nil {
		fmt.Println("error while creating multiple income products", err.Error())
		return err
	}

	return nil
}
