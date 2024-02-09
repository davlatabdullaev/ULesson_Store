package service

import (
	"context"
	"fmt"
	"test/api/models"
	"test/storage"
)

type dealerService struct {
	storage storage.IStorage
}

func NewDealerService(storage storage.IStorage) dealerService {
	return dealerService{storage}
}

func (d dealerService) Delivery(ctx context.Context, sell models.ProductSell) error {
	var (
		totalSum = 0
	)

	for productID, quantity := range sell.NotEnoughProducts {
		totalSum += quantity * sell.NotEnoughProductPrices[productID]
	}

	budget, err := d.storage.Store().GetStoreBudget(ctx, sell.ProductsBranchID)
	if err != nil {
		fmt.Println("error in service layer while getting store budget", err.Error())
		return err
	}

	if budget < float32(totalSum) {
		fmt.Println("not enough budget", err)
		return err
	}

	if err = d.storage.Product().AddDeliveredProducts(ctx, models.DeliverProducts{
		NotEnoughProducts: sell.NotEnoughProducts,
	}, sell.ProductsBranchID); err != nil {
		fmt.Println("error in service layer while adding delivered products", err.Error())
		return err
	}

	if err = d.storage.Store().WithdrawalDeliveredSum(ctx, float32(totalSum), sell.ProductsBranchID); err != nil {
		fmt.Println("error in service layer while remove delivered sum", err.Error())
		return err
	}

	if err = d.storage.Dealer().AddSum(ctx, totalSum); err != nil {
		fmt.Println("error in service layer while add sum to dealer", err.Error())
		return err
	}

	return nil
}
