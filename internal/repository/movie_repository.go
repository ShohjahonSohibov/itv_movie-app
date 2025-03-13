package repository

import (
	"itv_movie_app/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r *MovieRepository) GetByID(id string) (*models.Movie, error) {
	var movie models.Movie

	// Parse string ID to UUID
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	// Safe: GORM handles parameter sanitization
	if err := r.db.Where("id = ?", parsedID).First(&movie).Error; err != nil {
		return nil, err
	}

	return &movie, nil
}

func (r *MovieRepository) GetAll(filter *models.MovieListRequest) (*models.MovieListResponse, error) {
	query := r.db.Model(&models.Movie{})

	// Apply filters
	if filter.Title != "" {
		query = query.Where("title ILIKE ?", "%"+filter.Title+"%")
	}
	if filter.Director != "" {
		query = query.Where("director ILIKE ?", "%"+filter.Director+"%")
	}
	if filter.YearFrom != 0 {
		query = query.Where("year >= ?", filter.YearFrom)
	}
	if filter.YearTo != 0 {
		query = query.Where("year <= ?", filter.YearTo)
	}
	if filter.MinImdbRating > 0 {
		query = query.Where("imdb_rating >= ?", filter.MinImdbRating)
	}
	if filter.MinItvRating > 0 {
		query = query.Where("itv_rating >= ?", filter.MinItvRating)
	}
	if filter.MinKinoRating > 0 {
		query = query.Where("kinopoisk_rating >= ?", filter.MinKinoRating)
	}

	query = query.Offset(filter.Offset).Limit(filter.Limit)

	var movies []*models.Movie
	if err := query.Find(&movies).Error; err != nil {
		return nil, err
	}

	res := models.MovieListResponse{
		Count: len(movies),
		Items: movies,
	}

	return &res, nil
}

func (r *MovieRepository) Create(movie *models.Movie) error {
    return r.db.Transaction(func(tx *gorm.DB) error {
        // Generate new UUID
        movie.ID = uuid.New()
        movie.CreatedAt = time.Now()

        if err := tx.Create(movie).Error; err != nil {
            return err
        }

        // other entities

        return nil
    })
}

func (r *MovieRepository) Update(movie *models.Movie) error {
    return r.db.Transaction(func(tx *gorm.DB) error {
        movie.UpdatedAt = time.Now()

        result := tx.Model(&models.Movie{}).
            Where("id = ?", movie.ID).
            Updates(movie)

        if result.Error != nil {
            return result.Error
        }

        if result.RowsAffected == 0 {
            return gorm.ErrRecordNotFound
        }

        // other entities

        return nil
    })
}

func (r *MovieRepository) Delete(id string) error {
    return r.db.Transaction(func(tx *gorm.DB) error {
        // Parse UUID
        parsedID, err := uuid.Parse(id)
        if err != nil {
            return err
        }

        result := tx.Where("id = ?", parsedID).Delete(&models.Movie{})
        if result.Error != nil {
            return result.Error
        }

        if result.RowsAffected == 0 {
            return gorm.ErrRecordNotFound
        }

        // other entities

        return nil
    })
}
