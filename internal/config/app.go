package config

import (
	"github.com/BernardBerenes/stockflow-api/api/routes"
	"github.com/BernardBerenes/stockflow-api/pkg/store"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type App struct {
	fiber     *fiber.App
	viper     *viper.Viper
	gorm      *gorm.DB
	validator *validator.Validate
}

func NewApp(fiber *fiber.App, viper *viper.Viper, gorm *gorm.DB, validator *validator.Validate) *App {
	return &App{
		fiber:     fiber,
		viper:     viper,
		gorm:      gorm,
		validator: validator,
	}
}

func (a *App) Bootstrap() {
	Migrate(a.gorm)

	storeRepository := store.NewRepository(a.gorm)
	storeService := store.NewService(storeRepository, a.validator)

	api := a.fiber.Group("api")
	storeRoute := api.Group("store")

	routes.StoreRouter(storeRoute, storeService)
}
