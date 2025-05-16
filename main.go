package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/api/sheets/v4"
)

func main() {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Warning: .env file not found. Falling back to environment variables.")
	}

	// Initialize GCP services
	var sheetsService = setUpGoogleSheetsAPI()

	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", handleRoot)
	e.GET("/get_networkpoints", func(c echo.Context) error { return handleGetNetworkpoints(c, sheetsService) })
	e.GET("/get_level1", func(c echo.Context) error { return handleGetLevel1(c, sheetsService) })
	e.GET("/get_level2", func(c echo.Context) error { return handleGetLevel2(c, sheetsService) })
	e.GET("/get_level3", func(c echo.Context) error { return handleGetLevel3(c, sheetsService) })

	e.GET("/thisisatest", handleTest)

	// Start server
	port := os.Getenv("PORT")
	address := ":" + port
	e.Logger.Fatal(e.Start(address))
}

func handleTest(c echo.Context) error {
	return c.String(http.StatusOK, "Happy Friday!")
}

// Handlers

func handleRoot(c echo.Context) error {
	return c.String(http.StatusOK, "PCW Network Map GeoJson API")
}

func handleGetNetworkpoints(c echo.Context, sheetsService *sheets.Service) error {
	sheet_values := get_sheet_values(sheetsService, "networkpoints")
	networkpoints := process_networkpoints(sheet_values)
	shelled := prep_for_export(networkpoints)
	return c.JSON(http.StatusOK, shelled)
}

func handleGetLevel1(c echo.Context, sheetsService *sheets.Service) error {
	sheet_values := get_sheet_values(sheetsService, "level1")
	networkpoints := process_level1(sheet_values)
	shelled := prep_for_export_level1(networkpoints)
	return c.JSON(http.StatusOK, shelled)
}

func handleGetLevel2(c echo.Context, sheetsService *sheets.Service) error {
	sheet_values := get_sheet_values(sheetsService, "level2")
	networkpoints := process_level2_level3(sheet_values)
	shelled := prep_for_export_level2_3(networkpoints)
	return c.JSON(http.StatusOK, shelled)
}

func handleGetLevel3(c echo.Context, sheetsService *sheets.Service) error {
	sheet_values := get_sheet_values(sheetsService, "level3")
	networkpoints := process_level2_level3(sheet_values)
	shelled := prep_for_export_level2_3(networkpoints)
	return c.JSON(http.StatusOK, shelled)
}
