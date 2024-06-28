package usecase

type HelloOuputDTO struct{
	Message string `json:"message"`
}

type HelloUseCase struct{}


func NewHelloUseCase(
) *HelloUseCase {
	return &HelloUseCase{}
}

func(h *HelloUseCase) Hello() *HelloOuputDTO{
	return &HelloOuputDTO{
		"hello",
	}
}