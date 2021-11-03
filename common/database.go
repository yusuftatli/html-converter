package common

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() * gorm.DB {
	_cnn := os.Getenv("POSTGRE_DATABASE_URL")
	db, err := gorm.Open(postgres.Open(_cnn), & gorm.Config {})
	if err != nil {
			log.Fatal(err)
			return nil
	}
	return db
}

func CloseDBConnection(db * gorm.DB) {
	con, err := db.DB()
	if err != nil {
			log.Fatal(err)
			return
	}
	if err := con.Close();
	err != nil {
			log.Fatal(err)
	}
}