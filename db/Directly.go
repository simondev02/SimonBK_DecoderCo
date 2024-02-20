package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func ConnectDirectly() (*sql.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	DSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("pgx", DSN)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
		return nil, err
	}

	// Verificar la conexión
	err = db.Ping()
	if err != nil {
		log.Fatal("Error al verificar la conexión a la base de datos:", err)
		return nil, err
	}

	log.Println("Conexión directa a DB establecida")
	return db, nil
}
