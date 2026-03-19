package config

import "github.com/gofiber/fiber/v3"

func NewFiber() *fiber.App {
	return fiber.New(fiber.Config{
		AppName:       "StockFlow",
		CaseSensitive: true,
		StrictRouting: true,
	})
}
