package service

import (
	"context"
	"fmt"
	"io"

	"github.com/andremelinski/pos-goexpert/13-grpc/internal/database"
	"github.com/andremelinski/pos-goexpert/13-grpc/internal/pb"
)

type CategoryService struct{
	pb.UnimplementedCategoryServiceServer
	categoryDb database.Category
}

func NewCategoryService(categoryDb database.Category) *CategoryService {
	return &CategoryService{
		categoryDb: categoryDb,
	}
}

func (cs *CategoryService) CreateCategory(cx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := cs.categoryDb.Create(in.Name, in.Description)

	if err != nil{
		return nil, err
	}

	generateCategoryForRPC := &pb.Category{
		Id: category.ID,
		Name: category.Name,
		Description: category.Description,
	}
	// a CategoryResponse tem que ser to tipo Category, como foi feito  no .proto
	return &pb.CategoryResponse{
		Category: generateCategoryForRPC,
	}, nil
}

func (cs *CategoryService) ListCategory(ctx context.Context, in *pb.Blank) (*pb.CategoryListResponse, error){
	categories, err := cs.categoryDb.LoadAll()
	if err != nil{
		return nil, err
	}

	categoriesResponse := []*pb.Category{}

	for _, category := range categories {
		categoryResponse := &pb.Category{
		Id: category.ID,
		Name: category.Name,
		Description: category.Description,
	}
		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return &pb.CategoryListResponse{
		Categories: categoriesResponse,
	}, nil
}

func (cs *CategoryService) GetCategory(ctx context.Context, in *pb.CategoryIdRequest) (*pb.CategoryResponse, error){
	fmt.Println()
	category, err := cs.categoryDb.LoadCategoryById(in.Id)
	if err != nil{
		return nil, err
	}

	categoriesResponse := &pb.Category{
		Id: category.ID,
		Name: category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoriesResponse,
	}, nil
}

func (cs *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryListResponse{}
	//  cria um loop infinito pq ele eh responsavel por mandar esse stream de dados
	for{
		// recebe os dados pra criar a categoria
		category, err := stream.Recv()
		//  final do stream
		if err == io.EOF{
			return stream.SendAndClose(categories)
		}
		if err != nil {
			return err
		}

		newCategory, err := cs.categoryDb.Create(category.Name, category.Description)
		if err != nil {
			return err
		}
		categories.Categories = append(categories.Categories, &pb.Category{
			Id: newCategory.ID,
			Name: newCategory.Name,
			Description: newCategory.Description,
		})
	}
}