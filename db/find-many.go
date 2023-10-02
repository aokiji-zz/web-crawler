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
	_, err := c.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
}
