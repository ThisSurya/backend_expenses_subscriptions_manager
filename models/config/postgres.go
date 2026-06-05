package db

import (
	"backend/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPostgresDB(host, port, user, password, dbname string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL database: \n", err)
	}

	log.Println("Successfully connected to PostgreSQL database")
	if err := ensureEnumTypes(db); err != nil {
		log.Fatal("Failed to ensure enum types: \n", err)
	}
	if err := db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Expense{},
		&models.Subscription{},
		&models.Notification{},
		&models.RefreshToken{},
	); err != nil {
		log.Fatal("Failed to auto-migrate models: \n", err)
	}

	return db
}

func ensureEnumTypes(db *gorm.DB) error {
	statements := []string{
		"DO $$ BEGIN CREATE TYPE user_role AS ENUM ('basic', 'premium', 'admin'); EXCEPTION WHEN duplicate_object THEN null; END $$;",
		"DO $$ BEGIN CREATE TYPE notification_type AS ENUM ('email'); EXCEPTION WHEN duplicate_object THEN null; END $$;",
		"DO $$ BEGIN CREATE TYPE notification_status AS ENUM ('pending', 'sent', 'failed'); EXCEPTION WHEN duplicate_object THEN null; END $$;",
	}

	for _, stmt := range statements {
		if err := db.Exec(stmt).Error; err != nil {
			return err
		}
	}

	return nil
}
