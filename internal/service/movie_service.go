package service

import (
	"itv_movie_app/internal/models"
	"itv_movie_app/internal/repository"
)

type MovieService struct {
	repo *repository.MovieRepository
}

func NewMovieService(repo *repository.MovieRepository) *MovieService {
	return &MovieService{repo: repo}
}

func (s *MovieService) GetMovie(id string) (*models.Movie, error) {
	return s.repo.GetByID(id)
}

func (s *MovieService) GetAllMovies(filter *models.MovieListRequest) (*models.MovieListResponse, error) {
	return s.repo.GetAll(filter)
}

func (s *MovieService) CreateMovie(movie *models.Movie) error {
	return s.repo.Create(movie)
}

func (s *MovieService) UpdateMovie(movie *models.Movie) error {
	return s.repo.Update(movie)
}

func (s *MovieService) DeleteMovie(id string) error {
	return s.repo.Delete(id)
}
