package ports

import "github.com/fillipehmeireles/order-service/core/domain/order"

type OrderRepository interface {
	Create(newOrder order.Order) error
	GetAll() (order.Orders, error)
	GetByUser(userID int) (order.Orders, error)
	GetOne(id int) (order.Order, error)
	Delete(id int) error
}
