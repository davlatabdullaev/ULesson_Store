package postgres

import (
	"context"
	"test/api/models"
	"test/config"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestBasketRepo_Create(t *testing.T) {

	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	createBasket := models.CreateBasket{
		CustomerID: "6df52ee6-8dbf-42fb-be0d-9810038f909fq",
		TotalSum:   1000,
	}

	basketID, err := pgStore.Basket().Create(context.Background(), createBasket)
	if err != nil {
		t.Errorf("error while creating basket error : %v", err)
	}

	basket, err := pgStore.Basket().GetByID(context.Background(), models.PrimaryKey{
		ID: basketID,
	})
	if err != nil {
		t.Errorf("error while getting basket error: %v", err)
	}

	assert.Equal(t, basket.CustomerID, createBasket.CustomerID)
	assert.Equal(t, basket.TotalSum, createBasket.TotalSum)

}

func TestBasketRepo_GetByID(t *testing.T) {

	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}
	createBasket := models.CreateBasket{
		TotalSum: 0,
	}

	basketID, err := pgStore.Basket().Create(context.Background(), createBasket)
	if err != nil {
		t.Errorf("error while creating basket error: %v", err)
	}

	basket, err := pgStore.Basket().GetByID(context.Background(), models.PrimaryKey{
		ID: basketID,
	})
	if err != nil {
		t.Errorf("error while getting primary key : %v", err)
	}

	assert.Equal(t, basket.ID, basketID)

}

func TestBasketRepo_GetList(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	basketResp, err := pgStore.Basket().GetList(context.Background(), models.GetListRequest{
		Page:  1,
		Limit: 1000,
	})
	if err != nil {
		t.Errorf("error while getting basketResp error: %v", err)
	}

	if len(basketResp.Baskets) != 1 {
		t.Errorf("expected 1, but got: %d", len(basketResp.Baskets))
	}

	assert.Equal(t, len(basketResp.Baskets), 2)

}

func TestBasketRepo_Update(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}
	createBasket := models.CreateBasket{
		CustomerID: "6df52ee6-8dbf-42fb-be0d-9810038f909f",
		TotalSum:   0,
	}

	basketID, err := pgStore.Basket().Create(context.Background(), createBasket)
	if err != nil {
		t.Errorf("error while creating basket error: %v", err)
	}

	updateBasket := models.UpdateBasket{
		ID:         basketID,
		CustomerID: "6df52ee6-8dbf-42fb-be0d-9810038f909f",
		TotalSum:   0,
	}

	updateBasketID, err := pgStore.Basket().Update(context.Background(), updateBasket)
	if err != nil {
		t.Errorf("error while creating basket error: %v", err)
	}

	assert.Equal(t, basketID, updateBasketID)
}

func TestBasketRepo_Delete(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	createBasket := models.CreateBasket{
		CustomerID: "",
		TotalSum:   0,
	}

	basketID, err := pgStore.Basket().Create(context.Background(), createBasket)
	if err != nil {
		t.Errorf("error while creating basket error: %v", err)
	}

	if err := pgStore.Basket().Delete(context.Background(), models.PrimaryKey{
		ID: basketID,
	}); err != nil {
		t.Errorf("Error deleting basket: %v", err)
	}
}
