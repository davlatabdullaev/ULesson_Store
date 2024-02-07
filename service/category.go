package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"test/api/models"
	"test/storage"
)

type categoryService struct {
	storage storage.IStorage
}

func NewCategoryService(storage storage.IStorage) categoryService {
	return categoryService{
		storage: storage,
	}
}

func (c categoryService) Create(ctx context.Context, createCategory models.CreateCategory) (models.Category, error) {
	pKey, err := c.storage.Category().Create(ctx, createCategory)
	if err != nil {
		fmt.Println("ERROR in service layer while creating category", err.Error())
		return models.Category{}, err
	}

	category, err := c.storage.Category().GetByID(ctx, models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		fmt.Println("ERROR in service layer while getting category", err.Error())
		return models.Category{}, err
	}

	return category, nil
}

func (c categoryService) Get(ctx context.Context, key models.PrimaryKey) (models.Category, error) {
	category, err := c.storage.Category().GetByID(ctx, key)
	if err != nil {
		fmt.Println("error is in service layer while getting by id", err.Error())
		return models.Category{}, err
	}
	return category, nil
}

func (c categoryService) GetList(ctx context.Context, request models.GetListRequest) (models.CategoryResponse, error) {
	categories, err := c.storage.Category().GetList(ctx, request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting list", err.Error())
			return models.CategoryResponse{}, err
		}
	}
	return categories, nil
}

func (c categoryService) Update(ctx context.Context, category models.UpdateCategory) (models.Category, error) {
	id, err := c.storage.Category().Update(ctx, category)
	if err != nil {
		fmt.Println("error in service layer while updating category", err.Error())
		return models.Category{}, err
	}

	updatedCategory, err := c.storage.Category().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		fmt.Println("error in service layer while getting by id", err.Error())
		return models.Category{}, err
	}

	return updatedCategory, nil
}

func (c categoryService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := c.storage.Category().Delete(ctx, key)

	return err
}