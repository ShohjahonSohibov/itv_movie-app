package repository

import "gorm.io/gorm"

type Repositories struct {
    db    *gorm.DB
    Movie *MovieRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
    return &Repositories{
        db:    db,
        Movie: NewMovieRepository(db),
    }
}