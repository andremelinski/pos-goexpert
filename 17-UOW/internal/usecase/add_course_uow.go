package usecase

import (
	"context"

	"github.com/andremelinski/pos-goexpert/17-UOW/internal/entity"
	"github.com/andremelinski/pos-goexpert/17-UOW/internal/repository"
	"github.com/andremelinski/pos-goexpert/17-UOW/pkg"
)

type InputUseCaseUOW struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUseCaseUOW struct {
	Uow pkg.UowInterface
}


func NewAddCourseUseCaseUOW(uow pkg.UowInterface) *AddCourseUseCaseUOW {
	return &AddCourseUseCaseUOW{
		uow,
	}
}

func (a *AddCourseUseCaseUOW) ExecuteUOW(ctx context.Context, input InputUseCase) error {
	a.Uow.Do(ctx, func(uow *pkg.UOW) error {
		catRepo := a.getCategoryRepository(ctx)
		courseRepo := a.getCourseRepository(ctx)

		category := entity.Category{
			Name: input.CategoryName,
		}
		
		err := catRepo.Insert(ctx, category)
		if err != nil {
			return err
		}

		course := entity.Course{
			Name:       input.CourseName,
			CategoryID: input.CourseCategoryID,
		}
	
		err = courseRepo.Insert(ctx, course)
		if err != nil {
			return err
		}

		return nil
	})

	return nil
}


// para poder implementar o GetRepository, eh necessario que ele volte com a interface repository.<RepoName>RepositoryInterface
func (a *AddCourseUseCaseUOW) getCategoryRepository(ctx context.Context) repository.CategoryRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.CategoryRepositoryInterface)
}

func (a *AddCourseUseCaseUOW) getCourseRepository(ctx context.Context) repository.CourseRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CourseRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.CourseRepositoryInterface)
}