package database

import (
	"fmt"
	"itv_movie_app/config"
	"itv_movie_app/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase() (*gorm.DB, error) {

	config, err := config.Load()
	if err != nil {
		return nil, err
	}
	fmt.Println("DB_HOST:", config.Database.Host)
	cfg := models.DatabaseConfig{
		Host:     config.Database.Host,
		Port:     config.Database.Port,
		User:     config.Database.User,
		Password: config.Database.Password,
		DBName:   config.Database.DBName,
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s",
	cfg.User, cfg.Password, cfg.Host, cfg.DBName)

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	// 	cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true, // Allow schema modifications
	})
	if err != nil {
		return nil, fmt.Errorf("postgres connection error: %w", err)
	}

	// Auto-migrate with cleanup
	if err := db.AutoMigrate(&models.Movie{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, nil
}
