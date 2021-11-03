package handlers

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

//fileName: dosyanın adı
func MoveFile(fileName string) error {
    var err error
    _movePath := os.Getenv("MOVE_PATH")
    _mainRoot :=  os.Getenv("MAIN_ROOT")
    if (_movePath == "true") {
        err = os.Rename(_mainRoot + "./pages/" + fileName, _mainRoot + "./readedPage/" + fileName)
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