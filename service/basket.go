package service

import (
	"context"
	"fmt"
	"test/api/models"
	"test/storage"
)

type basketService struct {
	storage storage.IStorage
}

func NewBasketService(storage storage.IStorage) basketService {
	return basketService{storage: storage}
}

func (b basketService) Create(ctx context.Context, basket models.CreateBasket) (models.Basket, error) {
	id, err := b.storage.Basket().Create(ctx, basket)
	if err != nil {
		fmt.Println("error in service layer while creating basket", err.Error())
		return models.Basket{}, err
	}

	createdBasket, err := b.storage.Basket().GetByID(context.Background(), models.PrimaryKey{ID: id})
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		return models.Basket{}, err
	}

	return createdBasket, err
}

func (b basketService) Get(ctx context.Context, id string) (models.Basket, error) {
	basket, err := b.storage.Basket().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		fmt.Println("error in service layer while getting by id", err.Error())
		return models.Basket{}, err
	}

	return basket, nil
}

func (b basketService) GetList(ctx context.Context, request models.GetListRequest) (models.BasketResponse, error) {
	baskets, err := b.storage.Basket().GetList(ctx, request)
	if err != nil {
		fmt.Println("error in service layer  while getting list", err.Error())
		return models.BasketResponse{}, err
	}

	return baskets, nil
}

func (b basketService) Update(ctx context.Context, basket models.UpdateBasket) (models.Basket, error) {
	id, err := b.storage.Basket().Update(ctx, basket)
	if err != nil {
		fmt.Println("error in service layer while updating", err.Error())
		return models.Basket{}, err
	}

	updatedBasket, err := b.storage.Basket().GetByID(context.Background(), models.PrimaryKey{ID: id})
	if err != nil {
		fmt.Println("error in service layer while getting basket by id", err.Error())
		return models.Basket{}, err
	}

	return updatedBasket, nil
}

func (b basketService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := b.storage.Basket().Delete(ctx, key)

	return err
}
