package main

import (
	"context"
	"log"

	"firebase.google.com/go"
	"firebase.google.com/go/auth"
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
          "X-Hasura-User-Id": token.UID,
          "X-Hasura-Role": "user",
        })
}

func extractToken(c *gin.Context) string {
        authHeader := c.Request.Header["Authorization"][0]
        log.Printf("ID token: %v\n", authHeader)
        return authHeader
}

func signUp(c *gin.Context) {
        
        ctx := context.Background()
        app, err := firebase.NewApp(ctx, nil)
        if err != nil {
                log.Fatalf("error initializing app: %v\n", err)
        }

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
        }
        
        params := (&auth.UserToCreate{}).
                Email("user@example.com").
                EmailVerified(false).
                PhoneNumber("+15555550100").
                Password("secretPassword").
                DisplayName("John Doe").
                PhotoURL("http://www.example.com/12345678/photo.png").
                Disabled(false)
        u, err := client.CreateUser(ctx, params)
        if err != nil {
                log.Printf("error creating user: %v\n", err)
                c.JSON(200, gin.H{
                "status": "Failed",
                "message": "Failed to create new user",
                "error": err.Error(),
                })
        }
        log.Printf("Successfully created user: %v\n", u)

        // Add the new user in Hasura DB with GraphQL Mutation

        c.JSON(200, gin.H{
          "status": "OK",
        })
}

func main() {
	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
	r.GET("/auth", validateToken)
	r.GET("/signup", signUp)
	r.Run() // listen and serve on 0.0.0.0:8080
}
