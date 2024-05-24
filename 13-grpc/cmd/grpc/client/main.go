package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

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
		
	// newUser, err := client.CreateCategory(context.Background(), &pb.CreateCategoryRequest{Name:"random name", Description: "dsescrpition"})
	// if err != nil {
	// 	panic(err)
	// }

	// log.Println(newUser)

	// in := &pb.Blank{}
	// categories, err := client.ListCategory(context.Background(), in)
	// if err != nil {
	// 	panic(err)
	// }

	// log.Println(categories.Categories)
	// id := &pb.CategoryIdRequest{Id:"2ee6e3c1-02fc-48ff-8b55-7ba878b98183"}
	// category, err := client.GetCategory(context.Background(), id)
	// if err != nil {
	// 	panic(err)
	// }

	// log.Println(category)
	// defer conn.Close()

	// stream, err := client.CreateCategoryStream(context.Background())
	// for i := 0; i < 10; i++ {
	// 	name := fmt.Sprintf("category from stream %d", i)
	// 	err := stream.Send(&pb.CreateCategoryRequest{Name: name, Description: "description"})
	// 	if err != nil {
	// 		log.Fatalln("Send stream",err)
	// 	}

	// 	time.Sleep(100*time.Millisecond)
	// }
	// res, err := stream.CloseAndRecv()
	// if err != nil {
	// 	log.Fatalln("CloseAndRecv stream",err)
	// }

	// fmt.Println(res.Categories)
	// defer conn.Close()

	stream, err := client.CreateCategoryStreamBidirectional(context.Background())
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("category from bidirectional stream %d", i)
		err := stream.Send(&pb.CreateCategoryRequest{Name: name, Description: "description"})
		if err != nil {
			log.Fatalln("Send stream",err)
		}
		time.Sleep(100*time.Millisecond)
		
		res, err := stream.Recv()
		if err != nil {
			log.Fatalln("CloseAndRecv stream",err)
		}
	
		fmt.Println(res)
	}
	defer conn.Close()
}
