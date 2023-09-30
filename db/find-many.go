package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type VisitedLinkFindMany struct {
	Website     string    `bson:"website"`
	Link        string    `bson:"link"`
	VisitedDate time.Time `bson:"visitedDate"`
}

func FindMany() {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)
	c := client.Database("Crawler").Collection(("links"))
	// opts := options.Count().SetLimit(1)
	// n, err := c.CountDocuments(context.TODO(), bson.D{{"link", link}},
	// 	opts)
	_, err := c.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	// _, err := c.Find(context.Background(), data)
}
