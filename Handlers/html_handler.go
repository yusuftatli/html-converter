package handlers

import (
	"regexp"
	"strings"

	"github.com/orgs/mdyazilim/html-converter-new/models"
	"github.com/orgs/mdyazilim/html-converter-new/utils"
)

func GetParsedHtml(data string, fileName string, uniqueID string, _packageCreataDate string, _packageNumber string)[] models.ExcelDto {
    var _ListHtml[] models.ExcelDto
    _ListHtml = append(_ListHtml, * utils.GetDefaultColumns())
    html := utils.GetDataFromHtml(data, "table tr td table tr td table")
    parsedHtml := strings.Split(html, "<td class=\"BG0 S\" colspan=\"9\">")
    for _, masterRow := range parsedHtml {
        utils.WriteValue(masterRow)
        if strings.Contains(masterRow, "<td class=\"BG1 S\" colspan=\"16\">") {
            firstSplit := strings.Split(masterRow, "<td class=\"BG1 S\" colspan=\"16\">")
            for _,
            firstRow := range firstSplit {
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
                        } else if strings.Contains(secondRow, "<td class=\"BG2 S\" colspan=\"16\">") {
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
                            } else if (strings.Contains(secondRow, "<td class=\"BG0 S\" colspan=\"15\">")) {
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
                                                    _index = "3"
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
                            } else
                        if strings.Contains(secondRow, "<td class=\"BG0 S\" colspan=\"15\">") {
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
                    splitData := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"14\">")
                    for _,
                    bg2Data := range splitData {
                        utils.WriteValue(bg2Data)
                        if strings.Contains(bg2Data, "<td class=\"BG0 S\" colspan=\"16\">") {
                            split_span16 := strings.Split(bg2Data, "<td class=\"BG0 S\" colspan=\"16\">")
                            for _,
                            row_16 := range split_span16 {
                                trSplitdata := strings.Split(row_16, "<tr>")
                                for _,
                                row := range trSplitdata {
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
                        } else {
                            trSplitdata := strings.Split(bg2Data, "<tr>")
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
        } else if (strings.Contains(masterRow, "<td class=\"BG1 S\" colspan=\"15\">")) {
            split_colspan_16 := strings.Split(masterRow, "<td class=\"BG1 S\" colspan=\"15\">")
            for _,
            colspan_firstdata := range split_colspan_16 {
                utils.WriteValue(colspan_firstdata)
                if strings.Contains(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"14\">") {
                    utils.WriteValue(colspan_firstdata)
                    split_colspan_14 := strings.Split(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"14\">")
                    for _,
                    colspan_second := range split_colspan_14 {
                        utils.WriteValue(colspan_second)
                        if strings.Contains(colspan_second, "<td class=\"BG0 S\" colspan=\"16\">") {
                            utils.WriteValue(colspan_second)
                            splitData2 := strings.Split(colspan_second, "<td class=\"BG0 S\" colspan=\"16\">")
                            utils.WriteValue("")
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
                            trSplitdata := strings.Split(colspan_second, "<tr>")
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
                } else if strings.Contains(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"16\">") {
                    splitData := strings.Split(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"16\">")
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
                } else if strings.Contains(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"13\">") {
                    splitData := strings.Split(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"13\">")
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
                } else if strings.Contains(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"17\">") {
                    splitData := strings.Split(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"17\">")
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
                } else if strings.Contains(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"15\">") {
                    splitData := strings.Split(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"15\">")
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
                } else {
                    trSplitdata := strings.Split(colspan_firstdata, "<tr>")
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
        } else if (strings.Contains(masterRow, "<td class=\"BG1 S\" colspan=\"17\">")) {
            split_colspan_17 := strings.Split(masterRow, "<td class=\"BG1 S\" colspan=\"17\">")
            for _,
            colspan_firstdata := range split_colspan_17 {
                utils.WriteValue(colspan_firstdata)
                if strings.Contains(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"14\">") {
                    utils.WriteValue(colspan_firstdata)
                    split_colspan_14 := strings.Split(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"14\">")
                    for _,
                    colspan_second := range split_colspan_14 {
                        utils.WriteValue(colspan_second)
                        if strings.Contains(colspan_second, "<td class=\"BG0 S\" colspan=\"16\">") {
                            utils.WriteValue(colspan_second)
                            splitData2 := strings.Split(colspan_second, "<td class=\"BG0 S\" colspan=\"16\">")
                            utils.WriteValue("")
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
                                    } else if strings.Contains(row, "class=\"BG2\">") {
                                        _index = "2"
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
                        } else if (strings.Contains(colspan_second, "<td class=\"BG2 S\" colspan=\"16\">")) {
                            trSplitdata := strings.Split(colspan_second, "<tr>")
                            for _,
                            row := range trSplitdata {
                                utils.WriteValue(row)
                                _index := "0"
                                if strings.Contains(row, "class=\"BG0") {
                                    _index = "3"
                                } else if strings.Contains(row, "class=\"BG1") {
                                    _index = "4"
                                } else if strings.Contains(row, "class=\"BG2\">") {
                                    _index = "2"
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
                        } else {
                            trSplitdata := strings.Split(colspan_second, "<tr>")
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
                } else if strings.Contains(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"16\">") {
                    splitData := strings.Split(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"16\">")
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
                } else if (strings.Contains(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"17\">")) {
                    splitData := strings.Split(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"17\">")
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
                } else if (strings.Contains(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"15\">")) {
                    utils.WriteValue(colspan_firstdata)
                    if strings.Contains(colspan_firstdata, "<td class=\"BG0 S\" colspan=\"15\">") {
                        utils.WriteValue(colspan_firstdata)
                        splitData2 := strings.Split(colspan_firstdata, "<td class=\"BG0 S\" colspan=\"16\">")
                        utils.WriteValue("")
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
                                } else if strings.Contains(row, "class=\"BG2\">") {
                                    _index = "2"
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
                    trSplitdata := strings.Split(colspan_firstdata, "<tr>")
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
        } else if (strings.Contains(masterRow, "<td class=\"BG1 S\" colspan=\"18\">")) {
            split_colspan_17 := strings.Split(masterRow, "<td class=\"BG1 S\" colspan=\"18\">")
            for _,
            colspan_firstdata := range split_colspan_17 {
                utils.WriteValue(colspan_firstdata)
                if strings.Contains(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"14\">") {
                    utils.WriteValue(colspan_firstdata)
                    split_colspan_14 := strings.Split(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"14\">")
                    for _,
                    colspan_second := range split_colspan_14 {
                        utils.WriteValue(colspan_second)
                        if strings.Contains(colspan_second, "<td class=\"BG0 S\" colspan=\"16\">") {
                            utils.WriteValue(colspan_second)
                            splitData2 := strings.Split(colspan_second, "<td class=\"BG0 S\" colspan=\"16\">")
                            utils.WriteValue("")
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
                                    } else if strings.Contains(row, "class=\"BG2\">") {
                                        _index = "2"
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
                        } else if (strings.Contains(colspan_second, "<td class=\"BG2 S\" colspan=\"16\">")) {
                            trSplitdata := strings.Split(colspan_second, "<tr>")
                            for _,
                            row := range trSplitdata {
                                utils.WriteValue(row)
                                _index := "0"
                                if strings.Contains(row, "class=\"BG0") {
                                    _index = "3"
                                } else if strings.Contains(row, "class=\"BG1") {
                                    _index = "4"
                                } else if strings.Contains(row, "class=\"BG2\">") {
                                    _index = "2"
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
                        } else {
                            trSplitdata := strings.Split(colspan_second, "<tr>")
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
                } else if strings.Contains(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"16\">") {
                        splitData := strings.Split(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"16\">")
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
                    } else if (strings.Contains(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"17\">")) {
                        splitData := strings.Split(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"17\">")
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
                    } else {
                        trSplitdata := strings.Split(colspan_firstdata, "<tr>")
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
        } else if strings.Contains(masterRow, "<td class=\"BG1 S\" colspan=\"14\">") {
                firstSplit := strings.Split(masterRow, "<td class=\"BG1 S\" colspan=\"14\">")
                for _,
                firstRow := range firstSplit {
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
                            } else if strings.Contains(secondRow, "<td class=\"BG2 S\" colspan=\"16\">") {
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
                                } else if (strings.Contains(secondRow, "<td class=\"BG0 S\" colspan=\"15\">")) {
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
                                                        _index = "3"
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
                                } else  if strings.Contains(secondRow, "<td class=\"BG0 S\" colspan=\"15\">") {
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
                        splitData := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"14\">")
                        for _,
                        bg2Data := range splitData {
                            utils.WriteValue(bg2Data)
                            if strings.Contains(bg2Data, "<td class=\"BG0 S\" colspan=\"16\">") {
                                split_span16 := strings.Split(bg2Data, "<td class=\"BG0 S\" colspan=\"16\">")
                                for _,
                                row_16 := range split_span16 {
                                    trSplitdata := strings.Split(row_16, "<tr>")
                                    for _,
                                    row := range trSplitdata {
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
                            } else {
                                trSplitdata := strings.Split(bg2Data, "<tr>")
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
            } else if (strings.Contains(masterRow, "<td class=\"BG1 S\" colspan=\"13\">")) {
                split_colspan_16 := strings.Split(masterRow, "<td class=\"BG1 S\" colspan=\"13\">")
                for _,
                colspan_firstdata := range split_colspan_16 {
                    utils.WriteValue(colspan_firstdata)
                    if strings.Contains(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"14\">") {
                        utils.WriteValue(colspan_firstdata)
                        split_colspan_14 := strings.Split(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"14\">")
                        for _,
                        colspan_second := range split_colspan_14 {
                            utils.WriteValue(colspan_second)
                            if strings.Contains(colspan_second, "<td class=\"BG0 S\" colspan=\"16\">") {
                                utils.WriteValue(colspan_second)
                                splitData2 := strings.Split(colspan_second, "<td class=\"BG0 S\" colspan=\"16\">")
                                utils.WriteValue("")
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
                                trSplitdata := strings.Split(colspan_second, "<tr>")
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
                    } else if strings.Contains(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"16\">") {
                        splitData := strings.Split(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"16\">")
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
                    } else if strings.Contains(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"13\">") {
                        splitData := strings.Split(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"13\">")
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
                    } else if strings.Contains(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"17\">") {
                        splitData := strings.Split(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"17\">")
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
                    } else if strings.Contains(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"15\">") {
                        splitData := strings.Split(colspan_firstdata, "<td class=\"BG2 S\" colspan=\"15\">")
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
                    } else {
                        trSplitdata := strings.Split(colspan_firstdata, "<tr>")
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
            } else {
                trSplitdata := strings.Split(masterRow, "<tr>")
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
    return _ListHtml
}

func ConverToMasterDto(dto models.ConvertMasterDto) * models.MasterProduct {
    result := models.MasterProduct {}
    result.Tipi = dto.E.Col1
    result.Uy = dto.E.Col3
    result.Processid = dto.UniqueID
    result.MasterKod = dto.E.Col4
    result.Aciklama = dto.E.Col5
    result.Tedarikci = dto.E.Col6
    result.Fiyat = dto.E.Col7
    result.Parabirimi = dto.E.Col8
    result.FileName = dto.FileName
    result.PackageCretaDate = dto.PackageCreataDate
    result.PackageNumber = dto.PackageNumber
    return &result
}

func ConverToDetailDto(dto models.ConvertDetailDto) * models.DetailProduct {
    r := models.DetailProduct {}
    r = col3(dto.E, dto.H, "")
    r.Processid = dto.UniqueID
    r.MasterKod = dto.MasterCode
    r.MasteriId = dto.Masterid
    r.DetayId = dto.Dept
    return &r
}

func col3(d models.ExcelDto, h models.ExcelDto, _type string) models.DetailProduct {
    r := models.DetailProduct {}
    r.Tipi = d.Col1

    //col3 
    ////--------Col3-----------------		
    if h.Col1 == "Açıklama" {
        r.Aciklama = d.Col1
    }

    ////--------Col3-----------------		
    if h.Col2 == "Esk.Formülü" {
        r.EskFormülü = d.Col2
    } else if h.Col2 == "Açıklama" {
        r.Aciklama = d.Col2
    }

    ////--------Col3-----------------		
    if h.Col3 == "Malzeme" {
        r.Malzeme = d.Col3
    } else if h.Col3 == "Esk.Formülü" {
        r.EskFormülü = d.Col3
    } else if h.Col3 == "Fiyat" {
        r.Fiyat = d.Col3
    }

    ////--------Col4-----------------		
    if h.Col4 == "Açıklama" {
        r.Aciklama = d.Col4
    } else if h.Col4 == "Mlyt(Sa)" {
        r.Maliyetsa = d.Col4
    } else if h.Col4 == "Malzeme" {
        r.Malzeme = d.Col4
    }

    ////--------Col5-----------------		
    if h.Col5 == "Sarf" {
        r.Sarf = d.Col5
    } else if h.Col5 == "Tedarikçi" {
        r.Tedarikci = d.Col5
    } else if h.Col5 == "Op.Süresi(dk)" {
        r.OpSuresi = d.Col5
    } else if h.Col5 == "Açıklama" {
        r.Aciklama = d.Col5
    }

    ////--------Col6-----------------		
    if h.Col6 == "Brm" {
        r.Brm = d.Col6
    } else if h.Col6 == "Fire" {
        r.Fire = d.Col6
    } else if h.Col6 == "Fiyat" {
        r.Fiyat = d.Col6
    } else if h.Col6 == "Çvrm. Adt" {
        r.ÇevrimAdet = d.Col6
    } else if h.Col6 == "Tedarikçi" {
        r.Tedarikci = d.Col6
    }

    ////--------Col7-----------------		
    if h.Col7 == "MMO%" {
        r.MMO = d.Col7
    } else if h.Col7 == "Brm" {
        r.Brm = d.Col7
    } else if h.Col7 == "Pros.Net" {
        r.ProsNet = d.Col7
    } else if h.Col7 == "Fiyat" {
        r.Fiyat = d.Col7
    }

    ////--------Col8-----------------		
    if h.Col8 == "PMO%" {
        r.PMO = d.Col8
    } else if h.Col8 == "MMO%" {
        r.MMO = d.Col8
    } else if h.Col8 == "Pros.Net" {
        r.Prosfile = d.Col8
    } else if h.Col8 == "P.Brm" {
        r.Brm = d.Col8
    }

    ////--------Col9-----------------		
    if h.Col9 == "PFO%" {
        r.PFO = d.Col9
    } else if h.Col9 == "PMO%" {
        r.PMO = d.Col9
    } else if h.Col9 == "Pros." {
        r.Pros = d.Col9
    }

    ////--------Col10-----------------		
    if h.Col10 == "Esk.Formülü" {
        r.EskFormülü = d.Col10
    } else if h.Col10 == "Malz." {
        r.Malz = d.Col10
    } else if h.Col10 == "FO T." {
        r.FOT = d.Col10
    } else if h.Col10 == "PFO%" {
        r.PFO = d.Col10
    }

    ////--------Col11-----------------		
    if h.Col11 == "FO T." {
        r.FOT = d.Col11
    } else if h.Col11 == "Malz." {
        r.Malz = d.Col11
    } else if h.Col11 == "Pros." {
        r.Pros = d.Col11
    }

    ////--------Col12-----------------		
    if h.Col12 == "FO T." {
        r.FOT = d.Col12
    } else if h.Col12 == "Fiyat" {
        r.Fiyat = d.Col12
    } else if h.Col12 == "Pros." {
        r.Pros = d.Col12
    }

    ////--------Col13-----------------		
    if h.Col13 == "MO T." {
        r.MOT = d.Col13
    } else if h.Col13 == "Pros." {
        r.Pros = d.Col13
    }

    ////--------Col14-----------------		
    if h.Col14 == "MO T." {
        r.MOT = d.Col14
    } else if h.Col14 == "Fiyat" {
        r.Fiyat = d.Col14
    }

    ////--------Col15-----------------		
    if h.Col15 == "Fiyat" {
        r.Fiyat = d.Col15
    }

    return r
}


// func mlzControl(r models.DetailProduct, e models.ExcelDto, dept string) models.DetailProduct {
// 	if dept == "Col10" {
// 			r.Tipi = e.Col1
// 			r.Malzeme = e.Col3
// 			r.Aciklama = e.Col4
// 			r.Sarf = e.Col5
// 			r.Brm = e.Col6
// 			r.MMO = e.Col7
// 			r.PMO = e.Col8
// 			r.PFO = e.Col9
// 			r.Malz = e.Col10
// 			r.FOT = e.Col11
// 			r.Pros = e.Col12
// 			r.MOT = e.Col13
// 			r.Fiyat = e.Col14
// 	}
// 	return r
// }

// func ToplControl(r models.DetailProduct, e models.ExcelDto, dept string) models.DetailProduct {
// 	if e.Col4 == "Topl." {
// 			r.Toplam = e.Col4
// 			if e.Col14 == ""{
// 				r.Fiyat = e.Col15
// 			}else{
// 				r.Fiyat = e.Col14
// 			}
// 	}
// 	return r
// }


// func malzemeControl(r models.DetailProduct, e models.ExcelDto, dept string) models.DetailProduct {
//     if  e.Col10 == "Esk.Formülü" {
//         r.Tipi = e.Col1
//         r.Malzeme = e.Col3
//         r.Aciklama = e.Col4
//         r.Sarf = e.Col5
//         r.Brm = e.Col6
//         r.MMO = e.Col7
//         r.PMO = e.Col8
//         r.PFO = e.Col9
//         r.EskFormülü = e.Col10
//         r.Malz = e.Col11
//         r.FOT = e.Col12
//         r.Pros = e.Col13
//         r.MOT = e.Col14
//         r.Fiyat = e.Col15
//     }else if e.Col0 == "Malz." {

// 		}else if e.Col0 == "FOT." {

// 		}
//     return r
// }

// func AciklamaControl(r models.DetailProduct, e models.ExcelDto, dept string) models.DetailProduct {
// 	if dept == "Col2" {
// 			r.Tipi = e.Col1
// 			r.Aciklama = e.Col2
// 			r.EskFormülü = e.Col3
// 		  r.Maliyetsa = e.Col4
// 			r.OpSuresi = e.Col5
// 			r.ÇevrimAdet = e.Col6
// 			r.ProsNet = e.Col7
// 			r.Prosfile = e.Col8
// 			r.Pros = e.Col9
// 	}
// 	return r
// }

// func ConverToDetailDto1(orders models.ExcelDto, uniqueID string, masterid uint64, detayid string, header models.ExcelDto, masterCode string) *models.DetailProduct {
// 	insertType := utils.GetInsertType(orders)
// 	utils.WriteValue(insertType)
// 	utils.WriteValue(header.Col5)
// 	result := models.DetailProduct{}
// 	result.MasteriId = masterid
// 	result.Processid = uniqueID
// 	result.MasterKod = masterCode
// 	result.DetayId = detayid
// 	result.Tipi = orders.Col1
// 	if insertType == "none" {
// 		result.Malzeme = orders.Col3
// 		result.Aciklama = orders.Col4
// 		result.Sarf = orders.Col5
// 		//-------col6-----------------------
// 		if header.Col6 == "Fire" {
// 			result.Fire = orders.Col6
// 		}
// 		if header.Col6 == "Brm" {
// 			result.Brm = orders.Col6
// 		}
// 		//-------col7-----------------------
// 		if header.Col7 == "Brm" {
// 			result.Brm = orders.Col7
// 		}
// 		if header.Col7 == "MMO%" {
// 			result.MMO = orders.Col7
// 		}
// 		//-------col8-----------------------
// 		if header.Col8 == "MMO%" {
// 			result.MMO = orders.Col8
// 		}
// 		if header.Col8 == "PMO%" {
// 			result.PMO = orders.Col8
// 		}
// 		//-------col9-----------------------
// 		if header.Col9 == "PMO%" {
// 			result.PMO = orders.Col9
// 		}
// 		if header.Col9 == "PFO%" {
// 			result.PFO = orders.Col9
// 		}
// 		//-------col10-----------------------
// 		if header.Col10 == "PFO%" {
// 			result.PFO = orders.Col10
// 		}
// 		if header.Col10 == "Esk.Formülü" {
// 			result.EskFormülü = orders.Col10
// 		}
// 	} else if insertType == "Proses" {
// 		result.Aciklama = orders.Col2
// 		result.EskFormülü = orders.Col3
// 		result.Maliyetsa = orders.Col4
// 		result.OpSuresi = orders.Col5
// 		result.ÇevrimAdet = orders.Col6
// 		result.ProsNet = orders.Col7
// 		result.Prosfile = orders.Col8
// 		result.Pros = orders.Col9
// 	} else {
// 		result.Aciklama = orders.Col4
// 	}
// 	if insertType == "op" {
// 		result.PMO = orders.Col9
// 		result.PFO = orders.Col10
// 		result.FOT = orders.Col12
// 		result.Pros = orders.Col13
// 		result.MOT = orders.Col14
// 		result.Fiyat = orders.Col15
// 	}
// 	if header.Col5 == "Tedarikçi" {
// 		result.Tipi = orders.Col2
// 		result.Malzeme = orders.Col3
// 		result.Aciklama = orders.Col4
// 		result.Tedarikci = orders.Col5
// 		result.Fiyat = orders.Col6
// 	}
// 	if insertType == "none" {
// 		result.Malz = orders.Col12
// 		result.FOT = orders.Col12
// 		result.Pros = orders.Col13
// 		result.MOT = orders.Col14
// 		result.Fiyat = orders.Col15
// 	}
// 	if insertType == "Topl" {
// 		result.Aciklama = orders.Col1
// 		result.Fiyat = orders.Col3
// 	}
// 	if insertType == "adetay" {
// 		result.Aciklama = orders.Col4
// 		result.EskFormülü = orders.Col11
// 	}
// 	if insertType == "Topl_4" {
// 		result.Aciklama = orders.Col4
// 		result.EskFormülü = orders.Col6
// 	}
// 	if insertType == "Ambalaj" {
// 		result.EskFormülü = strings.ReplaceAll(orders.Col2, "*[TRY]", "")
// 		result.Fiyat = orders.Col3
// 	}
//     if(header.Col2 == "Esk.Formülü"){
//         result.Aciklama = orders.Col1
//         result.EskFormülü = orders.Col2
//         result.Fiyat = orders.Col3

//     }
//     if(header.Col3 == "Esk.Formülü"){
// 			result.Aciklama = orders.Col2
// 			result.EskFormülü = orders.Col3
// 			result.Maliyetsa = orders.Col4
// 			result.OpSuresi = orders.Col5
// 			result.ÇevrimAdet = orders.Col6
// 			result.ProsNet = orders.Col7
// 			result.Prosfile = orders.Col8
// 			result.Pros = orders.Col9

//     }
// 	return &result
// }

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