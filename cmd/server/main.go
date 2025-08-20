package main

import (
	"fmt"
	"log"
	"mnp-tests-server/internal/db"

	"github.com/joho/godotenv"
)

func main() {
	// загружаем .env из текущей рабочей директории
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := db.LoadConfig()
	fmt.Printf("%+v\n", cfg) // проверка, что env читается

	database := db.Connect(cfg)
	defer func() {
		if err := database.Close(); err != nil {
			log.Printf("Error closing DB: %v", err)
		}
	}()

	fmt.Println("DB connected successfully!")
}
