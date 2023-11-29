package dto

import "github.com/fillipehmeireles/order-service/core/domain/order"

type orderRequestDto struct {
	UserID    int     `json:"user_id"`
	Amount    float32 `json:"amount"`
	Direction int     `json:"direction"`
	OrderType int     `json:"order_type"`
}

type OrderResponseDto struct {
	ID        int     `json:"id"`
	UserID    int     `json:"user_id"`
	Pair      string  `json:"pair"`
	Amount    float32 `json:"amount"`
	Direction int     `json:"direction"`
	OrderType int     `json:"order_type"`
}

type requestDto struct {
	orderRequestDto
}
type responseManyOrdersDto struct {
	Orders []OrderResponseDto `json:"orders"`
}

type (
	CreateOrderRequestDto struct {
		requestDto
	}
	GetAllOrdersResponseDto struct {
		responseManyOrdersDto
	}
	GetOneResponseDto struct {
		OrderResponseDto
	}
	GetByUserResponseDto struct {
		responseManyOrdersDto
	}
)

func (rDto *requestDto) ToDomain(orderDomain *order.Order) {
	orderDomain.UserID = rDto.UserID
	orderDomain.Amount = rDto.Amount
	orderDomain.Direction = order.Direction(rDto.Direction)
	orderDomain.OrderType = order.OrderType(rDto.OrderType)
}

func (gODto *GetOneResponseDto) FromDomain(orderDomain order.Order) {
	gODto.ID = orderDomain.ID
	gODto.UserID = orderDomain.UserID
	gODto.Pair = orderDomain.Pair
	gODto.Amount = orderDomain.Amount
	gODto.Direction = int(orderDomain.Direction)
	gODto.OrderType = int(orderDomain.OrderType)
}

func (gDto *responseManyOrdersDto) FromDomain(ordersDomain order.Orders) {
	for _, o := range ordersDomain {
		gDto.Orders = append(gDto.Orders, OrderResponseDto{
			ID:        o.ID,
			UserID:    o.UserID,
			Pair:      o.Pair,
			Amount:    o.Amount,
			Direction: int(o.Direction),
			OrderType: int(o.OrderType),
		})
	}
}
