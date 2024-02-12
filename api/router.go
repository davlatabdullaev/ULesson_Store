package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "test/api/docs"
	"test/api/handler"
	"test/service"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
func New(services service.IServiceManager) *gin.Engine {
	h := handler.New(services)

	r := gin.New()

	r.POST("/user", h.CreateUser)
	r.GET("/user/:id", h.GetUser)
	r.GET("/users", h.GetUserList)
	r.PUT("/user/:id", h.UpdateUser)
	r.DELETE("/user/:id", h.DeleteUser)
	r.PATCH("/user/:id", h.UpdateUserPassword)

	r.POST("/category", h.CreateCategory)
	r.GET("/category/:id", h.GetCategory)
	r.GET("/categories", h.GetCategoryList)
	r.PUT("/category/:id", h.UpdateCategory)
	r.DELETE("/category/:id", h.DeleteCategory)

	r.POST("/product", h.CreateProduct)
	r.GET("/product/:id", h.GetProduct)
	r.GET("/products", h.GetProductList)
	r.PUT("/product/:id", h.UpdateProduct)
	r.DELETE("/product/:id", h.DeleteProduct)

	r.POST("/basket", h.CreateBasket)
	r.GET("/basket/:id", h.GetBasket)
	r.GET("/baskets", h.GetBasketList)
	r.PUT("basket/:id", h.UpdateBasket)
	r.DELETE("basket/:id", h.DeleteBasket)

	r.POST("/basketProduct", h.CreateBasketProduct)
	r.GET("/basketProduct/:id", h.GetBasketProduct)
	r.GET("/basketProducts", h.GetBasketProductList)
	r.PUT("/basketProduct/:id", h.UpdateBasketProduct)
	r.DELETE("/basketProduct/:id", h.DeleteBasketProduct)

	r.POST("/branch", h.CreateBranch)
	r.GET("/branch/:id", h.GetBranch)
	r.GET("/branches", h.GetBranchList)
	r.PUT("/branch/:id", h.UpdateBranch)
	r.DELETE("/branch/:id", h.DeleteBranch)

	r.POST("/income", h.CreateIncome) // create
	r.GET("/income/:id", h.GetIncome)              // get by id
	r.GET("/incomes", h.GetIncomeList)                 // get list
	r.DELETE("/income/:id", h.DeleteIncome)           // delete

	r.POST("/income_products", h.CreateIncomeProducts) // create multiple
	r.GET("/income_products", h.GetIncomeProductsList)   // get income products (filter => by income_id)
	r.PUT("/income_products", h.UpdateIncomeProducts)    // update multiple
	r.DELETE("/income_products", h.DeleteIncomeProducts) // delete multiple

	r.POST("/sell-new", h.StartSellNew)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}