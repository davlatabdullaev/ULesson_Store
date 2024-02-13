package postgres

import (
	"context"
	"test/api/models"
	"test/config"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCategoryRepo_Create(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	createCategory := models.CreateCategory{
		Name: "category name",
	}

	categoryID, err := pgStore.Category().Create(context.Background(), createCategory)
	if err != nil {
		t.Errorf("error while creating catigory error: %v", err)
	}

	user, err := pgStore.Category().GetByID(context.Background(), models.PrimaryKey{
		ID: categoryID,
	})
	if err != nil {
		t.Errorf("error while getting catigory error: %v", err)
	}

	assert.Equal(t, user.Name, createCategory.Name)

}

func TestCategoryRepo_GetByID(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}
	createCategory := models.CreateCategory{
		Name: "category_name_2",
	}

	categoryID, err := pgStore.Category().Create(context.Background(), createCategory)
	if err != nil {
		t.Errorf("error while creating user error: %v", err)
	}

	category, err := pgStore.Category().GetByID(context.Background(), models.PrimaryKey{
		ID: categoryID,
	})
	if err != nil {
		t.Errorf("error while getting primary key : %v", err)
	}

	assert.Equal(t, category.ID, categoryID)
}

func TestCategoryRepo_GetList(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	categoryResp, err := pgStore.Category().GetList(context.Background(), models.GetListRequest{
		Page:  1,
		Limit: 1000,
	})
	if err != nil {
		t.Errorf("error while getting categoryResp error: %v", err)
	}

	if len(categoryResp.Category) != 7 {
		t.Errorf("expected 7, but got: %d", len(categoryResp.Category))
	}

	assert.Equal(t, len(categoryResp.Category), 7)

}

func TestCategoryRepo_Update(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}
	createCategory := models.CreateCategory{
		Name: "name3",
	}

	categoryID, err := pgStore.Category().Create(context.Background(), createCategory)
	if err != nil {
		t.Errorf("error while creating category error: %v", err)
	}

	updateCategory := models.UpdateCategory{
		ID:   categoryID,
		Name: "updatedName",
	}

	categoryUpdateID, err := pgStore.Category().Update(context.Background(), updateCategory)
	if err != nil {
		t.Errorf("error while creating category error: %v", err)
	}

	assert.Equal(t, categoryID, categoryUpdateID)
}

func TestCategoryRepo_Delete(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	createCategory := models.CreateCategory{
		Name: "name4",
	}

	categoryID, err := pgStore.Category().Create(context.Background(), createCategory)
	if err != nil {
		t.Errorf("error while creating category error: %v", err)
	}

	if err := pgStore.Category().Delete(context.Background(), models.PrimaryKey{
		ID: categoryID,
	}); err != nil {
		t.Errorf("Error deleting category: %v", err)
	}
}
