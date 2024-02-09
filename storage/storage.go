package storage

import (
	"context"
	"test/api/models"
)

type IStorage interface {
	Close()
	User() IUserStorage
	Category() ICategoryStorage
	Product() IProductStorage
	Basket() IBasketStorage
	BasketProduct() IBasketProductStorage
	Store() IStoreStorage
	Branch() IBranchStorage
	Dealer() IDealerStorage
	Income() IIncomeStorage
	IncomeProduct() IIncomeProductStorage
}

type IUserStorage interface {
	Create(context.Context, models.CreateUser) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.User, error)
	GetList(context.Context, models.GetListRequest) (models.UsersResponse, error)
	Update(context.Context, models.UpdateUser) (string, error)
	Delete(context.Context, models.PrimaryKey) error
	GetPassword(context.Context, string) (string, error)
	UpdatePassword(context.Context, models.UpdateUserPassword) error
	UpdateCustomerCash(context.Context, string, int) error
}

type ICategoryStorage interface {
	Create(context.Context, models.CreateCategory) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.Category, error)
	GetList(context.Context, models.GetListRequest) (models.CategoryResponse, error)
	Update(context.Context, models.UpdateCategory) (string, error)
	Delete(context.Context, models.PrimaryKey) error
}

type IProductStorage interface {
	Create(context.Context, models.CreateProduct) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.Product, error)
	GetList(context.Context, models.GetListRequest) (models.ProductResponse, error)
	Update(context.Context, models.UpdateProduct) (string, error)
	Delete(context.Context, models.PrimaryKey) error
	Search(context.Context, map[string]int) (models.ProductSell, error)
	TakeProducts(context.Context, map[string]int) error
	AddDeliveredProducts(context.Context, models.DeliverProducts, string) error
}
type IBasketStorage interface {
	Create(context.Context, models.CreateBasket) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.Basket, error)
	GetList(context.Context, models.GetListRequest) (models.BasketResponse, error)
	Update(context.Context, models.UpdateBasket) (string, error)
	Delete(context.Context, models.PrimaryKey) error
}

type IBasketProductStorage interface {
	Create(context.Context, models.CreateBasketProduct) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.BasketProduct, error)
	GetList(context.Context, models.GetListRequest) (models.BasketProductResponse, error)
	Update(context.Context, models.UpdateBasketProduct) (string, error)
	Delete(context.Context, models.PrimaryKey) error
	AddProducts(context.Context, string, map[string]int) error
}

type IStoreStorage interface {
	AddProfit(ctx context.Context, profit float32, branchID string) error
	GetStoreBudget(context.Context, string) (float32, error)
	WithdrawalDeliveredSum(context.Context, float32, string) error
}

type IDealerStorage interface {
	AddSum(context.Context, int) error
}

type IBranchStorage interface {
	Create(context.Context, models.CreateBranch) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.Branch, error)
	GetList(context.Context, models.GetListRequest) (models.BranchResponse, error)
	Update(context.Context, models.UpdateBranch) (string, error)
	Delete(context.Context, models.PrimaryKey) error
}

type IIncomeStorage interface {
	Create(ctx context.Context) (models.Income, error)
}

type IIncomeProductStorage interface {
	CreateMultiple(context.Context, models.CreateIncomeProducts) error
}
