package main

import (
	"FS01/database"
	"FS01/server"
	"log"
	"os"
)

func main() {
	if err := database.InitDBConnection(); err != nil {
		log.Fatalln("error connecting to db", err)
	}

	s := server.SetupServer()
	p := os.Getenv("PORT")
	if p == "" {
		p = "8000"
	}
	if err := s.Run(":" + p); err != nil {
		panic(err)
	}
}
