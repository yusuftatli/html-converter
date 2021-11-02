package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/orgs/mdyazilim/html-converter-new/common"
	"github.com/orgs/mdyazilim/html-converter-new/models"
)

func SaveMaster(masterData * models.MasterProduct) (uint64, string, error) {
    var _masterId uint64
    var _masterCode string
    var err error
    _prod, _ := strconv.ParseBool(common.GetEnvironment().Production)
    if (!_prod) {
        _masterId, _masterCode, err = savePostgreMaster(masterData)
    } else {
        _masterId, _masterCode, err =  saveOracleMaster(masterData)
    }
    return _masterId, _masterCode, err
}

func savePostgreMaster(masterData * models.MasterProduct) (uint64, string, error) {
    db := common.ConnectDB()
    err := db.Create( & masterData).Error
    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }
    common.CloseDBConnection(db)
    return masterData.Id, masterData.MasterKod, err
}

func saveOracleMaster(masterData *models.MasterProduct) (uint64, string, error) {
    db, err := sql.Open("goracle", common.GetEnvironment().OracleDatabaseUrl)
    if err != nil {
        fmt.Println("... DB Setup Failed")
        fmt.Println(err)
    }
    defer db.Close()
    dbQuery := "EXECUTE MASTERPRODUCT_INSERT('" + masterData.Processid + "','" + masterData.MasterKod + "','" + masterData.Tipi + "','" + masterData.Uy + "','" + masterData.Aciklama + "','" + masterData.Tedarikci + "','" + masterData.Fiyat + "','" + masterData.Parabirimi + "');"

    rows, err := db.Query(dbQuery)
    if err != nil {
        fmt.Println(".....Error processing query")
        fmt.Println(err)
    }
    defer rows.Close()
    return masterData.Id, masterData.MasterKod, err
}