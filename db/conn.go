package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
}

func ConnectDB() (*sql.DB, error) {
	var port int
	fmt.Sscanf(dbPort, "%d", &port)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, port, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("falha ao abrir a conexão: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar com o banco de dados: %w", err)
	}

	return db, nil
}
