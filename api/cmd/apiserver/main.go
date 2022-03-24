package main

import (
	"api/internal/kernel/app"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := app.Init()
	if err != nil {
		return
	}
}
