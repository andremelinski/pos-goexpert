package entity

// interface que comunica o struct do usecase (CreateOrderUseCase) e o struct do repository (OrdeRepository)
type OrderRepositoryInterface interface {
	Save(order *Order) error
	// GetTotal() (int, error)
}