package handlers

import (
	"regexp"
	"strings"

	"github.com/orgs/mdyazilim/html-converter-new/models"
	"github.com/orgs/mdyazilim/html-converter-new/utils"
)

func GetParsedHtml(data string, fileName string, uniqueID string) []models.ExcelDto {
	var _ListHtml[] models.ExcelDto
	_ListHtml = append(_ListHtml, * utils.GetDefaultColumns())
	html := utils.GetDataFromHtml(data, "table tr td table tr td table")
	parsedHtml := strings.Split(html, "<td class=\"BG0 S\" colspan=\"9\">")
	for _, masterRow := range parsedHtml {
			utils.WriteValue(masterRow)
			firstSplit := strings.Split(masterRow, "<td class=\"BG1 S\" colspan=\"16\">")
			for _, firstRow := range firstSplit {
					utils.WriteValue(firstRow)
					if strings.Contains(firstRow, "<td class=\"BG2 S\" colspan=\"15\">") {
							secondSplit := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"15\">")
							for _,
							secondRow := range secondSplit {
									utils.WriteValue(secondRow)
									if strings.Contains(secondRow, "<td class=\"BG0 S\" colspan=\"16\">") {
											splitData2 := strings.Split(secondRow, "<td class=\"BG0 S\" colspan=\"16\">")
											for _,
											bg0Data := range splitData2 {
													utils.WriteValue(bg0Data)
													trSplitdata := strings.Split(bg0Data, "<tr>")
													for _, row := range trSplitdata {
															utils.WriteValue(row)
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
															val := utils.ParseData(row, _index)
															if utils.FieldControl(val) {
																_ListHtml = append(_ListHtml, * val)
															}
													}
											}
									} else if strings.Contains(secondRow, "<td class=\"BG2 S\" colspan=\"14\">") {
											splitData := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"16\">")
											for _,
											bg2Data := range splitData {
													utils.WriteValue(bg2Data)
													trSplitdata := strings.Split(bg2Data, "<tr>")
													for _, row := range trSplitdata {
															utils.WriteValue(row)
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
															val := utils.ParseData(row, _index)
															if utils.FieldControl(val) {
																_ListHtml = append(_ListHtml, * val)
															}
													}
											}

									} else if strings.Contains(secondRow, "<td class=\"BG0 S\" colspan=\"14\">") {
											splitDataThird := strings.Split(secondRow, "<td class=\"BG0 S\" colspan=\"14\">")
											for _,
											thirdData := range splitDataThird {
													utils.WriteValue(thirdData)
													if strings.Contains(thirdData, "<td class=\"BG1 S\" colspan=\"16\">") {
															splitBg016 := strings.Split(thirdData, "<td class=\"BG1 S\" colspan=\"16\">")
															for _,
															bg2Data := range splitBg016 {
																	utils.WriteValue(bg2Data)
																	trSplitdata := strings.Split(bg2Data, "<tr>")
																	for _, row := range trSplitdata {
																			utils.WriteValue(row)
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
																			val := utils.ParseData(row, _index)
																			if utils.FieldControl(val) {
																				_ListHtml = append(_ListHtml, * val)
																			}
																	}
															}
													} else if strings.Contains(firstRow, "<td class=\"BG2 S\" colspan=\"9\">") {
															splitbgos15 := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"9\">")
															for _,
															bgorow := range splitbgos15 {
																	trSplitdata := strings.Split(bgorow, "<tr>")
																	for _,
																	row := range trSplitdata {
																			utils.WriteValue(row)
																			_index := "0"
																			if strings.Contains(row, "class=\"BG0") {
																					_index = "2"
																			} else if strings.Contains(row, "class=\"BG1") {
																					_index = "4"
																			} else if strings.Contains(row, "class=\"BG2\">") {
																					_index = "5"
																			} else if strings.Contains(row, "class=\"BG2 S") {
																					_index = "2"
																			} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
																					_index = "3"
																			}
																			val := utils.ParseData(row, _index)
																			if utils.FieldControl(val) {
																					_ListHtml = append(_ListHtml, * val)
																			}
																	}
															}
													} else {
															trSplitdata := strings.Split(thirdData, "<tr>")
															for _,
															row := range trSplitdata {
																	utils.WriteValue(row)
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
																	val := utils.ParseData(row, _index)
																	if utils.FieldControl(val) {
																			_ListHtml = append(_ListHtml, * val)
																	}
															}
													}
											}
									} else if strings.Contains(secondRow, "<td class=\"BG0 S\" colspan=\"15\">") {
											splitbgos15 := strings.Split(secondRow, "<td class=\"BG0 S\" colspan=\"15\">")
											for _,
											bgorow := range splitbgos15 {
													utils.WriteValue(bgorow)
													if strings.Contains(bgorow, "<td class=\"BG1 S\" colspan=\"16\">") {
															splitBg016 := strings.Split(bgorow, "<td class=\"BG1 S\" colspan=\"16\">")
															for _,
															bg2Data := range splitBg016 {
																	utils.WriteValue(bg2Data)
																	trSplitdata := strings.Split(bg2Data, "<tr>")
																	for _, row := range trSplitdata {
																			utils.WriteValue(row)
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
																			val := utils.ParseData(row, _index)
																			if utils.FieldControl(val) {
																					_ListHtml = append(_ListHtml, * val)
																			}
																	}
															}
													} else {
															trSplitdata := strings.Split(bgorow, "<tr>")
															for _,
															row := range trSplitdata {
																	utils.WriteValue(row)
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
																	val := utils.ParseData(row, _index)
																	if utils.FieldControl(val) {
																			_ListHtml = append(_ListHtml, * val)
																	}
															}
													}
											}
									} else {
											trSplitdata := strings.Split(secondRow, "<tr>")
											for _,
											row := range trSplitdata {
													utils.WriteValue(row)
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
													val := utils.ParseData(row, _index)
													if utils.FieldControl(val) {
															_ListHtml = append(_ListHtml, * val)
													}
											}
									}
							}
					} else if strings.Contains(firstRow, "<td class=\"BG2 S\" colspan=\"16\">") {
							splitData := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"16\">")
							for _,
							bg2Data := range splitData {
									utils.WriteValue(bg2Data)
									trSplitdata := strings.Split(bg2Data, "<tr>")
									for _, row := range trSplitdata {
											utils.WriteValue(row)
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
											val := utils.ParseData(row, _index)
											if utils.FieldControl(val) {
													_ListHtml = append(_ListHtml, * val)
											}
									}
							}
					} else if strings.Contains(firstRow, "<td class=\"BG2 S\" colspan=\"14\">") {
							splitData := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"16\">")
							for _,
							bg2Data := range splitData {
									utils.WriteValue(bg2Data)
									trSplitdata := strings.Split(bg2Data, "<tr>")
									for _, row := range trSplitdata {
											utils.WriteValue(row)
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
											val := utils.ParseData(row, _index)
											if utils.FieldControl(val) {
													_ListHtml = append(_ListHtml, * val)
											}
									}
							}
					} else if strings.Contains(firstRow, "<td class=\"BG2 S\" colspan=\"13\">") {
							splitData := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"13\">")
							for _,
							bg2Data := range splitData {
									utils.WriteValue(bg2Data)
									trSplitdata := strings.Split(bg2Data, "<tr>")
									for _, row := range trSplitdata {
											utils.WriteValue(row)
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
											val := utils.ParseData(row, _index)
											if utils.FieldControl(val) {
													_ListHtml = append(_ListHtml, * val)
											}
									}
							}
					} else if strings.Contains(firstRow, "<td class=\"BG2 S\" colspan=\"17\">") {
							splitData := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"17\">")
							for _,
							bg2Data := range splitData {
									utils.WriteValue(bg2Data)
									trSplitdata := strings.Split(bg2Data, "<tr>")
									for _, row := range trSplitdata {
											utils.WriteValue(row)
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
											val := utils.ParseData(row, _index)
											if utils.FieldControl(val) {
													_ListHtml = append(_ListHtml, * val)
											}
									}
							}
					} else if strings.Contains(firstRow, "<td class=\"BG0 S\" colspan=\"15\">") {
							splitbgos15 := strings.Split(firstRow, "<td class=\"BG0 S\" colspan=\"15\">")
							for _,
							bgorow := range splitbgos15 {
									utils.WriteValue(bgorow)
									if strings.Contains(bgorow, "<td class=\"BG1 S\" colspan=\"16\">") {
											splitBg016 := strings.Split(bgorow, "<td class=\"BG1 S\" colspan=\"16\">")
											for _,
											bg2Data := range splitBg016 {
													utils.WriteValue(bg2Data)
													trSplitdata := strings.Split(bg2Data, "<tr>")
													for _, row := range trSplitdata {
															utils.WriteValue(row)
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
															val := utils.ParseData(row, _index)
															if utils.FieldControl(val) {
																	_ListHtml = append(_ListHtml, * val)
															}
													}
											}
									} else {
											trSplitdata := strings.Split(bgorow, "<tr>")
											for _,
											row := range trSplitdata {
													utils.WriteValue(row)
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
													val := utils.ParseData(row, _index)
													if utils.FieldControl(val) {
															_ListHtml = append(_ListHtml, * val)
													}
											}
									}
							}
					} else if strings.Contains(firstRow, "<td class=\"BG2 S\" colspan=\"9\">") {
							splitbgos15 := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"9\">")
							for _,
							bgorow := range splitbgos15 {
									trSplitdata := strings.Split(bgorow, "<tr>")
									for _,
									row := range trSplitdata {
											utils.WriteValue(row)
											_index := "0"
											if strings.Contains(row, "class=\"BG0") {
													_index = "2"
											} else if strings.Contains(row, "class=\"BG1") {
													_index = "4"
											} else if strings.Contains(row, "class=\"BG2\">") {
													_index = "5"
											} else if strings.Contains(row, "class=\"BG2 S") {
													_index = "2"
											} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
													_index = "3"
											}
											val := utils.ParseData(row, _index)
											if utils.FieldControl(val) {
													_ListHtml = append(_ListHtml, * val)
											}
									}
							}
					} else {
							trSplitdata := strings.Split(firstRow, "<tr>")
							for _,
							row := range trSplitdata {
									utils.WriteValue(row)
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
									val := utils.ParseData(row, _index)
									if utils.FieldControl(val) {
											_ListHtml = append(_ListHtml, * val)
									}
							}
					}
			}
	}
	return _ListHtml
}

