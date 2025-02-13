package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"short-url-service/internal/handler"
	"short-url-service/internal/storage"
)

func main() {
	// Получаем тип хранилища из флага или переменной окружения
	dbType := flag.String("db", os.Getenv("DB_TYPE"), "Choose storage: 'memory' or 'postgres'")
	flag.Parse()

	var store storage.Storage
	var err error

	if *dbType == "postgres" {
		log.Println("DB_TYPE=PG")
		dsn := os.Getenv("DATABASE_URL")
		if dsn == "" {
			log.Fatal("DATABASE_URL is not set")
		}
		store, err = storage.NewPostgresStorage(dsn)
		if err != nil {
			log.Fatal("Failed to connect to PostgreSQL:", err)
		}
	} else {
		log.Println("DB_TYPE=memory")
		store = storage.NewInMemoryStorage()
	}

	http.HandleFunc("/api/url", handler.PostHandler(store))
	http.HandleFunc("/api/url/", handler.GetHandler(store))

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
