package handlers

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/orgs/mdyazilim/html-converter-new/common"
)

//fileName: dosyanın adı
func MoveFile(fileName string) error {
    var err error
    _movePath, err := strconv.ParseBool(common.GetEnvironment().MovePath)
    if (_movePath) {
        err = os.Rename(common.GetEnvironment().MainRoot + "./pages/" + fileName, common.GetEnvironment().MainRoot + "./readedPage/" + fileName)
    }
    return err
}

func GetFiles() []fs.FileInfo {
    path := os.Getenv("MAIN_ROOT")
    files, err := ioutil.ReadDir(os.Getenv("MAIN_ROOT") + "pages")
    if err != nil {
        log.Println("path=> ",path)
        log.Fatal("hata burada")
        
    }
    return files
}