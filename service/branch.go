package service

import (
	"context"
	"fmt"
	"test/api/models"
	"test/storage"
)

type branchService struct {
	storage storage.IStorage
}

func NewBranchService(storage storage.IStorage) branchService {
	return branchService{storage: storage}
}

func (b branchService) Create(ctx context.Context, branch models.CreateBranch) (models.Branch, error) {
	id, err := b.storage.Branch().Create(ctx, branch)
	if err != nil {
		fmt.Println("error in service layer while creating branch", err.Error())
		return models.Branch{}, err
	}

	createdBranch, err := b.storage.Branch().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		fmt.Println("error in service layer while getting branch by id", err.Error())
		return models.Branch{}, err
	}

	return createdBranch, nil
}

func (b branchService) Get(ctx context.Context, id string) (models.Branch, error) {
	branch, err := b.storage.Branch().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		fmt.Println("error in service layer while getting branch by id", err.Error())
		return models.Branch{}, err
	}

	return branch, nil
}

func (b branchService) GetList(ctx context.Context, request models.GetListRequest) (models.BranchResponse, error) {
	branches, err := b.storage.Branch().GetList(ctx, request)
	if err != nil {
		fmt.Println("error in service layer while getting list", err.Error())
		return models.BranchResponse{}, err
	}

	return branches, nil
}

func (b branchService) Update(ctx context.Context, branch models.UpdateBranch) (models.Branch, error) {
	id, err := b.storage.Branch().Update(ctx, branch)
	if err != nil {
		fmt.Println("error in service layer while updating branch", err.Error())
		return models.Branch{}, err
	}

	updatedBranch, err := b.storage.Branch().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		fmt.Println("error in service layer while getting  branch by id", err.Error())
		return models.Branch{}, err
	}

	return updatedBranch, nil
}

func (b branchService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := b.storage.Branch().Delete(ctx, key)

	return err
}
