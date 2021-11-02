package models

import "time"


type MasterProduct struct {
	Id         uint64    `gorm:"primaryKey"`
	Processid  string    `gorm:"column:processid"`
	MasterKod  string    `gorm:"column:malzemecode"`
	Tipi       string    `gorm:"column:tipi"`
	Uy         string    `gorm:"column:uy"`
	Aciklama   string    `gorm:"column:aciklama"`
	Tedarikci  string    `gorm:"column:tedarikci"`
	Fiyat      string    `gorm:"column:fiyat"`
	Parabirimi string    `gorm:"column:parabirimi"`
  PackageCretaDate string    `gorm:"column:packagecretadate"`
  PackageNumber string `gorm:"column:packagenumber"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time ` gorm:"column:updated_at"`
}

type DetailProduct struct {
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