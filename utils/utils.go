package utils

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/orgs/mdyazilim/html-converter-new/models"
)


func IsCheckSpace(value string) string {
    if value == "&nbsp;" {
        return "-"
    } else {
        return value
    }
}


func GetInsertType(masterRow models.ExcelDto) string {
    if strings.Contains(masterRow.Col4, "Plan") {
        return "op"
    } else if strings.Contains(masterRow.Col1, "Proses") {
        return "Proses"
    } else if strings.Contains(masterRow.Col1, "Topl") {
        return "Topl"
    } else if strings.Contains(masterRow.Col1, "Ambalaj") {
        return "Ambalaj"
    } else if strings.Contains(masterRow.Col5, "Tedar") {
        return "Tedarikci"
    } else if strings.Contains(masterRow.Col4, "Ambalaj") {
        return "adetay"
    } else if strings.Contains(masterRow.Col4, "Topl") {
        return "Topl_4"
    } else if strings.Contains(masterRow.Col4, "Esk") {
        return "Esk"
    } else if strings.Contains(masterRow.Col2, "Esk") {
        return "Eskcol2"
    } else {
        return "none"
    }
}

func WriteValue(value string) {
	 _writeValue := os.Getenv("WRITE_VALUES")
    if _writeValue == "true" {
        fmt.Println("-------------------------------------")
        fmt.Println(value)
    }
}

func GetDefaultColumns() * models.ExcelDto {
    columns := models.ExcelDto {}
    columns.Col0 = "col0"
    columns.Col1 = "col1"
    columns.Col2 = "col2"
    columns.Col3 = "col3"
    columns.Col4 = "col4"
    columns.Col5 = "col5"
    columns.Col6 = "col6"
    columns.Col7 = "col7"
    columns.Col8 = "col8"
    columns.Col9 = "col9"
    columns.Col10 = "col10"
    columns.Col11 = "col11"
    columns.Col12 = "col12"
    columns.Col13 = "col13"
    columns.Col14 = "col14"
    columns.Col15 = "col15"
    columns.Col16 = "col16"
    columns.Col17 = "col17"
    columns.Col18 = "col18"
    return &columns
}

func ParseData(value string, dept string) * models.ExcelDto {
    e := models.ExcelDto {}
    if strings.Contains(value, "<tbody>") && strings.Contains(value, "<table>") {
        e.Col0 = ""
    } else {
         value = strings.ReplaceAll(value, "</td>", "</td>\n")
                value = strings.ReplaceAll(value,"</tr>","")
            //  value = strings.ReplaceAll(value,"class=\"BG0\"","")
            //  value = strings.ReplaceAll(value,"class=\"BG0\"","")
            //  vaslue = strings.ReplaceAll(value,"class=\"BG0 L MLZ\"","")
            //  value = strings.ReplaceAll(value,"\u00a0</td>","")
            //  value = strings.ReplaceAll(value,"<td >","<td>")
            //  value = strings.ReplaceAll(value,"<td></td>","<td>-</td>")
            //  value = strings.ReplaceAll(value,"class=\"BG0 R FP1\"","")
        re := regexp.MustCompile(`<td.*?>(.*)</td>`)
        e.Col18 = dept
				_writeHeaders := os.Getenv("WRITE_HEADERS")
        if _writeHeaders == "true" && strings.Contains(value, "th") {
            value = strings.ReplaceAll(value, "</th>", "</th>\n")
            re = regexp.MustCompile(`<th.*?>(.*)</th>`)
            e.Col18 = dept + "_header_"
        }
        submatchall := re.FindAllStringSubmatch(value, -1)
        for i, element := range submatchall {
            switch i {
                case 0:
                    e.Col0 = IsCheckSpace(element[1])
                case 1:
                    e.Col1 = IsCheckSpace(element[1])
                case 2:
                    e.Col2 = IsCheckSpace(element[1])
                case 3:
                    e.Col3 = IsCheckSpace(element[1])
                case 4:
                    e.Col4 = IsCheckSpace(element[1])
                case 5:
                    e.Col5 = IsCheckSpace(element[1])
                case 6:
                    e.Col6 = IsCheckSpace(element[1])
                case 7:
                    e.Col7 = IsCheckSpace(element[1])
                case 8:
                    e.Col8 = IsCheckSpace(element[1])
                case 9:
                    e.Col9 = IsCheckSpace(element[1])
                case 10:
                    e.Col10 = IsCheckSpace(element[1])
                case 11:
                    e.Col11 = IsCheckSpace(element[1])
                case 12:
                    e.Col12 = IsCheckSpace(element[1])
                case 13:
                    e.Col13 = IsCheckSpace(element[1])
                case 14:
                    e.Col14 = IsCheckSpace(element[1])
                case 15:
                    e.Col15 = IsCheckSpace(element[1])
                case 16:
                    e.Col16 = IsCheckSpace(element[1])
                case 17:
                    e.Col17 = IsCheckSpace(element[1])
            }
        }
    }
    return &e
}

func GetDataFromHtml(data string, pattern string) string {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(data))
	s := doc.Find(pattern)
	html, _ := s.Html()
	return html
}

func FieldControl(row * models.ExcelDto) bool {
	if row.Col0 != "" {
			return true
	} else {
			return false
	}
}