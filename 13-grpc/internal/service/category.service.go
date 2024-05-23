package service

import (
	"context"
	"fmt"

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
		categoriesResponse = append(categoriesResponse, &pb.Category{
		Id: category.ID,
		Name: category.Name,
		Description: category.Description,
	})
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