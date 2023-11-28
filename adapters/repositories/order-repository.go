package repositories

import (
	"log"

	"github.com/fillipehmeireles/order-service/core/domain/order"
	"github.com/fillipehmeireles/order-service/core/domain/order/ports"
	"gorm.io/gorm"
)

type OrderRepository struct {
	dbInstance *gorm.DB
}

// Create implements ports.OrderRepository.
func (oRepo *OrderRepository) Create(newOrder order.Order) error {
	if err := oRepo.dbInstance.Create(&newOrder).Error; err != nil {
		log.Printf("[OrderRepository:Create] Error on creating new order: %s", err)
		return err
	}

	return nil
}

// Delete implements ports.OrderRepository.
func (oRepo *OrderRepository) Delete(id int) error {
	if err := oRepo.dbInstance.Delete(&order.Order{}, id).Error; err != nil {
		log.Printf("[OrderRepository:Delete] Error on deleting order %d: %s", id, err)
		return err
	}

	return nil
}

// GetAll implements ports.OrderRepository.
func (oRepo *OrderRepository) GetAll() (order.Orders, error) {
	var orders order.Orders
	if err := oRepo.dbInstance.Find(&orders).Error; err != nil {
		log.Printf("[OrderRepository:GetAll] Error on retrieving all orders: %s", err)
		return order.Orders{}, err
	}

	return orders, nil
}

// GetByUser implements ports.OrderRepository.
func (oRepo *OrderRepository) GetByUser(userID int) (order.Orders, error) {
	var orders order.Orders

	if err := oRepo.dbInstance.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		log.Printf("[OrderRepository:GetByUser] Error on retrieving user %d orders: %s", userID, err)
		return order.Orders{}, err
	}

	return orders, nil
}

// GetOne implements ports.OrderRepository.
func (oRepo *OrderRepository) GetOne(id int) (order.Order, error) {
	var ordr order.Order

	if err := oRepo.dbInstance.First(&ordr, id).Error; err != nil {
		log.Printf("[OrderRepository:GetOne] Error on retrieving one order: %s", err)
		return order.Order{}, err
	}

	return ordr, nil
}

func NewOrderRepository(dbInstance *gorm.DB) ports.OrderRepository {
	return &OrderRepository{
		dbInstance: dbInstance,
	}
}
