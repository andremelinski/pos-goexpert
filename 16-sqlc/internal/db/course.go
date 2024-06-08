package db

import (
	"context"
	"database/sql"
	"fmt"
)

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

type CourseDB struct {
	dbConn *sql.DB
	*Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: New(dbConn),
	}
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	tx := NewTransaction(c.dbConn)
	err := tx.callTx(ctx, func(q *Queries) error {
		var err error
		_, err = q.CreateCategory(ctx, CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return err
		}
		err = q.CreateCourse(ctx, CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			CategoryID:  argsCategory.ID,
			Price:       argsCourse.Price,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *CourseDB) ListCourses(ctx context.Context) ([]ListCoursesRow, error) {
	courses, err := c.Queries.ListCourses(ctx)
	if err != nil {
		return nil, err
	}
	for _, course := range courses {
		fmt.Printf("Category: %s, Course ID: %s, Course Name: %s, Course Description: %s, Course Price: %f",
			course.CategoryName, course.ID, course.Name, course.Description.String, course.Price)
	}

	return courses, nil
}