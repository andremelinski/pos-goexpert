package main

import (
	"database/sql"
	"encoding/json"
	"goexpert/desafios/Client-Server-API/server/api"
	"goexpert/desafios/Client-Server-API/server/db"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// https://mholt.github.io/json-to-go/

// var dbConfig *gorm.DB
var dbConfig *sql.DB

func main() {

	// dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	// config, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	config, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	config.Ping()
	
	if err != nil {
		panic(err)
	}

	// config.AutoMigrate(&db.UsdBrlGormModel{})
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
	return
}
