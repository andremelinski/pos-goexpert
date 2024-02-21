package db

import (
	"fmt"
	"goexpert/desafios/Client-Server-API/server/api"

	"gorm.io/gorm"
)

type NewRepo struct {
	db *gorm.DB
}

type UsdBrlGormModel struct {
	ID int `gorm:"primaryKey"`
	Code string
	Codein string
	Name string
	High string
	Low string
	VarBid string
	PctChange string
	Bid string
	Ask string
	Timestamp string
	CreateDate string
	gorm.Model
}

func RepoInit(db *gorm.DB) *NewRepo {
	return &NewRepo{
		db: db,
	}
}

func (gorm *NewRepo) CreateInfoDb(payload *api.USDBRL) *UsdBrlGormModel{
	fmt.Println(payload)
	info := payload.Usdbrl
	gorm.db.Create(&UsdBrlGormModel{
		Code: info.Code,
		Codein: info.Codein,
		Name: info.Name,
		High: info.High,
		Low: info.Low,
		VarBid: info.VarBid,
		PctChange: info.PctChange,
		Bid: info.Bid,
		Ask: info.Ask,
		Timestamp: info.Timestamp,
		CreateDate: info.CreateDate,
	})

	usdBrlGormFound := UsdBrlGormModel{}

	gorm.db.First(&usdBrlGormFound, "timestamp=?", info.Timestamp)

	return &usdBrlGormFound
}