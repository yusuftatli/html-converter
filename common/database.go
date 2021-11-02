package common

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() * gorm.DB {
	db, err := gorm.Open(postgres.Open(GetEnvironment().PostgreDatabaseUrl), & gorm.Config {})
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