package models

import (
	"time"
)


type MasterProduct struct {
	Id         uint64    `gorm:"primaryKey"`
	Processid  string    `gorm:"column:processid"`
	MasterKod  string    `gorm:"column:malzemecode"`
	FileName   string		 `gorm:"column:filename"`
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
	DetayId    string `gorm:"column:detayid"`
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
	PaketTarihi  string `gorm:"column:pakettarihi"`
	PaketNumarasi  string `gorm:"column:paketnumarasi"`
}

type ExcelDto struct {
	Col0 string
	Col1 string	
	Col2 string
	Col3 string
	Col4 string
	Col5 string
	Col6 string
	Col7 string
	Col8 string
	Col9 string
	Col10 string
	Col11 string
	Col12 string
	Col13 string
	Col14 string
	Col15 string
	Col16 string
	Col17 string
	Col18 string
}

type SaveDatabaseDto struct{
	O []ExcelDto
	UniqueID string
	PackageCreataDate string
	PackageNumber string
	FileName string
}

type ConvertDetailDto struct{
	E ExcelDto
	UniqueID string
	Masterid uint64
	Dept string
	H ExcelDto
	MasterCode string
	PackageCreataDate string
	PackageNumber string
}

type ConvertMasterDto struct{
	E ExcelDto 
	UniqueID string
	PackageCreataDate string
	PackageNumber string 
	FileName string
}

const ( 
	Dept_1 = "1"
	Dept_2 = "2"
	Dept_3 = "3"
	Dept_4 = "4"
	Dept_5 = "5"
)