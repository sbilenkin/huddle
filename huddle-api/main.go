package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
)

func init() {
	// Path to .env in root directory
	rootDir, err := filepath.Abs("../")
	if err != nil {
		log.Fatalf("Error getting root directory: %v\n", err)
	}
	envPath := filepath.Join(rootDir, ".env")
	// Load environment variables from .env file
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}

	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8081/auth/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func main() {
	router := gin.Default()

	// Route for root URL
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// Route for Google OAuth callback
	router.GET("/auth/google/callback", handleGoogleCallback)

	// Start the server
	log.Fatal(router.Run(":8081"))
}
