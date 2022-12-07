package main

import (
	"cghendrix/nbfacts/db"
	"cghendrix/nbfacts/docs"
	"cghendrix/nbfacts/internal/facts"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

func main() {
	viper.SetConfigFile("./env/.env")
	viper.ReadInConfig()

	dbUrl := viper.Get("DB_URL").(string)
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	engine := setupServerAndRoutes(dbUser + ":" + dbPass + "@" + dbUrl)

	port := viper.Get("PORT").(string)
	err := engine.Run(port)
	if err != nil {
		log.Fatalf("Error starting server %s", err)
		return
	}
}

func setupServerAndRoutes(dbUrl string) *gin.Engine {
	server := gin.Default()
	factsDb := db.Init(dbUrl)

	// Facts
	factsRepository := facts.NewRepository(factsDb)
	factsService := facts.NewService(factsRepository)
	facts.NewHandler(server, factsService)

	// Swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Description = "A robust API for all of your Nickleback facts!"
	docs.SwaggerInfo.Title = "Nickleback Facts API"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return server
}
