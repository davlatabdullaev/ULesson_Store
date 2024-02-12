package service

import (
	"context"
	"fmt"
	"log"
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

func (i incomeProductService) GetList(ctx context.Context, request models.GetListRequest) (models.IncomeProductsResponse, error) {

	incomeProducts, err := i.storage.IncomeProduct().GetList(ctx, request)
	if err != nil {
		log.Println("error in service layer getting list", err.Error())
		return models.IncomeProductsResponse{}, err
	}
	return incomeProducts, nil

}

func (i incomeProductService) UpdateMultiple(ctx context.Context, request models.IncomeProducts) error {

	err := i.storage.IncomeProduct().UpdateMultiple(ctx, request)
	if err != nil {
		log.Println("error in service layer update multiple", err.Error())
		return err
	}
	return nil

}

func (i incomeProductService) DeleteMultiple(ctx context.Context, request models.DeleteIncomeProducts) error {

	err := i.storage.IncomeProduct().DeleteMultiple(ctx, request)

	return err
}
