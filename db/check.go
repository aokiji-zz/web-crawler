package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func VisitedLink(link string) bool {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)
	c := client.Database("Crawler").Collection(("links"))
	opts := options.Count().SetLimit(1)
	n, err := c.CountDocuments(context.TODO(), bson.D{{"link", link}},
		opts)
	if err != nil {
		panic(err)
	}
	// _, err := c.Find(context.Background(), data)
	return n > 0
}
