package db

import (
	"context"
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
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root_username:root_password@localhost:27017"))
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
