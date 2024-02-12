package service

import (
	"context"
	"errors"
	"fmt"
	"test/api/models"
	"test/pkg/check"
	"test/storage"

	"github.com/jackc/pgx/v5"
)

type userService struct {
	storage storage.IStorage
}

func NewUserService(storage storage.IStorage) userService {
	return userService{
		storage: storage,
	}
}

func (u userService) Create(ctx context.Context, createUser models.CreateUser) (models.User, error) {
	pKey, err := u.storage.User().Create(ctx, createUser)
	if err != nil {
		fmt.Println("ERROR in service layer while creating user", err.Error())
		return models.User{}, err
	}

	user, err := u.storage.User().GetByID(ctx, models.PrimaryKey{
		ID: pKey,
	})

	return user, nil
}

func (u userService) GetUser(ctx context.Context, pKey models.PrimaryKey) (models.User, error) {
	user, err := u.storage.User().GetByID(ctx, pKey)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("ERROR in service layer while getting user by id", err.Error())
			return models.User{}, err
		}
	}

	return user, nil
}

func (u userService) GetUsers(ctx context.Context, request models.GetListRequest) (models.UsersResponse, error) {
	usersResponse, err := u.storage.User().GetList(ctx, request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("ERROR in service layer while getting users list", err.Error())
			return models.UsersResponse{}, err
		}
	}

	return usersResponse, err
}

func (u userService) Update(ctx context.Context, updateUser models.UpdateUser) (models.User, error) {

	pKey, err := u.storage.User().Update(ctx, updateUser)
	if err != nil {
		fmt.Println("ERROR in service layer while updating updateUser", err.Error())
		return models.User{}, err
	}

	user, err := u.storage.User().GetByID(ctx, models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		fmt.Println("ERROR in service layer while getting user after update", err.Error())
		return models.User{}, err
	}

	return user, nil
}

func (u userService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := u.storage.User().Delete(ctx, key)
	return err
}

func (u userService) UpdatePassword(ctx context.Context, request models.UpdateUserPassword) error {
	oldPassword, err := u.storage.User().GetPassword(ctx, request.ID)
	if err != nil {
		fmt.Println("ERROR in service layer while getting user password", err.Error())
		return err
	}

	if oldPassword != request.OldPassword {
		fmt.Println("ERROR in service old password is not correct")
		return errors.New("old password did not match")
	}

	if err = check.ValidatePassword(request.NewPassword); err != nil {
		fmt.Println("ERROR in service layer new password is weak", err.Error())
		return err
	}

	if err = u.storage.User().UpdatePassword(context.Background(), request); err != nil {
		fmt.Println("ERROR in service layer while updating password", err.Error())
		return err
	}

	return nil
}
