package dto

type CreateCategoryInput struct{
	Name string `json:"name"`
	Description string `json:"description"`
}

type CreateCourseInput struct {
	Name        string    `json:"name"`
	// ,omitempty
	Description string   `json:"description"`
	Price float64   `json:"price"`
}