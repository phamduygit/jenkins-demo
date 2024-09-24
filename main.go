package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/phamduygit/jenkins-demo/domain"
	"github.com/phamduygit/jenkins-demo/service"
	"github.com/phamduygit/jenkins-demo/usecase"
)

func main() {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Health is OK!!")
	})

	sumService := service.NewSumService()
	sumHandler := usecase.NewSumHandler(sumService)

	// POST /sum endpoint to sum two numbers
	e.POST("/sum", func(c echo.Context) error {
		// Parse the request body into SumRequest struct
		sumRequest := new(domain.SumNumberRequest)
		if err := c.Bind(sumRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		}

		res, err := sumHandler.Sum(sumRequest)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
		}

		// Return the sum in the response
		return c.JSON(http.StatusOK, res)
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		log.Fatal("Not found PORT")
	}


	e.Logger.Fatal(e.Start(":" + httpPort))
}
