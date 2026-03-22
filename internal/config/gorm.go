package config

import (
	"fmt"

	"github.com/BernardBerenes/stockflow-api/pkg/entities"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGorm(dbConfig *viper.Viper) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", dbConfig.GetString("DB_HOST"), dbConfig.GetUint("DB_PORT"), dbConfig.GetString("DB_USER"), dbConfig.GetString("DB_PASSWORD"), dbConfig.GetString("DB_NAME"), dbConfig.GetString("APP_TIMEZONE"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(fmt.Errorf("fatal error opening postgres connection: %w", err))
	}

	return db
}

func SetupEnums(db *gorm.DB) {
	db.Exec(`
	DO $$ BEGIN
		CREATE TYPE transaction_type AS ENUM ('IN', 'OUT');
	EXCEPTION
		WHEN duplicate_object THEN null;
	END $$;
	`)

	db.Exec(`
	DO $$ BEGIN
		CREATE TYPE payment_status AS ENUM ('UNPAID', 'PAID');
	EXCEPTION
		WHEN duplicate_object THEN null;
	END $$;
	`)

	db.Exec(`
	DO $$ BEGIN
		CREATE TYPE delivery_status AS ENUM ('ON_DELIVERY', 'DELIVERED');
	EXCEPTION
		WHEN duplicate_object THEN null;
	END $$;
	`)
}

func Migrate(db *gorm.DB) {
	SetupEnums(db)

	err := db.AutoMigrate(&entities.Product{}, &entities.Store{}, &entities.Transaction{}, &entities.TransactionDetail{})
	if err != nil {
		panic(fmt.Errorf("fatal error migrating gorm: %w", err))
	}
}
