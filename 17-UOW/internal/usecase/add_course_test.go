package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/andremelinski/pos-goexpert/17-UOW/internal/repository"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func TestAddCourse(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	assert.NoError(t, err)

	dbt.Exec("DROP TABLE if exists `courses`;")
	dbt.Exec("DROP TABLE if exists `categories`;")

	dbt.Exec("CREATE TABLE IF NOT EXISTS `categories` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL);")
	dbt.Exec("CREATE TABLE IF NOT EXISTS `courses` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL, category_id INTEGER NOT NULL, FOREIGN KEY (category_id) REFERENCES categories(id));")

	input := InputUseCase{
		CategoryName:     "Category 1", // ID->1
		CourseName:       "Course 1",
		CourseCategoryID: 2, // quebra
		// CourseCategoryID: 1, // da certo
	}

	ctx := context.Background()

	useCase := NewAddCourseUseCase(repository.NewCourseRepository(dbt), repository.NewCategoryRepository(dbt))
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}

/*
utilizando courseCategoryID: 2 da erro: Cannot add or update a child row: a foreign key constraint fails
pq? pq criamos a category id 1 e pra salvar o curso estamos usando 2. Com isso, nao eh criado curso mas eh criado uma categoria, geralmente inconsistencia. 
Correto: se der erro na transacao, dar rollback em tudo. Pra isso deve se usar transaction.
Problema: como para o use case esta usando repositorios e esses repositorios devem ser independnetes,
como fazer essa transancao se eles sao independentes? eh de resposa do usecase cuidar disso?
*/