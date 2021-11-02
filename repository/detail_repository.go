package repository

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/orgs/mdyazilim/html-converter-new/common"
	"github.com/orgs/mdyazilim/html-converter-new/models"
)


func SaveDetail(masterData * models.DetailProduct) (error) {
	var err error
	_prod, _ := strconv.ParseBool(common.GetEnvironment().Production)
	if (!_prod) {
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
	db, err := sql.Open("goracle", common.GetEnvironment().OracleDatabaseUrl)
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