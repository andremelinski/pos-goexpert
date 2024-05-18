package service

import (
	"context"

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