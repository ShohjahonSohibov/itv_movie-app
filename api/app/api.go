package app

import (
	"itv_movie_app/api/handlers"
	"itv_movie_app/api/middleware"
	"itv_movie_app/config"
	"itv_movie_app/internal/models"
	"itv_movie_app/internal/repository"
	"itv_movie_app/internal/service"
	"itv_movie_app/pkg/auth"
	"itv_movie_app/pkg/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(router *gin.Engine, db *gorm.DB) error {
    // Migrate database
    if err := database.AutoMigrate(db); err != nil {
        return err
    }

    // Initialize managers
    repoManager := repository.NewRepositories(db)
    serviceManager := service.NewServices(repoManager)
    
	// Initialize handlers
	movieHandler := handlers.NewMovieHandler(serviceManager.MovieService)

	// API routes
	api := router.Group("/api/v1")
	{
		protected := api.Group("")
		config, err := config.Load()
		if err != nil {
			return err
		}

		secret := models.JWTConfig{
			Secret: config.JWT.Secret,
		}
		auth := auth.NewJWTManager(secret)
		protected.Use(middleware.AuthMiddleware(auth))

		// Public routes (no auth required)
		movies := api.Group("/movies")
		{
			movies.GET("", movieHandler.GetAllMovies)
			movies.GET("/:id", movieHandler.GetMovie)
			movies.POST("", movieHandler.CreateMovie)
			movies.PUT("/:id", movieHandler.UpdateMovie)
			movies.DELETE("/:id", movieHandler.DeleteMovie)
		}
	}

	return nil
}
