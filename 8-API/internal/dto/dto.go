package dto

// utilizado para limpar o que vem da request antes de ir para as outras camadas
//  assim fica mais facil de fazer o bind entre camadas pq a info ta limpa

type CreateProductInput struct{
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type CreateUserInput struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type GetJWTInput struct{
	Email string `json:"email"`
	Password string `json:"password"`
}