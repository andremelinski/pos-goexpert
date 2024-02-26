package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/andremelinski/pos-goexpert/desafios/Client-Server-API/server/api"
	"github.com/andremelinski/pos-goexpert/desafios/Client-Server-API/server/db"
	_ "github.com/mattn/go-sqlite3"
)

// https://mholt.github.io/json-to-go/

// var dbConfig *gorm.DB
var dbConfig *sql.DB

func main() {

	// config, err := gorm.Open(sqlite.Open("sqlite3"), &gorm.Config{})

	config, err := sql.Open("sqlite3", "./db/bid.db")
	config.Ping()
	
	if err != nil {
		panic(err)
	}

	// config.AutoMigrate(&db.UsdBrlGormModel{})
	path := filepath.Join("./db/sql","usdBrl.sql")
	c, _ := os.ReadFile(path)
	s := string(c)
	_ , err = config.Exec(s)
	if err != nil{
		panic(err)
	}
	dbConfig = config
	
	http.HandleFunc("/", getPrice)
	server := http.Server{
		Addr: ":8080",
		WriteTimeout: 2 * time.Second,
	}
	server.ListenAndServe()
}

func getPrice( w http.ResponseWriter, r *http.Request) {
	payload, err := api.USDBRLInit().GetDataFromApi()
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w ).Encode(err.Error())
		return
	}

	// newItem := db.RepoInit(dbConfig).CreateInfoDb(payload)
	newItem, err := db.RepoInitSql(dbConfig).CreateInfoDbSql(payload)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w ).Encode(err.Error())
		return
	}

	payloadNormalized := api.UsdbrlDto{}
	jsonPayload, _ := json.Marshal(newItem)
	json.Unmarshal([]byte(jsonPayload), &payloadNormalized)
	payloadNormalized.CreateDate = newItem.CreateDate

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w ).Encode(payloadNormalized)
}
