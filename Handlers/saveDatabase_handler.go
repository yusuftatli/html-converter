package handlers

import (
	"fmt"
	"os"
	"strings"

	"github.com/orgs/mdyazilim/html-converter-new/common"
	"github.com/orgs/mdyazilim/html-converter-new/models"
	"github.com/orgs/mdyazilim/html-converter-new/repository"
	"github.com/orgs/mdyazilim/html-converter-new/utils"
)


func SaveToDatabase(orders []models.ExcelDto, uniqueID string, _packageCreataDate string, _packageNumber string, fileName string) (error) {
	_prod := os.Getenv("PRODUCTION")
	if (_prod == "false") {
			db := common.ConnectDB()
			_ = db.AutoMigrate( & models.MasterProduct {})
			_ = db.AutoMigrate( & models.DetailProduct {})
	}
	var header_1, header_2, header_3, header_4, header_5 models.ExcelDto
	var masterId uint64 = 0
	var masterCode string = ""
	var _err error
	for _, masterRow := range orders {
			utils.WriteValue(masterRow.Col18)
			if masterRow.Col18 == "1_header_" {
					header_1 = masterRow
			} else if masterRow.Col18 == "2_header_" {
					header_2 = masterRow
			} else if masterRow.Col18 == "3_header_" {
					header_3 = masterRow
			} else if masterRow.Col18 == "4_header_" {
					header_4 = masterRow
			} else if masterRow.Col18 == "5_header_" {
					header_5 = masterRow
			} else if masterRow.Col18 == "master" {
					if strings.Contains(masterRow.Col4, "Op.Plan") {
							fmt.Println(masterRow)
					}
					masterId = 0
					masterCode = ""
					master := ConverToMasterDto(masterRow, uniqueID, _packageCreataDate, _packageNumber, fileName)
					masterId, masterCode, _err	= repository.SaveMaster(master)
					
			} else if masterRow.Col18 == "1" {
					detaildata := ConverToDetailDto(masterRow, uniqueID, masterId, 1, header_1, masterCode)
					repository.SaveDetail(detaildata)
			} else if masterRow.Col18 == "2" {
					detaildata := ConverToDetailDto(masterRow, uniqueID, masterId, 2, header_2, masterCode)
					repository.SaveDetail(detaildata)
			} else if masterRow.Col18 == "3" {
					detaildata := ConverToDetailDto(masterRow, uniqueID, masterId, 3, header_3, masterCode)
					repository.SaveDetail(detaildata)
			} else if masterRow.Col18 == "4" {
					detaildata := ConverToDetailDto(masterRow, uniqueID, masterId, 4, header_4, masterCode)
					repository.SaveDetail(detaildata)
			} else if masterRow.Col18 == "5" {
					detaildata := ConverToDetailDto(masterRow, uniqueID, masterId, 5, header_5, masterCode)
					repository.SaveDetail(detaildata)
			}
	}
	return _err
}