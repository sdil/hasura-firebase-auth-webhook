package main

import (
	"context"
	"log"

	"firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

var ctx = context.Background()
var app, _ = firebase.NewApp(ctx, nil)
var client, _ = app.Auth(ctx)

func validateToken(c *gin.Context) {

        idToken := extractToken(c)

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
                log.Printf("error verifying ID token: %v\n", err)

                c.JSON(400, gin.H{
                "status": "Failed",
                "message": "Failed to verify User ID",
                "error": err.Error(),
                })

                return
	}

        log.Printf("Verified ID token: %v\n", token)
	c.JSON(200, gin.H{
          "X-Hasura-User-Id": token.UID,
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
			"status": "OK",
		})
	})
	r.GET("/auth", validateToken)
	r.Run() // listen and serve on 0.0.0.0:8080
}
