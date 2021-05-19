package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DBClient *sqlx.DB

func InitDBConnection() error {
	godotenv.Load(".env")
	dial := os.Getenv("DBDIAL")
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("DBUSER")
	name := os.Getenv("DBNAME")
	pssw := os.Getenv("DBPSSW")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=require password=%s port=%s", host, user, name, pssw, port)
	db, err := sqlx.Open(dial, dbURI)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	DBClient = db
	return nil
}
