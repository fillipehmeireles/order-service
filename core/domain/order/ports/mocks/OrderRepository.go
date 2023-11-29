// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	order "github.com/fillipehmeireles/order-service/core/domain/order"
	mock "github.com/stretchr/testify/mock"
)

// OrderRepository is an autogenerated mock type for the OrderRepository type
type OrderRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: newOrder
func (_m *OrderRepository) Create(newOrder order.Order) error {
	ret := _m.Called(newOrder)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(order.Order) error); ok {
		r0 = rf(newOrder)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *OrderRepository) Delete(id int) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *OrderRepository) GetAll() (order.Orders, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 order.Orders
	var r1 error
	if rf, ok := ret.Get(0).(func() (order.Orders, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() order.Orders); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(order.Orders)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUser provides a mock function with given fields: userID
func (_m *OrderRepository) GetByUser(userID int) (order.Orders, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for GetByUser")
	}

	var r0 order.Orders
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (order.Orders, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(int) order.Orders); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(order.Orders)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOne provides a mock function with given fields: id
func (_m *OrderRepository) GetOne(id int) (order.Order, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetOne")
	}

	var r0 order.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (order.Order, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) order.Order); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(order.Order)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOrderRepository creates a new instance of OrderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrderRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrderRepository {
	mock := &OrderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
