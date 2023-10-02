package db

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type VisitedLinkFindMany struct {
	Website     string    `bson:"website"`
	Link        string    `bson:"link"`
	VisitedDate time.Time `bson:"visitedDate"`
}

func FindMany() {
	//var resp []VisitedLinkFindMany
	client, ctx := getConnection()
	defer client.Disconnect(ctx)
	c := client.Database("Crawler").Collection(("links"))
	cursor, err := c.Find(context.TODO(), bson.D{{"website", "aprendagolang.com.br"}})
	if err != nil {
		panic(err)
	}
	var results []VisitedLinkFindMany
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
}
