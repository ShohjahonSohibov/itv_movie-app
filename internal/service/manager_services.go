package service

import "itv_movie_app/internal/repository"

type Services struct {
    MovieService *MovieService
}

func NewServices(repos *repository.Repositories) *Services {
    return &Services{
        MovieService: NewMovieService(repos.Movie),
    }
}