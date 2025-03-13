package handlers

import (
	"errors"
	"itv_movie_app/internal/models"
	"itv_movie_app/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MovieHandler struct {
	movieService *service.MovieService
}

func NewMovieHandler(movieService *service.MovieService) *MovieHandler {
	return &MovieHandler{movieService: movieService}
}

// GetMovie godoc
// @Summary Get a movie by ID
// @Description Get details of a specific movie by its ID
// @Tags movies
// @Accept json
// @Produce json
// @Param id path string true "Movie ID"
// @Success 200 {object} models.Movie
// @Failure 400 {object} map[string]string "Invalid ID format"
// @Failure 404 {object} map[string]string "Movie not found"
// @Router /movies/{id} [get]
func (h *MovieHandler) GetMovie(c *gin.Context) {
	id := c.Param("id")

	movie, err := h.movieService.GetMovie(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

// GetAllMovies godoc
// @Summary List all movies
// @Description Get a list of all movies with optional filtering
// @Tags movies
// @Accept json
// @Produce json
// @Param request query models.MovieListRequest false "Filter parameters"
// @Success 200 {object} models.MovieListResponse
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /movies [get]
func (h *MovieHandler) GetAllMovies(c *gin.Context) {
	var filter models.MovieListRequest

	offset, limit, err := getPageOffsetLimit(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	filter.Filter.Offset = offset
	filter.Filter.Limit = limit
	filter.Director = c.Query("director")
	filter.Title = c.Query("title")

	// Convert string to int for numeric fields
	if yearFrom := c.Query("year_from"); yearFrom != "" {
		if year, err := strconv.Atoi(yearFrom); err == nil {
			filter.YearFrom = year
		}
	}
	if yearTo := c.Query("year_to"); yearTo != "" {
		if year, err := strconv.Atoi(yearTo); err == nil {
			filter.YearTo = year
		}
	}

	// Convert string to float32 for ratings
	if rating := c.Query("min_imdb_rating"); rating != "" {
		if r, err := strconv.ParseFloat(rating, 32); err == nil {
			filter.MinImdbRating = float32(r)
		}
	}
	if rating := c.Query("min_itv_rating"); rating != "" {
		if r, err := strconv.ParseFloat(rating, 32); err == nil {
			filter.MinItvRating = float32(r)
		}
	}
	if rating := c.Query("min_kinopoisk_rating"); rating != "" {
		if r, err := strconv.ParseFloat(rating, 32); err == nil {
			filter.MinKinoRating = float32(r)
		}
	}

	movies, err := h.movieService.GetAllMovies(&filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

// CreateMovie godoc
// @Summary Create a new movie
// @Description Create a new movie with the provided details
// @Tags movies
// @Accept json
// @Produce json
// @Param movie body models.CreateMovieRequest true "Movie details"
// @Success 201 {object} models.CreateMovieRequest
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /movies [post]
func (h *MovieHandler) CreateMovie(c *gin.Context) {
	var request models.Movie
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.movieService.CreateMovie(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "movie created successfully")
}

// UpdateMovie godoc
// @Summary Update a movie
// @Description Update an existing movie's details
// @Tags movies
// @Accept json
// @Produce json
// @Param id path string true "Movie ID"
// @Param movie body models.UpdateMovieRequest true "Movie details to update"
// @Success 200 {object} models.Movie
// @Failure 400 {object} map[string]string "Invalid request payload or ID"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /movies/{id} [put]
func (h *MovieHandler) UpdateMovie(c *gin.Context) {
	id := c.Param("id")

	// Parse string ID to UUID
	parsedID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	// Fetch the existing movie from the database
	existingMovie, err := h.movieService.GetMovie(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Bind the request JSON
	var request models.UpdateMovieRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Prepare the movie to update, preserving Title if not provided
	movie := models.Movie{
		ID:              parsedID,
		Title:           existingMovie.Title, // Default to existing Title
		Director:        request.Director,
		Year:            request.Year,
		Plot:            request.Plot,
		ImdbRating:      request.ImdbRating,
		ItvRating:       request.ItvRating,
		KinopoiskRating: request.KinopoiskRating,
		Duration:        request.Duration,
		Budget:          request.Budget,
		UpdatedAt:       request.UpdatedAt,
	}

	// Only update Title if provided in the request
	if request.Title != "" {
		movie.Title = request.Title
	}

	// Update the movie in the database
	if err := h.movieService.UpdateMovie(&movie); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}

// DeleteMovie godoc
// @Summary Delete a movie
// @Description Delete a movie by its ID
// @Tags movies
// @Accept json
// @Produce json
// @Param id path string true "Movie ID"
// @Success 200 {object} map[string]string "Success message"
// @Failure 400 {object} map[string]string "Invalid ID format"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /movies/{id} [delete]
func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	if err := h.movieService.DeleteMovie(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}
