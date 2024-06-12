package product


type ProductUseCase struct{
	repository *ProductRepository
}

func NewProductUseCase(repository *ProductRepository) *ProductUseCase{
	return &ProductUseCase{
		repository,
	}
}

// GetProduct retorna a propia entidade, o que eh errado. Ao inves disso, deveria utilizar DTO.
// usecase nao deve ter acesso a entidade, no caso Product 
func (pu *ProductUseCase) GetProduct(id int) (Product, error){
	return pu.repository.GetProductId(id)
}