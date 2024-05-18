package main

import (
	"context"
	"flag"
	"log"

	"github.com/andremelinski/pos-goexpert/13-grpc/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", "world", "Name to greet")
)

var opts []grpc.DialOption

func main(){
	
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
	  panic(err)
	}

	client := pb.NewCategoryServiceClient(conn)

	feature, err := client.CreateCategory(context.Background(), &pb.CreateCategoryRequest{Name:"random name", Description: "dsescrpition"})
	if err != nil {
		panic(err)
	}

	log.Println(feature)
	defer conn.Close()
}
