package gomongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InnitConnection() {
	var err error
	ctx := context.Background()
	if client, err = mongo.Connect(ctx, options.Client().ApplyURI(dbAddress)); err != nil {
		log.Fatal(err)
	}

	log.Print("MongoDB Create Client")
}

// ClientClose close the client
func ClientClose() {
	ctx, cancel := context.WithTimeout(context.Background(), 32*time.Second)
	defer cancel()
	if err := client.Disconnect(ctx); err != nil {
		log.Printf("Err at Close the mongo client %s", err.Error())
	}

	log.Println("goodbye MongoDB")
}
