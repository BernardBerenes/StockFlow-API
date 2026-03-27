package config

import (
	"fmt"

	"github.com/BernardBerenes/stockflow-api/api/routes"
	"github.com/BernardBerenes/stockflow-api/pkg/helper"
	"github.com/BernardBerenes/stockflow-api/pkg/product"
	"github.com/BernardBerenes/stockflow-api/pkg/store"
	"github.com/BernardBerenes/stockflow-api/pkg/transaction"
	"github.com/BernardBerenes/stockflow-api/pkg/transaction_detail"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type App struct {
	fiber     *fiber.App
	viper     *viper.Viper
	gorm      *gorm.DB
	minio     *minio.Client
	validator *validator.Validate
}

func NewApp(fiber *fiber.App, viper *viper.Viper, gorm *gorm.DB, minio *minio.Client, validator *validator.Validate) *App {
	return &App{
		fiber:     fiber,
		viper:     viper,
		gorm:      gorm,
		minio:     minio,
		validator: validator,
	}
}

func (a *App) Bootstrap() {
	Migrate(a.gorm)
	helpers := helper.NewHelper(a.viper, a.minio)

	err := a.validator.RegisterValidation("check_duplicate_products", CheckDuplicateProduct)
	if err != nil {
		panic(fmt.Errorf("fatal error registering validation: %w", err))
	}

	storeRepository := store.NewRepository(a.gorm)
	storeService := store.NewService(storeRepository, a.validator)
	productRepository := product.NewRepository(a.gorm)
	productService := product.NewService(productRepository, helpers, a.validator)
	transactionRepository := transaction.NewRepository(a.gorm)
	transactionService := transaction.NewService(transactionRepository, storeRepository, a.validator)
	transactionDetailRepository := transaction_detail.NewRepository(a.gorm)
	transactionDetailService := transaction_detail.NewService(transactionDetailRepository, transactionRepository, productRepository, a.validator)

	api := a.fiber.Group("api")
	storeRoute := api.Group("store")
	productRoute := api.Group("product")
	transactionRoute := api.Group("transaction")
	transactionDetailRoute := api.Group("transaction-detail")

	routes.StoreRouter(storeRoute, storeService)
	routes.ProductRouter(productRoute, productService)
	routes.TransactionRouter(transactionRoute, transactionService)
	routes.TransactionDetailRouter(transactionDetailRoute, transactionDetailService)
}
