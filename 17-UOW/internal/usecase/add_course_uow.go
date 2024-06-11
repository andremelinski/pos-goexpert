package usecase

import (
	"context"

	"github.com/andremelinski/pos-goexpert/17-UOW/internal/entity"
	"github.com/andremelinski/pos-goexpert/17-UOW/internal/repository"
	"github.com/andremelinski/pos-goexpert/17-UOW/pkg"
)

type InputUseCaseUow struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUseCaseUOW struct{
	uow pkg.UowInterface
}

func NewAddCourseUseCaseUow(uow pkg.UowInterface) *AddCourseUseCaseUOW{
	return &AddCourseUseCaseUOW{
	uow,
	}
}


func (a *AddCourseUseCaseUOW)Execute(ctx context.Context, input InputUseCase) error{
	err := a.uow.Do(ctx, func(uow *pkg.Uow) error{
		// para ocorrer uma transaction o que precisa:
		// pegar o repo -> mandar executar o que vc quer, tratar erro e fazer isso no outro repo.
		// transaction ocorre pelo context, por isso deve ser carregado

		category := entity.Category{
			Name: input.CategoryName,
		}

		// como usecase nao deve ter acesso a camada repo, deve utilizar o Uow.GetRepository
		// catRepo := repository.NewCategoryRepository(uow.Db)
		catRepo := a.getCategoryRepository(ctx)
		err := catRepo.Insert(ctx, category)
		if err != nil {
			return err
		}

		course := entity.Course{
			Name:       input.CourseName,
			CategoryID: input.CourseCategoryID,
		}

		// courseRepo := repository.NewCourseRepository(uow.Db)sss
		courseRepo := a.getCourseRepository(ctx)
		err = courseRepo.Insert(ctx, course)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil{
		return err
	}
	return nil
}

// para poder implementar o GetRepository, eh necessario que ele volte com a interface repository.<RepoName>RepositoryInterface
func (a *AddCourseUseCaseUOW) getCategoryRepository(ctx context.Context) repository.CategoryRepositoryInterface {
	repo, err := a.uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.CategoryRepositoryInterface)
}

func (a *AddCourseUseCaseUOW) getCourseRepository(ctx context.Context) repository.CourseRepositoryInterface {
	repo, err := a.uow.GetRepository(ctx, "CourseRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.CourseRepositoryInterface)
}