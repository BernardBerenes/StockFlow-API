package main

import (
	"fmt"

	"github.com/BernardBerenes/stockflow-api/internal/config"
)

func main() {
	fiber := config.NewFiber()
	viper := config.NewViper()
	gorm := config.NewGorm(viper)
	validator := config.NewValidator()

	app := config.NewApp(fiber, viper, gorm, validator)
	app.Bootstrap()

	err := fiber.Listen(fmt.Sprintf(":%d", viper.GetInt("APP_PORT")))
	if err != nil {
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}
}
