package dto

type CreateCategoryInput struct{
	Name string `json:"name"`
	Description string `json:"description"`
}

type CreateCourseInput struct {
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
}