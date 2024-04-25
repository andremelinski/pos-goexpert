package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andremelinski/pos-goexpert/sse/server/internal/infra/interfaces"
)

type DynamoInfo struct{
	ID int 
	Title string `json:"title"`
}

type NotificationHandler struct{
	NotificationDB interfaces.NotificationInterface
}
func NewNotificationHandler(db interfaces.NotificationInterface)*NotificationHandler{
	return &NotificationHandler{
		db,
	}
}

func(productHandler *NotificationHandler) GetNotification( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Cache-Control", "no-cache")
	notifications, err := productHandler.NotificationDB.GetAll()
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}
	w.WriteHeader(http.StatusOK)
	normalizedParamArr := []DynamoInfo{}
	for i, v := range notifications {
		normalizedParam := DynamoInfo{}
		 // convert map to json
		jsonString, _ := json.Marshal(v)
		fmt.Println(string(jsonString))
		// convert json to struct
		json.Unmarshal(jsonString, &normalizedParam)
		normalizedParam.ID = i

		normalizedParamArr = append(normalizedParamArr, normalizedParam)

		// fmt.Fprintf(w, "data: {index %d, title: %s} \n\n", m, normalizedParam.Title)
	}
	
	fmt.Fprintf(w, "event: message\n")
	fmt.Fprintf(w, "data: %+v\n\n", normalizedParamArr)
	w.(http.Flusher).Flush()
}