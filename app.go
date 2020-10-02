package main

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"

	"firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

var ctx = context.Background()
var app, _ = firebase.NewApp(ctx, nil)
var client, _ = app.Auth(ctx)

func validateToken(c *gin.Context) {

	idToken, err := extractToken(c)
	if err != nil {
		c.Header("Cache-Control", os.Getenv("CACHE_DURATION"))
		c.JSON(200, gin.H{
			"X-Hasura-Role": "anonymous",
		})
		return
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Printf("error verifying ID token: %v\n", err)
		c.JSON(400, gin.H{
			"status":  "Failed",
			"message": "Failed to verify User ID",
			"error":   err.Error(),
		})
		return
	}

	log.Printf("Verified ID token: %v\n", token)

	c.Header("Cache-Control", os.Getenv("CACHE_DURATION"))
	c.JSON(200, gin.H{
		"X-Hasura-User-Id": token.UID,
		"X-Hasura-Role":    "user",
	})
}

func extractToken(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("authorization")
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) == 2 {
		token := splitToken[1]
		log.Printf("ID token: %v\n", token)
		return token, nil
	}
	log.Printf("Invalid auth header is given")
	return "", errors.New("Invalid token given")
}

func main() {
	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
	r.GET("/auth", validateToken)
	r.Run(":8081") // listen and serve on 0.0.0.0:8081
}
