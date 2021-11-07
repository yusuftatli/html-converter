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


func SaveToDatabase(model models.SaveDatabaseDto) (error) {
	masterDto := models.ConvertMasterDto{}
	detailDto := models.ConvertDetailDto{}
	
	detailDto.UniqueID = model.UniqueID
	detailDto.PackageCreataDate = model.PackageCreataDate
	detailDto.PackageNumber = model.PackageNumber

	masterDto.UniqueID = model.UniqueID
	masterDto.PackageCreataDate = model.PackageCreataDate
	masterDto.PackageNumber = model.PackageNumber 
	masterDto.FileName = model.FileName


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
	for _, masterRow := range model.O {
			utils.WriteValue(masterRow.Col18)
			if masterRow.Col18 == "1_header_"  {
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
					masterDto.E = masterRow 
					master := ConverToMasterDto(masterDto)
					masterId, masterCode, _err	= repository.SaveMaster(master)
					
			} else if masterRow.Col18 == models.Dept_1 {
				detailDto.E = masterRow
				detailDto.Masterid = masterId
				detailDto.Dept = "1"
				detailDto.H = header_1
				detailDto.MasterCode = masterCode
				detaildata := ConverToDetailDto(detailDto)
				repository.SaveDetail(detaildata)
			} else if masterRow.Col18 == models.Dept_2 {
				detailDto.E = masterRow
				detailDto.Masterid = masterId
				detailDto.Dept = "2"
				detailDto.H = header_2
				detailDto.MasterCode = masterCode
				detaildata := ConverToDetailDto(detailDto)
				repository.SaveDetail(detaildata)
			} else if masterRow.Col18 == models.Dept_3 {
				detailDto.E = masterRow
				detailDto.Masterid = masterId
				detailDto.Dept = "3"
				detailDto.H = header_3
				detailDto.MasterCode = masterCode
				detaildata := ConverToDetailDto(detailDto)
				repository.SaveDetail(detaildata)
			} else if masterRow.Col18 == models.Dept_4 {
				detailDto.E = masterRow
				detailDto.Masterid = masterId
				detailDto.Dept = "4"
				detailDto.H = header_4
				detailDto.MasterCode = masterCode
				detaildata := ConverToDetailDto(detailDto)
				repository.SaveDetail(detaildata)
			} else if masterRow.Col18 == models.Dept_5 {
				detailDto.E = masterRow
				detailDto.Masterid = masterId
				detailDto.Dept = "5"
				detailDto.H = header_5
				detailDto.MasterCode = masterCode
				detaildata := ConverToDetailDto(detailDto)
				repository.SaveDetail(detaildata)
			}
	}
	return _err
}