package usescases

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/fillipehmeireles/order-service/core/domain/order"
	"github.com/fillipehmeireles/order-service/core/domain/order/ports"
	"github.com/fillipehmeireles/order-service/pkg/handlers/order/dto"
	userPorts "github.com/fillipehmeireles/user-service/core/domain/user/ports"
)

type OrderUseCase struct {
	orderRepo ports.OrderRepository
	userRepo  userPorts.UserRepository
}

// Create implements ports.OrderUseCase.
func (oUC *OrderUseCase) Create(newOrder dto.CreateOrderRequestDto) error {
	_, err := oUC.userRepo.GetOne(newOrder.UserID)
	if err != nil {
		return err
	}
	var order order.Order

	newOrder.ToDomain(&order)
	t := time.Now().Format(time.RFC850)
	tokenPair := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d%f%d%d%s", order.UserID, order.Amount, order.Direction, order.OrderType, t)))

	order.Pair = tokenPair
	if err := oUC.orderRepo.Create(order); err != nil {
		return err
	}

	return nil
}

// Delete implements ports.OrderUseCase.
func (oUC *OrderUseCase) Delete(id int) error {
	if err := oUC.orderRepo.Delete(id); err != nil {
		return err
	}

	return nil
}

// GetAll implements ports.OrderUseCase.
func (oUC *OrderUseCase) GetAll() (dto.GetAllOrdersResponseDto, error) {
	orders, err := oUC.orderRepo.GetAll()
	if err != nil {
		return dto.GetAllOrdersResponseDto{}, err
	}

	var ordrs dto.GetAllOrdersResponseDto

	ordrs.FromDomain(orders)
	return ordrs, nil
}

// GetByUser implements ports.OrderUseCase.
func (oUC *OrderUseCase) GetByUser(userID int) (dto.GetByUserResponseDto, error) {
	o, err := oUC.orderRepo.GetByUser(userID)
	if err != nil {
		return dto.GetByUserResponseDto{}, err
	}

	var ordr dto.GetByUserResponseDto

	ordr.FromDomain(o)
	return ordr, nil
}

// GetOne implements ports.OrderUseCase.
func (oUC *OrderUseCase) GetOne(id int) (dto.GetOneResponseDto, error) {
	o, err := oUC.orderRepo.GetOne(id)
	if err != nil {
		return dto.GetOneResponseDto{}, err
	}

	var ordr dto.GetOneResponseDto

	ordr.FromDomain(o)
	return ordr, nil
}

func NewOrderUseCase(orderRepo ports.OrderRepository, userRepo userPorts.UserRepository) ports.OrderUseCase {
	return &OrderUseCase{
		orderRepo: orderRepo,
		userRepo:  userRepo,
	}
}