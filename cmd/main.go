package main

import (
	"itv_movie_app/api/app"
	"itv_movie_app/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	_ "itv_movie_app/docs"

	"go.uber.org/fx"
)

// @title           Movie API
// @version         1.0
// @description     A Movie management API.
// @host           localhost:8080
// @BasePath       /api/v1

func NewGinEngine() *gin.Engine {
	return gin.Default()
}

func NewDB() (*gorm.DB, error) {
	return database.NewDatabase()
}

func main() {

	router := gin.Default()

	// CORS middleware configuration
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	db, err := NewDB()
	if err != nil {
		log.Error(err)
		return
	}

	app := fx.New(
		fx.Provide(
			NewGinEngine,
			NewDB,
		),
		fx.Invoke(func() {
			app.InitRoutes(router, db)
			router.Run(":8080")
		}),
	)

	app.Run()
}
