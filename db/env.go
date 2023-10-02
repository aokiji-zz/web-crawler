package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type User struct {
	user string
	pass string
	host string
	port string
}

func getEnv() User {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env %s", err)
	}
	user := os.Getenv("USR")
	pass := os.Getenv("PASS")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	admin := User{
		user: user,
		pass: pass,
		host: host,
		port: port,
	}
	return admin
}
