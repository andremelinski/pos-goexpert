package usecase

import (
	"github.com/andremelinski/pos-goexpert/19-Clean-Arch/internal/entity"
	"github.com/andremelinski/pos-goexpert/19-Clean-Arch/pkg/events"
)
type OrderInputDTO struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type OrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type CreateOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface // capaz de criar uma ordem no banco de dados
	OrderCreated    events.EventInterface // evento 
	EventDispatcher events.EventDispatcherInterface // o cara que dispara o evento
}

func NewCreateOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: OrderRepository,
		OrderCreated:    OrderCreated,
		EventDispatcher: EventDispatcher,
	}
}

func (c *CreateOrderUseCase) Execute(input OrderInputDTO) (OrderOutputDTO, error) {
	order := entity.Order{
		ID:    input.ID,
		Price: input.Price,
		Tax:   input.Tax,
	}
	// executa regra de negocio
	order.CalculateFinalPrice()
	// acesso ao repo que executa uma acao (Save). Usecase nao sabe onde o dado vai ser salvo, apenas que vai ser salvo
	if err := c.OrderRepository.Save(&order); err != nil {
		return OrderOutputDTO{}, err
	}
	// prepara o output
	dto := OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice, // order.Price + order.Tax
	}

	c.OrderCreated.SetPayload(dto)
	c.EventDispatcher.Dispatch(c.OrderCreated)

	return dto, nil
}