package repository

import (
	"database/sql"
	"log"
	"os"

	"github.com/orgs/mdyazilim/html-converter-new/common"
	"github.com/orgs/mdyazilim/html-converter-new/models"
)


func SaveDetail(masterData * models.DetailProduct) (error) {
	var err error
	_prod := os.Getenv("PRODUCTION")
	if (_prod == "false") {
			err = savePostgreDetail(masterData)
	} else {
		err =  saveOracleDetail(masterData)
	}
	return err
}

func savePostgreDetail(data * models.DetailProduct) (error) {
	db := common.ConnectDB()
	err := db.Create(&data).Error
	if err != nil {
			log.Fatalf("Unable to execute the query. %v", err)
			return err
	}
	common.CloseDBConnection(db)
	return err
}

func saveOracleDetail(data *models.DetailProduct) ( error) {
	db, err := sql.Open("goracle", os.Getenv("ORACLE_DATABASE_URL"))
	if err != nil {
		return err
	}
	defer db.Close()
	dbQuery := ""

	rows, err := db.Query(dbQuery)
	if err != nil {
		return err
	}
	defer rows.Close()
	return err
}