package database

import (
    "itv_movie_app/internal/models"
    "gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
    return db.AutoMigrate(
        &models.Movie{},
    )
}