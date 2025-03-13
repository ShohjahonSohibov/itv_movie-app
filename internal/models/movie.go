package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Movie struct {
	ID              uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title           string         `gorm:"size:255;not null;index" json:"title" binding:"required,min=1,max=255"`
	Director        string         `gorm:"size:255;index" json:"director,omitempty" binding:"omitempty,max=255"`
	Year            int            `gorm:"index" json:"year,omitempty" binding:"omitempty,min=1888,max=2100"`
	Plot            string         `gorm:"type:text" json:"plot,omitempty" binding:"omitempty,max=5000"`
	ImdbRating      float32        `gorm:"type:decimal(3,1);index" json:"imdb_rating,omitempty" binding:"omitempty,min=0,max=10"`
	ItvRating       float32        `gorm:"type:decimal(3,1);index" json:"itv_rating,omitempty" binding:"omitempty,min=0,max=10"`
	KinopoiskRating float32        `gorm:"type:decimal(3,1);index" json:"kinopoisk_rating,omitempty" binding:"omitempty,min=0,max=10"`
	Duration        int            `gorm:"index" json:"duration,omitempty" binding:"omitempty,min=1,max=1000"`
	Budget          int64          `gorm:"index" json:"budget,omitempty" binding:"omitempty,min=0"`
	CreatedAt       time.Time      `json:"created_at,omitempty"`
	UpdatedAt       time.Time      `json:"updated_at,omitempty"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

type CreateMovieRequest struct {
	Title           string   `json:"title" binding:"required,min=1,max=255"`
	Director        string   `json:"director,omitempty" binding:"omitempty,max=255"`
	Year            int      `json:"year,omitempty" binding:"omitempty,min=1888,max=2100"`
	Plot            string   `json:"plot,omitempty" binding:"omitempty,max=5000"`
	ImdbRating      float32  `json:"imdb_rating,omitempty" binding:"omitempty,min=0,max=10"`
	ItvRating       float32  `json:"itv_rating,omitempty" binding:"omitempty,min=0,max=10"`
	KinopoiskRating float32  `json:"kinopoisk_rating,omitempty" binding:"omitempty,min=0,max=10"`
	Duration        int      `json:"duration,omitempty" binding:"omitempty,min=1,max=1000"`
	Budget          int64    `json:"budget,omitempty" binding:"omitempty,min=0"`
}

type UpdateMovieRequest struct {
	Title           string    `json:"title" binding:"omitempty,min=1,max=255"`
	Director        string    `json:"director,omitempty" binding:"omitempty,max=255"`
	Year            int       `json:"year,omitempty" binding:"omitempty,min=1888,max=2100"`
	Plot            string    `json:"plot,omitempty" binding:"omitempty,max=5000"`
	ImdbRating      float32   `json:"imdb_rating,omitempty" binding:"omitempty,min=0,max=10"`
	ItvRating       float32   `json:"itv_rating,omitempty" binding:"omitempty,min=0,max=10"`
	KinopoiskRating float32   `json:"kinopoisk_rating,omitempty" binding:"omitempty,min=0,max=10"`
	Duration        int       `json:"duration,omitempty" binding:"omitempty,min=1,max=1000"`
	Budget          int64     `json:"budget,omitempty" binding:"omitempty,min=0"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
}

type MovieListRequest struct {
	Filter
	Title         string  `json:"title,omitempty" binding:"omitempty,max=255"`
	Director      string  `json:"director,omitempty" binding:"omitempty,max=255"`
	YearFrom      int     `json:"year_from,omitempty" binding:"omitempty,min=1888,max=2100"`
	YearTo        int     `json:"year_to,omitempty" binding:"omitempty,min=1888,max=2100,gtefield=YearFrom"`
	MinImdbRating float32 `json:"min_imdb_rating,omitempty" binding:"omitempty,min=0,max=10"`
	MinItvRating  float32 `json:"min_itv_rating,omitempty" binding:"omitempty,min=0,max=10"`
	MinKinoRating float32 `json:"min_kinopoisk_rating,omitempty" binding:"omitempty,min=0,max=10"`
}

type MovieListResponse struct {
	Count int      `json:"count,omitempty"`
	Items []*Movie `json:"items,omitempty"`
}

// Add BeforeCreate hook to ensure UUID is set
func (m *Movie) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}
