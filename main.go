package main

import (
	"context"
	"log"

	"firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

func validateToken(c *gin.Context) {

        ctx := context.Background()
        app, err := firebase.NewApp(ctx, nil)
        if err != nil {
                log.Fatalf("error initializing app: %v\n", err)
        }

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

        idToken := extractToken(c)
        
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

        log.Printf("Verified ID token: %v\n", token)
        
	c.JSON(200, gin.H{
          "X-Hasura-User-Id": token.uid,
          "X-Hasura-Role": "user",
        })
}

func extractToken(c *gin.Context) string {
        authHeader := c.Request.Header["Authorization"][0]
        log.Printf("ID token: %v\n", authHeader)
        return authHeader
}

func main() {
	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"OK",
		})
	})
	r.GET("/auth", validateToken)
	r.Run() // listen and serve on 0.0.0.0:8080
}
