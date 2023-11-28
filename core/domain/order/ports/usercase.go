package ports

import "github.com/fillipehmeireles/order-service/pkg/handlers/order/dto"

type OrderUseCase interface {
	Create(newOrder dto.CreateOrderRequestDto) error
	GetAll() (dto.GetAllOrdersResponseDto, error)
	GetOne(id int) (dto.GetOneResponseDto, error)
	GetByUser(userID int) (dto.GetByUserResponseDto, error)
	Delete(id int) error
}
