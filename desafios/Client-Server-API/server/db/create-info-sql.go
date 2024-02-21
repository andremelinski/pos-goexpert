package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andremelinski/pos-goexpert/desafios/Client-Server-API/server/api"

	"time"

	"github.com/google/uuid"
)

type NewRepoSql struct {
	db *sql.DB
}

type UsdBrlModel struct {
	ID         string
	Code       string
	Codein     string
	Name       string
	High       string
	Low        string
	VarBid     string
	PctChange  string
	Bid        string
	Ask        string
	Timestamp  string
	CreateDate string
}

// Read implements io.Reader.
func (*UsdBrlModel) Read(p []byte) (n int, err error) {
	panic("unimplemented")
}

func RepoInitSql(db *sql.DB) *NewRepoSql {
	return &NewRepoSql{
		db: db,
	}
}

func (sql *NewRepoSql) CreateInfoDbSql(payload *api.USDBRL) (*UsdBrlModel, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	dataToAdd := newData(payload.Usdbrl)
	statement, err := sql.db.PrepareContext(ctx, `INSERT INTO usdBrl(
		id,
		code,
		codein,
		name,
		high,
		low,
		varBid,
		pctChange,
		bid,
		ask,
		timestamp,
		createDate
	) values(?,?,?,?,?,?,?,?,?,?,?,?)`)

	if err != nil {
		return nil, err
	}

	_, err = statement.Exec(dataToAdd.ID, dataToAdd.Code, dataToAdd.Codein, dataToAdd.Name, dataToAdd.High, dataToAdd.Low, dataToAdd.VarBid, dataToAdd.PctChange, dataToAdd.Bid, dataToAdd.Ask, dataToAdd.Timestamp, dataToAdd.CreateDate)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	stmt, err := sql.db.Prepare("SELECT bid, timestamp, createDate from usdBrl where id=?")

	if err != nil {
		return nil, err
	}

	getInfo := UsdBrlModel{}
	err = stmt.QueryRow(dataToAdd.ID).Scan(&getInfo.Bid, &getInfo.Timestamp, &getInfo.CreateDate)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	fmt.Println(getInfo.CreateDate)
	return &getInfo, nil
}

func newData(info api.UsdbrlDto) *UsdBrlModel {
	return &UsdBrlModel{
		ID:         uuid.New().String(),
		Code:       info.Code,
		Codein:     info.Codein,
		Name:       info.Name,
		High:       info.High,
		Low:        info.Low,
		VarBid:     info.VarBid,
		PctChange:  info.PctChange,
		Bid:        info.Bid,
		Ask:        info.Ask,
		Timestamp:  info.Timestamp,
		CreateDate: info.CreateDate,
	}

}
