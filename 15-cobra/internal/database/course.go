package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
    db *sql.DB
    ID string
    Name string
	Description string
    CategoryID string
}

func NewCourse(db *sql.DB) *Course {
    return &Course{db: db}
}
func (c *Course) Create(name, description, categoryId string) (Course, error){
	categoryIdResponse, err  := c.findCategoryById(categoryId)

	if err != nil || *categoryIdResponse =="" {
		return Course{}, err
	}

	id := uuid.New().String()
	_, err = c.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)", id, name, description, *categoryIdResponse)
	if err != nil {
		return Course{}, err
	}
	return Course{ID: id, Name: name, Description: description,CategoryID: *categoryIdResponse}, nil
}

func (c *Course) LoadAll() ([]Course, error){
	coursesRows, err := c.db.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}

	courses := []Course{}
	for coursesRows.Next() {
		course := Course{}
		err = coursesRows.Scan(&course.ID, &course.Name,  &course.Description, &course.CategoryID)
		if err != nil{
			return nil, err
		}
		courses = append(courses, course)
	}

	defer coursesRows.Close()

	return courses, nil
}

func (r Course) FindCoursesByCategoryId(categoryId string) ([]Course, error){
	query, err := r.db.Prepare("SELECT id, name, description, category_id FROM courses WHERE category_id = $1")
	if err != nil{
			return nil, err
	}
	
	coursesRows,err := query.Query(categoryId)

	if err != nil{
			return nil, err
	}
	defer coursesRows.Close()
	
	courses := []Course{}
	for coursesRows.Next() {
		course := Course{}
		err = coursesRows.Scan(&course.ID, &course.Name,  &course.Description, &course.CategoryID)
		if err != nil{
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

func (r Course) findCategoryById(categoryId string) (*string, error){
	queryRow := r.db.QueryRow("SELECT id FROM categories WHERE id=?", categoryId)
	if queryRow.Err() != nil{
		return nil, queryRow.Err()
	}
	var id string
	queryRow.Scan(&id)

	return &id, nil
}