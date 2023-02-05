package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	"github.com/gin-gonic/gin"
)

func main() {
	// initCloudFirestore()
	ginFramework()
}

// Firebase Cloud Firestore
func initCloudFirestore() {
	// Web config

	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("yapin-test-firebase-adminsdk-edjvy-676d58f24c.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	// addDataToFirestore(client, ctx)
	readDataToFirestore(client, ctx)
}

func addDataToFirestore(client *firestore.Client, ctx context.Context) {

	// Add data to the "cities" collection
	_, err := client.Collection("cities").Doc("LA").Set(context.Background(), map[string]interface{}{
		"name":      "Los Angeles",
		"state":     "CA",
		"country":   "USA",
		"createdAt": time.Now(),
	})
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
}

func readDataToFirestore(client *firestore.Client, ctx context.Context) {
	iter := client.Collection("cities").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
}

func ginFramework() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
