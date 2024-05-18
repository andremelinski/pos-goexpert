package main

import (
	"database/sql"
	"net"

	"github.com/andremelinski/pos-goexpert/13-grpc/internal/database"
	"github.com/andremelinski/pos-goexpert/13-grpc/internal/pb"
	"github.com/andremelinski/pos-goexpert/13-grpc/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main(){
	dbCon, err := sql.Open("sqlite3", "./data.db")

	if err != nil{
		panic(err)
	}

	categoryDB := database.NewCategory(dbCon)
	categoryService := service.NewCategoryService(*categoryDB)
	
	grpcserver := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcserver, categoryService)

	reflection.Register(grpcserver)

	// apos conectar o server com o grpc, passar a porta de para ouvir
	lis, err := net.Listen("tcp", ":50051")
	
	if err != nil{
		panic(err)
	}

	//  tenta conectar e se der erro, ja da panic
	if err := grpcserver.Serve(lis); err != nil{
		panic(err)
	}
}