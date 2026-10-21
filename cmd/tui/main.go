package main

import (
	"log"

	"github.com/RubFV/venky/internal/db"
	"github.com/RubFV/venky/internal/logger"
)

func main() {

	f, err := logger.Setup("venky", false)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("error al cerrar el archivo de logs")
		}
	}()

	db, err := db.Start("venkyData.db")
	if err != nil {
		log.Panicf("error al iniciar la base de datos: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("error al cerrar la base de datos: %v", err)
		}
	}()

}
