package interfaces

import (
	"context"
	"database/sql"

	"github.com/andremelinski/pos-goexpert/16-sqlc/internal/db"
)

type CategoryInterface interface {
	CreateCategory(ctx context.Context, arg db.CreateCategoryParams) (sql.Result, error)
	GetCategory(ctx context.Context, id string) (db.Category, error)
	ListCategories(ctx context.Context) ([]db.Category, error)
	DeleteCategory(ctx context.Context, id string) error
}