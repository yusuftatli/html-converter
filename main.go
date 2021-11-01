package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/postgres"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/PuerkitoBio/goquery"
	"github.com/google/uuid"
	_ "gopkg.in/goracle.v2"
	"gorm.io/gorm"
)

var _prod bool = false
var _writeValue = false
var withHedaer = true
var local = true
var mainRoot string = ""
var _conn string = ""

type masterproduct struct {
	Id         uint64    `gorm:"primaryKey"`
	Processid  string    `gorm:"column:processid"`
	MasterKod  string    `gorm:"column:malzemecode"`
	Tipi       string    `gorm:"column:tipi"`
	Uy         string    `gorm:"column:uy"`
	Aciklama   string    `gorm:"column:aciklama"`
	Tedarikci  string    `gorm:"column:tedarikci"`
	Fiyat      string    `gorm:"column:fiyat"`
	Parabirimi string    `gorm:"column:parabirimi"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time ` gorm:"column:updated_at"`
}

type detailproduct struct {
	Id         uint64 `gorm:"primaryKey"`
	Processid  string `gorm:"column:processid"`
	MasteriId  uint64 `gorm:"column:masterid"`
	MasterKod  string `gorm:"column:masterkod"`
	DetayId    uint64 `gorm:"column:detayid"`
	Tipi       string `gorm:"column:tipi"`
	Malzeme    string `gorm:"column:malzeme"`
	Aciklama   string `gorm:"column:aciklama"`
	Sarf       string `gorm:"column:sarf"`
	Fire       string `gorm:"column:fire"`
	Brm        string `gorm:"column:brm"`
	MMO        string `gorm:"column:mmo"`
	PMO        string `gorm:"column:pmo"`
	PFO        string `gorm:"column:pfo"`
	EskFormülü string `gorm:"column:eskformulu"`
	Malz       string `gorm:"column:aciklama"`
	FOT        string `gorm:"column:fot"`
	Pros       string `gorm:"column:pros"`
	MOT        string `gorm:"column:mot"`
	Fiyat      string `gorm:"column:fiyat"`
	Maliyetsa  string `gorm:"column:maliyetsa"`
	OpSuresi   string `gorm:"column:opsuresi"`
	ÇevrimAdet string `gorm:"column:cevrimadet"`
	ProsNet    string `gorm:"column:prossnet"`
	Prosfile   string `gorm:"column:prosfile"`
	Toplam     string `gorm:"column:toplam"`
	Tedarikci  string `gorm:"column:tedarikci"`
}

type OrderDto struct {
	col0,
	col1,
	col2,
	col3,
	col4,
	col5,
	col6,
	col7,
	col8,
	col9,
	col10,
	col11,
	col12,
	col13,
	col14,
	col15,
	col16,
	col17,
	col18 string
}


func main() {
	if local {
		mainRoot = "./"
	} else {
		mainRoot = "/Users/tcytatli/Documents/md/prod/"
	}

	files, err := ioutil.ReadDir(mainRoot + "pages")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// fmt.Println(file.Name(), file.IsDir())
		var path string
		var path2 string
		path = mainRoot + "pages/"
		path2 = path + file.Name()
		htmlBytes, err := ioutil.ReadFile(path2)
		if err != nil {
			fmt.Println("error parsing file")
			panic(err)
		}
		// fmt.Println(htmlBytes)
		// fmt.Println(string(htmlBytes))
		// parsedHTML := string(htmlBytes)

		// fmt.Println(parsedHTML)
		// content, _ := ioutil.ReadFile(path2)
		if strings.Contains(file.Name(), "ht") {
			fmt.Println("file==>", file.Name())
			newid := uuid.New()
			goGet3(string(htmlBytes), file.Name(), newid.String())
			fmt.Println(file.Name() + " => " + "Okundu")
			//   moveFile("./pages/"+fileName, "./readedPage/"+fileName)
		}
	}
	// fmt.Println("job started")
	// ticker := time.NewTicker(time.Second)
	//     for t := range ticker.C {
	//         fmt.Println("Read Date ", t)
	//         setTime()
	// }
}

// func setTime(){
//     files, err := ioutil.ReadDir("./pages")
//     if err != nil {
//         log.Fatal(err)
//     }

//     for _, file := range files {
//         fmt.Println(file.Name(), file.IsDir())
//         var path string
//         var path2 string
//           path = "./pages/"
//           path2 = path +file.Name()
//         content, _ := ioutil.ReadFile(path2)
//         goGet(string(content))

//     }
// }

func goGet3(data string, fileName string, uniqueID string) {
	var orders []OrderDto
	columns := OrderDto{}
	columns.col0 = "col0"
	columns.col1 = "col1"
	columns.col2 = "col2"
	columns.col3 = "col3"
	columns.col4 = "col4"
	columns.col5 = "col5"
	columns.col6 = "col6"
	columns.col7 = "col7"
	columns.col8 = "col8"
	columns.col9 = "col9"
	columns.col10 = "col10"
	columns.col11 = "col11"
	columns.col12 = "col12"
	columns.col13 = "col13"
	columns.col14 = "col14"
	columns.col15 = "col15"
	columns.col16 = "col16"
	columns.col17 = "col17"
	columns.col18 = "col18"

	orders = append(orders, *&columns)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(data))
	s := doc.Find("table tr td table tr td table")
	html, _ := s.Html()
	parsedHtml := strings.Split(html, "<td class=\"BG0 S\" colspan=\"9\">")
	for _, masterRow := range parsedHtml {
		writeValue(masterRow)
		firstSplit := strings.Split(masterRow, "td class=\"BG1 S\" colspan=\"16\"")
		for _, firstRow := range firstSplit {
			writeValue(firstRow)
			if strings.Contains(firstRow, "<td class=\"BG2 S\" colspan=\"15\">") {
				secondSplit := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"15\">")
				for _, secondRow := range secondSplit {
					writeValue(secondRow)
					if strings.Contains(secondRow, "<td class=\"BG0 S\" colspan=\"16\">") {
						splitData2 := strings.Split(secondRow, "<td class=\"BG0 S\" colspan=\"16\">")
						for _, bg0Data := range splitData2 {
							writeValue(bg0Data)
							trSplitdata := strings.Split(bg0Data, "<tr>")
							for _, row := range trSplitdata {
								writeValue(row)
								_index := "0"
								if strings.Contains(row, "class=\"BG0") {
									_index = "3"
								} else if strings.Contains(row, "class=\"BG1") {
									_index = "4"
								} else if strings.Contains(row, "class=\"BG2") {
									_index = "2"
								} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
									_index = "3"
								}
								val := parseData(row, _index, withHedaer)
								if fieldControl(val) {
									orders = append(orders, *val)
								}
							}
						}
					} else if strings.Contains(secondRow, "<td class=\"BG2 S\" colspan=\"14\">") {
						splitData := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"16\">")
						for _, bg2Data := range splitData {
							writeValue(bg2Data)
							trSplitdata := strings.Split(bg2Data, "<tr>")
							for _, row := range trSplitdata {
								writeValue(row)
								_index := "0"
								if strings.Contains(row, "class=\"BG0") {
									_index = "3"
								} else if strings.Contains(row, "class=\"BG1") {
									_index = "1"
								} else if strings.Contains(row, "class=\"BG2") {
									_index = "2"
								} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
									_index = "3"
								}
								val := parseData(row, _index, withHedaer)
								if fieldControl(val) {
									orders = append(orders, *val)
								}
							}
						}

					} else if strings.Contains(secondRow, "<td class=\"BG0 S\" colspan=\"14\">") {
						splitDataThird := strings.Split(secondRow, "<td class=\"BG0 S\" colspan=\"14\">")
						for _, thirdData := range splitDataThird {
							writeValue(thirdData)
							if strings.Contains(thirdData, "<td class=\"BG1 S\" colspan=\"16\">") {
								splitBg016 := strings.Split(thirdData, "<td class=\"BG1 S\" colspan=\"16\">")
								for _, bg2Data := range splitBg016 {
									writeValue(bg2Data)
									trSplitdata := strings.Split(bg2Data, "<tr>")
									for _, row := range trSplitdata {
										writeValue(row)
										_index := "0"
										if strings.Contains(row, "class=\"BG0") {
											_index = "3"
										} else if strings.Contains(row, "class=\"BG1") {
											_index = "4"
										} else if strings.Contains(row, "class=\"BG2") {
											_index = "5"
										} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
											_index = "3"
										}
										val := parseData(row, _index, withHedaer)
										if fieldControl(val) {
											orders = append(orders, *val)
										}
									}
								}
							} else {
								trSplitdata := strings.Split(thirdData, "<tr>")
								for _, row := range trSplitdata {
									writeValue(row)
									_index := "0"
									if strings.Contains(row, "class=\"BG0") {
										_index = "3"
									} else if strings.Contains(row, "class=\"BG1") {
										_index = "4"
									} else if strings.Contains(row, "class=\"BG2\">") {
										_index = "5"
									} else if strings.Contains(row, "class=\"BG2 S") {
										_index = "2"
									} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
										_index = "3"
									}
									val := parseData(row, _index, withHedaer)
									if fieldControl(val) {
										orders = append(orders, *val)
									}
								}
							}
						}
					} else {
						trSplitdata := strings.Split(secondRow, "<tr>")
						for _, row := range trSplitdata {
							writeValue(row)
							_index := "0"
							if strings.Contains(row, "class=\"BG0") {
								_index = "3"
							} else if strings.Contains(row, "class=\"BG1") {
								_index = "1"
							} else if strings.Contains(row, "class=\"BG2") {
								_index = "2"
							} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
								_index = "3"
							}
							val := parseData(row, _index, withHedaer)
							if fieldControl(val) {
								orders = append(orders, *val)
							}
						}
					}
				}
			} else if strings.Contains(firstRow, "<td class=\"BG2 S\" colspan=\"16\">") {
				splitData := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"16\">")
				for _, bg2Data := range splitData {
					writeValue(bg2Data)
					trSplitdata := strings.Split(bg2Data, "<tr>")
					for _, row := range trSplitdata {
						writeValue(row)
						_index := "0"
						if strings.Contains(row, "class=\"BG0") {
							_index = "3"
						} else if strings.Contains(row, "class=\"BG1") {
							_index = "1"
						} else if strings.Contains(row, "class=\"BG2") {
							_index = "2"
						} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
							_index = "3"
						}
						val := parseData(row, _index, withHedaer)
						if fieldControl(val) {
							orders = append(orders, *val)
						}
					}
				}
			} else if strings.Contains(firstRow, "<td class=\"BG2 S\" colspan=\"14\">") {
				splitData := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"16\">")
				for _, bg2Data := range splitData {
					writeValue(bg2Data)
					trSplitdata := strings.Split(bg2Data, "<tr>")
					for _, row := range trSplitdata {
						writeValue(row)
						_index := "0"
						if strings.Contains(row, "class=\"BG0") {
							_index = "2"
						} else if strings.Contains(row, "class=\"BG1") {
							_index = "1"
						} else if strings.Contains(row, "class=\"BG2") {
							_index = "2"
						} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
							_index = "3"
						}
						val := parseData(row, _index, withHedaer)
						if fieldControl(val) {
							orders = append(orders, *val)
						}
					}
				}
			} else {
				trSplitdata := strings.Split(firstRow, "<tr>")
				for _, row := range trSplitdata {
					writeValue(row)
					_index := "0"
					if strings.Contains(row, "class=\"BG0") {
						_index = "master"
					} else if strings.Contains(row, "class=\"BG1") {
						_index = "1"
					} else if strings.Contains(row, "class=\"BG2") {
						_index = "2"
					} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
						_index = "3"
					}

					val := parseData(row, _index, withHedaer)
					if fieldControl(val) {
						orders = append(orders, *val)
					}
				}
			}
		}
	}
	saveTabels(orders, uniqueID)
	writeToExcel(orders, fileName)
}

func fieldControl(row *OrderDto) bool {
	if row.col0 != "" {
		return true
	} else {
		return false
	}
}

func writeValue(value string) {
	if _writeValue {
		fmt.Println("-------------------------------------")
		fmt.Println(value)
	}
}

func parseData(value string, dept string, header bool) *OrderDto {
	e := OrderDto{}
	if strings.Contains(value, "<tbody>") && strings.Contains(value, "<table>") {
		e.col0 = ""
	} else {
		value = strings.ReplaceAll(value, "</td>", "</td>\n")
		//   value = strings.ReplaceAll(value,"</tr>","")
		//  value = strings.ReplaceAll(value,"class=\"BG0\"","")
		//  value = strings.ReplaceAll(value,"class=\"BG0\"","")
		//  value = strings.ReplaceAll(value,"class=\"BG0 L MLZ\"","")
		//  value = strings.ReplaceAll(value,"\u00a0</td>","")
		//  value = strings.ReplaceAll(value,"<td >","<td>")
		//  value = strings.ReplaceAll(value,"<td></td>","<td>-</td>")
		//  value = strings.ReplaceAll(value,"class=\"BG0 R FP1\"","")
		re := regexp.MustCompile(`<td.*?>(.*)</td>`)
		e.col18 = dept
		if header && strings.Contains(value, "th") {
			value = strings.ReplaceAll(value, "</th>", "</th>\n")
			re = regexp.MustCompile(`<th.*?>(.*)</th>`)
			e.col18 = dept + "_header_"
		}
		submatchall := re.FindAllStringSubmatch(value, -1)
		for i, element := range submatchall {
			switch i {
			case 0:
				e.col0 = isCheckSpace(element[1])
			case 1:
				e.col1 = isCheckSpace(element[1])
			case 2:
				e.col2 = isCheckSpace(element[1])
			case 3:
				e.col3 = isCheckSpace(element[1])
			case 4:
				e.col4 = isCheckSpace(element[1])
			case 5:
				e.col5 = isCheckSpace(element[1])
			case 6:
				e.col6 = isCheckSpace(element[1])
			case 7:
				e.col7 = isCheckSpace(element[1])
			case 8:
				e.col8 = isCheckSpace(element[1])
			case 9:
				e.col9 = isCheckSpace(element[1])
			case 10:
				e.col10 = isCheckSpace(element[1])
			case 11:
				e.col11 = isCheckSpace(element[1])
			case 12:
				e.col12 = isCheckSpace(element[1])
			case 13:
				e.col13 = isCheckSpace(element[1])
			case 14:
				e.col14 = isCheckSpace(element[1])
			case 15:
				e.col15 = isCheckSpace(element[1])
			case 16:
				e.col16 = isCheckSpace(element[1])
			case 17:
				e.col17 = isCheckSpace(element[1])
			}
		}
	}
	return &e
}

func isCheckSpace(value string) string {
	if value == "&nbsp;" {
		return "-"
	} else {
		return value
	}
}

func writeToExcel(orders []OrderDto, fileName string) {

	f := excelize.NewFile()
	index := f.NewSheet("Sheet1")

	count := 1
	for _, element := range orders {

		f.SetCellValue("Sheet1", "A"+strconv.Itoa(count), element.col18)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(count), element.col0)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(count), element.col1)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(count), element.col2)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(count), element.col3)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(count), element.col4)
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(count), element.col5)
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(count), element.col6)
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(count), element.col7)
		f.SetCellValue("Sheet1", "J"+strconv.Itoa(count), element.col8)
		f.SetCellValue("Sheet1", "K"+strconv.Itoa(count), element.col9)
		f.SetCellValue("Sheet1", "L"+strconv.Itoa(count), element.col10)
		f.SetCellValue("Sheet1", "M"+strconv.Itoa(count), element.col11)
		f.SetCellValue("Sheet1", "N"+strconv.Itoa(count), element.col12)
		f.SetCellValue("Sheet1", "O"+strconv.Itoa(count), element.col13)
		f.SetCellValue("Sheet1", "P"+strconv.Itoa(count), element.col14)
		f.SetCellValue("Sheet1", "Q"+strconv.Itoa(count), element.col15)
		f.SetCellValue("Sheet1", "R"+strconv.Itoa(count), element.col16)
		f.SetCellValue("Sheet1", "S"+strconv.Itoa(count), element.col17)
		count++
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs(mainRoot + "exels/" + fileName + ".xlsx"); err != nil {
		fmt.Println(err)
	}

}

func moveFile(file1, file2 string) error {
	return os.Rename(file1, file2)
}

 func ConnectDB() * gorm.DB {
    //  log.Println("Database bağlantırı kuruluyor...")
     db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=12345 dbname=htmlconvert port=5435 sslmode=disable"), &gorm.Config{})
     if err != nil {
         log.Fatal(err)
         return nil
     }
    //  log.Println("Database Bağlantısı kuruldu.")
     return db
 }

func CloseDBConnection(db *gorm.DB) {
	con, err := db.DB()
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := con.Close(); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(con.Ping())
}

func saveTabels(orders []OrderDto, uniqueID string) {
    if(!_prod){
	    db := ConnectDB()
	    _ = db.AutoMigrate( & masterproduct {})
	    _ = db.AutoMigrate( & detailproduct {})
    }
	var header_1 OrderDto
	var header_2 OrderDto
	var header_3 OrderDto
	var header_4 OrderDto
	var header_5 OrderDto
	var masterId uint64 = 0
	var masterCode string = ""
	for _, masterRow := range orders {
		writeValue(masterRow.col18)
		if masterRow.col18 == "1_header_" {
			header_1 = masterRow
		} else if masterRow.col18 == "2_header_" {
			header_2 = masterRow
		} else if masterRow.col18 == "3_header_" {
			header_3 = masterRow
		} else if masterRow.col18 == "4_header_" {
			header_4 = masterRow
		} else if masterRow.col18 == "5_header_" {
			header_5 = masterRow
		} else if masterRow.col18 == "master" {
			if strings.Contains(masterRow.col4, "Op.Plan") {
				fmt.Println(masterRow)
			}
			masterId = 0
			masterCode = ""
			master := converToMasterDto(masterRow, uniqueID)
            if(_prod){
                masterId = CrateMaster(master)
                masterId = master.Id
                masterCode = master.MasterKod
            }else{
                db := ConnectDB()
                err := db.Create(&master).Error
                masterId = master.Id
                masterCode = master.MasterKod
                if err != nil {
                    log.Fatalf("Unable to execute the query. %v", err)
                }
                CloseDBConnection(db)
            }
		} else if masterRow.col18 == "1" {
			detaildata := converToDetailDto_dept_1(masterRow, uniqueID, masterId, 1, header_1, masterCode)
            WriteDetailData(detaildata)
		} else if masterRow.col18 == "2" {
			detaildata := converToDetailDto_dept_1(masterRow, uniqueID, masterId, 2, header_2, masterCode)
            WriteDetailData(detaildata)
		} else if masterRow.col18 == "3" {
			detaildata := converToDetailDto_dept_1(masterRow, uniqueID, masterId, 3, header_3, masterCode)
            WriteDetailData(detaildata)
		} else if masterRow.col18 == "4" {
			detaildata := converToDetailDto_dept_1(masterRow, uniqueID, masterId, 4, header_4, masterCode)
            WriteDetailData(detaildata)
		} else if masterRow.col18 == "5" {
			detaildata := converToDetailDto_dept_1(masterRow, uniqueID, masterId, 5, header_5, masterCode)
            WriteDetailData(detaildata)
		}
	}
}
func GetInsertType(masterRow OrderDto) string {
	if strings.Contains(masterRow.col4, "Plan") {
		return "op"
	} else if strings.Contains(masterRow.col1, "Proses") {
		return "Proses"
	} else if strings.Contains(masterRow.col1, "Topl") {
		return "Topl"
	} else if strings.Contains(masterRow.col1, "Ambalaj") {
		return "Ambalaj"
	} else if strings.Contains(masterRow.col5, "Tedar") {
		return "Tedarikci"
	} else if strings.Contains(masterRow.col4, "Ambalaj") {
		return "adetay"
	} else if strings.Contains(masterRow.col4, "Topl") {
		return "Topl_4"
	} else if strings.Contains(masterRow.col4, "Esk") {
		return "Esk"
	} else {
		return "none"
	}
}
func converToMasterDto(orders OrderDto, uniqueID string) *masterproduct {
	result := masterproduct{}
	result.Tipi = orders.col1
	result.Uy = orders.col3
	result.Processid = uniqueID
	result.MasterKod = orders.col4
	result.Aciklama = orders.col5
	result.Tedarikci = orders.col6
	result.Fiyat = orders.col7
	result.Parabirimi = orders.col8
	return &result
}
func converToDetailDto_dept_1(orders OrderDto, uniqueID string, masterid uint64, detayid uint64, header OrderDto, masterCode string) *detailproduct {
	insertType := GetInsertType(orders)
	writeValue(insertType)
    writeValue(header.col5)
	result := detailproduct{}
	result.MasteriId = masterid
	result.Processid = uniqueID
	result.MasterKod = masterCode
	result.DetayId = detayid
	result.Tipi = orders.col1
	if insertType == "none" {
		result.Malzeme = orders.col3
		result.Aciklama = orders.col4
		result.Sarf = orders.col5
		//-------col6-----------------------
		if header.col6 == "Fire" {
			result.Fire = orders.col6
		}
		if header.col6 == "Brm" {
			result.Brm = orders.col6
		}
		//-------col7-----------------------
		if header.col7 == "Brm" {
			result.Brm = orders.col7
		}
		if header.col7 == "MMO%" {
			result.MMO = orders.col7
		}
		//-------col8-----------------------
		if header.col8 == "MMO%" {
			result.MMO = orders.col8
		}
		if header.col8 == "PMO%" {
			result.PMO = orders.col8
		}
		//-------col9-----------------------
		if header.col9 == "PMO%" {
			result.PMO = orders.col9
		}
		if header.col9 == "PFO%" {
			result.PFO = orders.col9
		}
		//-------col10-----------------------
		if header.col10 == "PFO%" {
			result.PFO = orders.col10
		}
		if header.col10 == "Esk.Formülü" {
			result.EskFormülü = orders.col10
		}
	} else if insertType == "Proses" {
		result.Aciklama = orders.col2
		result.EskFormülü = strings.ReplaceAll(orders.col3, "*[TRY]", "")
		result.Maliyetsa = orders.col4
		result.OpSuresi = orders.col5
		result.ÇevrimAdet = orders.col6
		result.ProsNet = orders.col7
		result.Prosfile = orders.col8
		result.Pros = orders.col9
	} else {
		result.Aciklama = orders.col4
	}
	if insertType == "op" {
		result.PMO = orders.col9
		result.PFO = orders.col10
		result.FOT = orders.col12
		result.Pros = orders.col13
		result.MOT = orders.col14
		result.Fiyat = orders.col15
	}
	if header.col5 == "Tedarikçi" {
		result.Tipi = orders.col2
		result.Malzeme = orders.col3
		result.Aciklama = orders.col4
		result.Tedarikci = orders.col5
		result.Fiyat = orders.col6
	}
	if insertType == "none" {
		result.Malz = orders.col12
		result.FOT = orders.col12
		result.Pros = orders.col13
		result.MOT = orders.col14
		result.Fiyat = orders.col15
	}
	if insertType == "Topl" {
		result.Aciklama = orders.col1
		result.Fiyat = orders.col3
	}
	if insertType == "adetay" {
		result.Aciklama = orders.col4
		result.EskFormülü = orders.col11
	}
	if insertType == "Topl_4" {
		result.Aciklama = orders.col4
		result.EskFormülü = orders.col6
	}
	if insertType == "Ambalaj" {
		result.EskFormülü = strings.ReplaceAll(orders.col2, "*[TRY]", "")
		result.Fiyat = orders.col3
	}
	return &result
}

func CrateMaster(master *masterproduct) uint64 {
	db, err := sql.Open("goracle", _conn)
	if err != nil {
		fmt.Println("... DB Setup Failed")
		fmt.Println(err)
	}
	defer db.Close()
	dbQuery := ""

	rows, err := db.Query(dbQuery)
	if err != nil {
		fmt.Println(".....Error processing query")
		fmt.Println(err)
	}
	defer rows.Close()
	return 1
}

func CrateDetail(detail *detailproduct) {
	db, err := sql.Open("goracle", _conn)
	if err != nil {
		fmt.Println("... DB Setup Failed")
		fmt.Println(err)
	}
	defer db.Close()
	dbQuery := ""

	rows, err := db.Query(dbQuery)
	if err != nil {
		fmt.Println(".....Error processing query")
		fmt.Println(err)
	}
	defer rows.Close()
}


func WriteDetailData(detaildata *detailproduct){
    if(_prod){
        CrateDetail(detaildata)
    }else{
        db := ConnectDB()
        err := db.Create(&detaildata).Error
        if err != nil {
            log.Fatalf("Unable to execute the query. %v", err)
        }
        CloseDBConnection(db)
    }
}