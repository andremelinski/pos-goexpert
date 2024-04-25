package routes

import (
	"net/http"

	"github.com/andremelinski/pos-goexpert/sse/server/internal/infra/db"
	"github.com/andremelinski/pos-goexpert/sse/server/internal/infra/handlers"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type NotificationRoutes struct{
	mux *http.ServeMux
	dbConfig *dynamodb.Client
}

func NewNotificationRoutes(mux *http.ServeMux, dbConfig *dynamodb.Client)*NotificationRoutes{
	return &NotificationRoutes{
		mux,
		dbConfig,
	}
}

func(routes *NotificationRoutes) NotificationRoutesHandler(){
	userDB := db.NewNotificationDB(routes.dbConfig)
	userHandler := handlers.NewNotificationHandler(userDB)
	routes.mux.HandleFunc("/sse", userHandler.GetNotification)
}