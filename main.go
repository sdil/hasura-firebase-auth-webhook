package main

import (
	"context"
	"log"

	"firebase.google.com/go"
	// "firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

func ValidateToken(c *gin.Context) {

        ctx := context.Background()
        app, err := firebase.NewApp(ctx, nil)
        if err != nil {
                log.Fatalf("error initializing app: %v\n", err)
        }

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	idToken := c.Request.Header["Authorization"][0]
	log.Printf("ID token: %v\n", idToken)

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	log.Printf("Verified ID token: %v\n", token)
	c.JSON(200, gin.H{"token": idToken})
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/auth", ValidateToken)
	r.Run() // listen and serve on 0.0.0.0:8080
}
