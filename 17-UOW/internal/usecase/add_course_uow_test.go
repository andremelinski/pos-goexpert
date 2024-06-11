package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/andremelinski/pos-goexpert/17-UOW/internal/db"
	"github.com/andremelinski/pos-goexpert/17-UOW/internal/repository"
	"github.com/andremelinski/pos-goexpert/17-UOW/pkg"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func TestAddCourseUOW(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	assert.NoError(t, err)

	dbt.Exec("DROP TABLE if exists `courses`;")
	dbt.Exec("DROP TABLE if exists `categories`;")

	dbt.Exec("CREATE TABLE IF NOT EXISTS `categories` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL);")
	dbt.Exec("CREATE TABLE IF NOT EXISTS `courses` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL, category_id INTEGER NOT NULL, FOREIGN KEY (category_id) REFERENCES categories(id));")

	ctx := context.Background()
	uow := pkg.NewUow(ctx, dbt)

	uow.Register("CategoryRepository", func(tx *sql.Tx) interface{} {
		repoCat := repository.NewCategoryRepository(dbt)
		// no struct, NewCategoryRepository inicia o db.New. Entretanto, como nao queremos iniciar uma conexao mas a mesma transacao que NewCourseRepository utilizara, 
		// devemos falar que Queries aplica db.New(tx) em ambos.
		repoCat.Queries = db.New(dbt)
		return repoCat
	})

	uow.Register("CourseRepository", func(tx *sql.Tx) interface{} {
		repoCourse := repository.NewCourseRepository(dbt)
		repoCourse.Queries = db.New(tx)
		return repoCourse
	})
	
	input := InputUseCase{
	CategoryName:     "Category 1", // ID->1
	CourseName:       "Course 1",
	// CourseCategoryID: 2, // quebra
	CourseCategoryID: 1, // da certo
	}
	
	useCase := NewAddCourseUseCaseUow(uow)
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
