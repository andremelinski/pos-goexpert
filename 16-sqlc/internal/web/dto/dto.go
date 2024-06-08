package dto

type CreateCategoryInput struct{
	Name string `json:"name"`
	Description string `json:"description"`
}

type CreateCourseInput struct {
	Name        string    `json:"course_name"`
	Description string   `json:"course_description"`
	Price float64   `json:"price,omitempty"`
}