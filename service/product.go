package service

import (
	"context"
	"fmt"
	"test/api/models"
	"test/storage"
)

type productService struct {
	storage storage.IStorage
}

func NewProductService(storage storage.IStorage) productService {
	return productService{storage: storage}
}

func (p productService) Create(ctx context.Context, product models.CreateProduct) (models.Product, error) {
	id, err := p.storage.Product().Create(ctx, product)
	if err != nil {
		fmt.Println("error in service layer while creating product", err.Error())
		return models.Product{}, err
	}

	createdProduct, err := p.storage.Product().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		fmt.Println("error in service layer while getting by id", err.Error())
		return models.Product{}, err
	}

	return createdProduct, nil
}

func (p productService) Get(ctx context.Context, key models.PrimaryKey) (models.Product, error) {
	product, err := p.storage.Product().GetByID(ctx, key)
	if err != nil {
		fmt.Println("error in service layer while getting by id", err.Error())
		return models.Product{}, err
	}

	return product, nil
}

func (p productService) GetList(ctx context.Context, request models.GetListRequest) (models.ProductResponse, error) {
	products, err := p.storage.Product().GetList(ctx, request)
	if err != nil {
		fmt.Println("error in service layer while getting list", err.Error())
		return models.ProductResponse{}, err
	}

	return products, nil
}

func (p productService) Update(ctx context.Context, product models.UpdateProduct) (models.Product, error) {
	id, err := p.storage.Product().Update(ctx, product)
	if err != nil {
		fmt.Println("error in service layer while update", err.Error())
		return models.Product{}, err
	}

	updatedProduct, err := p.storage.Product().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		fmt.Println("error in service layer while getting by id", err.Error())
		return models.Product{}, err
	}

	return updatedProduct, nil
}

func (p productService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := p.storage.Product().Delete(ctx, key)

	return err
}

func (p productService) StartSellNew(ctx context.Context, request models.SellRequest) (models.ProductSell, error) {
	productSell, err := p.storage.Product().Search(ctx, request.Products)
	if err != nil {
		fmt.Println("error in service layer while searching product", err.Error())
		return models.ProductSell{}, err
	}

	basket, err := p.storage.Basket().GetByID(ctx, models.PrimaryKey{ID: request.BasketID})
	if err != nil {
		fmt.Println("error in service layer while getting basket by id", err.Error())
		return models.ProductSell{}, err
	}

	customer, err := p.storage.User().GetByID(ctx, models.PrimaryKey{ID: basket.CustomerID})
	if err != nil {
		fmt.Println("error in service layer while getting user by id", err.Error())
		return models.ProductSell{}, err
	}

	totalSum, profit := 0, float32(0.0)
	basketProducts := map[string]int{}

	for productID, price := range productSell.SelectedProducts.Products {
		customerQuantity := request.Products[productID]
		totalSum += price * customerQuantity

		//profit logic
		profit += float32(customerQuantity*price - productSell.ProductPrices[productID])
		basketProducts[productID] = customerQuantity
	}

	if customer.Cash < uint(totalSum) {
		fmt.Println("error in service layer while not enough customer cash", err.Error())
		return models.ProductSell{}, err
	}

	if err = p.storage.User().UpdateCustomerCash(ctx, customer.ID, totalSum); err != nil {
		fmt.Println("error in service layer while updating customer cash", err.Error())
		return models.ProductSell{}, err
	}

	if err = p.storage.Product().TakeProducts(ctx, basketProducts); err != nil {
		fmt.Println("error in service layer while taking product", err.Error())
		return models.ProductSell{}, err
	}

	if err = p.storage.BasketProduct().AddProducts(ctx, basket.ID, basketProducts); err != nil {
		fmt.Println("error in service later while adding products to basket", err.Error())
		return models.ProductSell{}, err
	}

	if err = p.storage.Store().AddProfit(ctx, profit, customer.BranchID); err != nil {
		fmt.Println("error in service layer while adding amount of profit", err.Error())
		return models.ProductSell{}, err
	}

	//prixod
	//check
	//report

	return productSell, nil
}
