package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func handleGoogleCallback(c *gin.Context) {
	// Extract the authorization code from the query parameters
	code := c.Query("code")
	if code == "" {
		c.String(http.StatusBadRequest, "No code in request")
		return
	}

	// Exchange the authorization code for an access token
	token, err := exchangeCodeForToken(code)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to exchange code for token")
		return
	}

	// Use the access token to get user info
	userInfo, err := getUserInfo(token.AccessToken)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get user info")
		return
	}

	// Respond with user info
	c.JSON(http.StatusOK, userInfo)
}

func exchangeCodeForToken(code string) (*oauth2.Token, error) {
	// Implement the logic to exchange the authorization code for an access token
	// This typically involves making a POST request to the OAuth2 token endpoint
	return &oauth2.Token{}, nil
}

func getUserInfo(accessToken string) (map[string]interface{}, error) {
	// Implement the logic to get user info using the access token
	// This typically involves making a GET request to the user info endpoint
	return map[string]interface{}{"name": "John Doe"}, nil
}

// func handleGoogleCallback(c *gin.Context) {
// 	ctx := context.Background()
// 	token, err := googleOauthConfig.Exchange(ctx, c.Query("code"))
// 	if err != nil {
// 		log.Printf("Failed to exchange token: %v\n", err)
// 		c.Redirect(http.StatusTemporaryRedirect, "/")
// 		return
// 	}

// 	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
// 	if err != nil {
// 		log.Printf("Failed getting user info: %v\n", err)
// 		c.Redirect(http.StatusTemporaryRedirect, "/")
// 		return
// 	}
// 	defer response.Body.Close()

// 	// Handle user info from Google here
// }
