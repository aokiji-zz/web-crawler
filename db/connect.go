package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getConnection() (client *mongo.Client, ctx context.Context) {
	user := getEnv().user
	pass := getEnv().pass
	host := getEnv().host
	port := getEnv().port
	uriTemplate := "mongodb://%s:%s@%s:%s"
	uri := fmt.Sprintf(fmt.Sprintf(uriTemplate, user, pass, host, port))
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return
}
