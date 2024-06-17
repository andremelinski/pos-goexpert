package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/andremelinski/pos-goexpert/19-Clean-Arch/cmd/queue"
	"github.com/andremelinski/pos-goexpert/19-Clean-Arch/configs"
	event "github.com/andremelinski/pos-goexpert/19-Clean-Arch/internal/events"
	"github.com/andremelinski/pos-goexpert/19-Clean-Arch/internal/events/handler"
	"github.com/andremelinski/pos-goexpert/19-Clean-Arch/internal/infra/graph"
	"github.com/andremelinski/pos-goexpert/19-Clean-Arch/internal/infra/grpc/pb"
	"github.com/andremelinski/pos-goexpert/19-Clean-Arch/internal/infra/grpc/service"
	"github.com/andremelinski/pos-goexpert/19-Clean-Arch/internal/infra/web/webserver"
	"github.com/andremelinski/pos-goexpert/19-Clean-Arch/pkg/events"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/go-sql-driver/mysql"
)


func main(){
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	// deve registrar um handle para colocar no rabbitMQ
	eventName := event.NewOrderCreated().Name
	eventDispatcher.Register(eventName, &handler.OrderCreatedHandler{
		RabbitMQChannel: queue.GetRabbitMQChannel(),
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler("/order", webOrderHandler.Create)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)
}