func ConverToMasterDto(orders models.ExcelDto, uniqueID string, _packageCreataDate string, _packageNumber string, fileName string) *models.MasterProduct {
	result := models.MasterProduct{}
	result.Tipi = orders.Col1
	result.Uy = orders.Col3
	result.Processid = uniqueID
	result.MasterKod = orders.Col4
	result.Aciklama = orders.Col5
	result.Tedarikci = orders.Col6
	result.Fiyat = orders.Col7
	result.Parabirimi = orders.Col8
	result.FileName = fileName
  result.PackageCretaDate = _packageCreataDate
  result.PackageNumber = _packageNumber
	return &result
}

func ConverToDetailDto(orders models.ExcelDto, uniqueID string, masterid uint64, detayid uint64, header models.ExcelDto, masterCode string) *models.DetailProduct {
	insertType := utils.GetInsertType(orders)
	utils.WriteValue(insertType)
	utils.WriteValue(header.Col5)
	result := models.DetailProduct{}
	result.MasteriId = masterid
	result.Processid = uniqueID
	result.MasterKod = masterCode
	result.DetayId = detayid
	result.Tipi = orders.Col1
	if insertType == "none" {
		result.Malzeme = orders.Col3
		result.Aciklama = orders.Col4
		result.Sarf = orders.Col5
		//-------col6-----------------------
		if header.Col6 == "Fire" {
			result.Fire = orders.Col6
		}
		if header.Col6 == "Brm" {
			result.Brm = orders.Col6
		}
		//-------col7-----------------------
		if header.Col7 == "Brm" {
			result.Brm = orders.Col7
		}
		if header.Col7 == "MMO%" {
			result.MMO = orders.Col7
		}
		//-------col8-----------------------
		if header.Col8 == "MMO%" {
			result.MMO = orders.Col8
		}
		if header.Col8 == "PMO%" {
			result.PMO = orders.Col8
		}
		//-------col9-----------------------
		if header.Col9 == "PMO%" {
			result.PMO = orders.Col9
		}
		if header.Col9 == "PFO%" {
			result.PFO = orders.Col9
		}
		//-------col10-----------------------
		if header.Col10 == "PFO%" {
			result.PFO = orders.Col10
		}
		if header.Col10 == "Esk.Formülü" {
			result.EskFormülü = orders.Col10
		}
	} else if insertType == "Proses" {
		result.Aciklama = orders.Col2
		result.EskFormülü = orders.Col3
		result.Maliyetsa = orders.Col4
		result.OpSuresi = orders.Col5
		result.ÇevrimAdet = orders.Col6
		result.ProsNet = orders.Col7
		result.Prosfile = orders.Col8
		result.Pros = orders.Col9
	} else {
		result.Aciklama = orders.Col4
	}
	if insertType == "op" {
		result.PMO = orders.Col9
		result.PFO = orders.Col10
		result.FOT = orders.Col12
		result.Pros = orders.Col13
		result.MOT = orders.Col14
		result.Fiyat = orders.Col15
	}
	if header.Col5 == "Tedarikçi" {
		result.Tipi = orders.Col2
		result.Malzeme = orders.Col3
		result.Aciklama = orders.Col4
		result.Tedarikci = orders.Col5
		result.Fiyat = orders.Col6
	}
	if insertType == "none" {
		result.Malz = orders.Col12
		result.FOT = orders.Col12
		result.Pros = orders.Col13
		result.MOT = orders.Col14
		result.Fiyat = orders.Col15
	}
	if insertType == "Topl" {
		result.Aciklama = orders.Col1
		result.Fiyat = orders.Col3
	}
	if insertType == "adetay" {
		result.Aciklama = orders.Col4
		result.EskFormülü = orders.Col11
	}
	if insertType == "Topl_4" {
		result.Aciklama = orders.Col4
		result.EskFormülü = orders.Col6
	}
	if insertType == "Ambalaj" {
		result.EskFormülü = strings.ReplaceAll(orders.Col2, "*[TRY]", "")
		result.Fiyat = orders.Col3
	}
    if(header.Col2 == "Esk.Formülü"){
        result.Aciklama = orders.Col1
        result.EskFormülü = orders.Col2
        result.Fiyat = orders.Col3

    }
    if(header.Col3 == "Esk.Formülü"){
        result.Aciklama = orders.Col1
        result.EskFormülü = orders.Col2
        result.Fiyat = orders.Col3

    }
	return &result
}

func GetHeaderData(data string)(string, string) {
	hedarData := utils.GetDataFromHtml(data, "html body table tr td table")
	hedarData = strings.ReplaceAll(hedarData, "</td>", "</td>\n")
	// hedarData = strings.ReplaceAll(hedarData,"<tbody>","")
	// hedarData = strings.ReplaceAll(hedarData,"<td=\"\"","")
	// hedarData = strings.ReplaceAll(hedarData,"<tr>","")
	// hedarData = strings.ReplaceAll(hedarData,"</tr>","")
	// hedarData = strings.ReplaceAll(hedarData,"</tbody>","")
	re := regexp.MustCompile(`<td.*?>(.*)</td>`)
	res := re.FindAllStringSubmatch(hedarData, -1)
	return res[2][1],
	res[8][1]
}