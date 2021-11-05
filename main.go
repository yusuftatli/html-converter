package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	handlers "github.com/orgs/mdyazilim/html-converter-new/Handlers"
	_ "gopkg.in/goracle.v2"
)


func main() {
    // common.GetEnvironment()
    // WRITE_VALUES=false
    os.Setenv("WRITE_VALUES", "false")
        // WRITE_HEADERS=true
    os.Setenv("WRITE_HEADERS", "true")
        // PRODUCTION=false
    os.Setenv("PRODUCTION", "false")
        // POSTGRE_DATABASE_URL=host=localhost user=postgres password=12345 dbname=htmlconvert port=5435 sslmode=disable
    os.Setenv("POSTGRE_DATABASE_URL", "host=localhost user=postgres password=12345 dbname=htmlconvert port=5435 sslmode=disable")
        // ORACLE_DATABASE_URL=
    os.Setenv("ORACLE_DATABASE_URL", "")
        // MAIN_ROOT="/Users/tcytatli/Documents/md/prod/"
    os.Setenv("MAIN_ROOT", "./")
        // MOVE_PATH=true
    os.Setenv("MOVE_PATH", "true")

    log.Println("job started")
    ticker :=  time.NewTicker(time.Second*100)
    for t :=  range ticker.C {
        log.Println("Job tetiklendi ", t)
            //  fmt.Print("\033[H\033[2J")
        readFiles()
    }
}

func readFiles() {
    files :=  handlers.GetFiles()
    if len(files) > 0 {
        for _, file :=  range files {
            var path string
            var path2 string
            path = os.Getenv("MAIN_ROOT") + "pages/"
            path2 = path + file.Name()
            htmlBytes, err :=  ioutil.ReadFile(path2)
            if err != nil {
                log.Panicln(err)
            }
            if strings.Contains(file.Name(), "ht") {
                log.Println(file.Name() + " => " + "Okunuyor...")
                newid :=  uuid.New()
                parseHtml(string(htmlBytes), file.Name(), newid.String())
                log.Println(file.Name() + " => " + "Okundu...")
            }
        }
    } else {
        log.Println("Okunacak dosya bulunamadı...")
    }
}

func parseHtml(htmlData string, fileName string, uniqueID string) {
    _packageCreataDate, _packageNumber :=  handlers.GetHeaderData(htmlData)
    parsedHtmlDataList :=  handlers.GetParsedHtml(htmlData, fileName, uniqueID)
    handlers.SaveToDatabase(parsedHtmlDataList, uniqueID, _packageCreataDate, _packageNumber, fileName)
    handlers.SaveToExcel(parsedHtmlDataList, fileName, _packageCreataDate, _packageNumber)
    handlers.MoveFile(fileName)
}