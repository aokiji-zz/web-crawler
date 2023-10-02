package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	user string
	pass string
}

func getEnv() User {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env %s", err)
	}
	user := os.Getenv("USR")
	pass := os.Getenv("PASS")
	admin := User{
		user: user,
		pass: pass,
	}
	return admin
}
func getConnection() (client *mongo.Client, ctx context.Context) {
	user := getEnv().user
	pass := getEnv().pass
	host := "localhost"
	port := 27017
	uriTemplate := "mongodb://%s:%s@%s:%d"
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
