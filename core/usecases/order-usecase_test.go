package usecases_test

import (
	"testing"

	"github.com/fillipehmeireles/order-service/core/domain/order"
	"github.com/fillipehmeireles/order-service/core/domain/order/ports"
	"github.com/fillipehmeireles/order-service/core/domain/order/ports/mocks"
	"github.com/fillipehmeireles/order-service/core/usecases"
	"github.com/fillipehmeireles/order-service/pkg/handlers/order/dto"
	"github.com/fillipehmeireles/order-service/pkg/utils"
	"github.com/fillipehmeireles/user-service/core/domain/user"
	userMocks "github.com/fillipehmeireles/user-service/core/domain/user/ports/mocks"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type OrderUseCaseTestSuite struct {
	suite.Suite
	orderMock                   *mocks.OrderRepository
	userMock                    *userMocks.UserRepository
	orderService                ports.OrderUseCase
	mockOrderCreateRequestDto   dto.CreateOrderRequestDto
	mockOrderDomain             order.Order
	mockOrdersDomain            order.Orders
	mockOrdersByUserResponseDto dto.GetByUserResponseDto
	mockUserDomain              user.User
	orderID                     int
	userID                      int
}

func (suite *OrderUseCaseTestSuite) TestCreate_ShouldThrowAnErrorWhenEntityIsNotCreated() {
	suite.userMock.On("GetOne", suite.orderID).Return(suite.mockUserDomain, nil).Once()
	tokenPairTest := utils.GenerateTokenPairWithOrderData(suite.mockOrderCreateRequestDto)
	suite.mockOrderDomain.Pair = tokenPairTest
	suite.orderMock.On("Create", suite.mockOrderDomain).Return(nil).Once()
	err := suite.orderService.Create(suite.mockOrderCreateRequestDto)
	require.Nil(suite.T(), err)
}
func (suite *OrderUseCaseTestSuite) TestDelete_ShouldThrowAnErrorWhenEntityIsNotDeleted() {
	suite.orderMock.On("Delete", suite.orderID).Return(nil).Once()

	err := suite.orderService.Delete(suite.orderID)
	require.Nil(suite.T(), err)
}

func (suite *OrderUseCaseTestSuite) TestGetAll_ShouldThrowAnErrorWhenRepoCannotFetchOrders() {
	suite.orderMock.On("GetAll").Return(suite.mockOrdersDomain, nil).Once()

	orders, err := suite.orderService.GetAll()

	require.Nil(suite.T(), err)
	require.Equal(suite.T(), len(orders.Orders), len(suite.mockOrdersDomain))
}

func (suite *OrderUseCaseTestSuite) TestGetOne_ShouldThrowAnErrorWhenRepoCannotFetchOneOrderByID() {
	suite.orderMock.On("GetOne", suite.orderID).Return(suite.mockOrderDomain, nil).Once()

	order, err := suite.orderService.GetOne(suite.orderID)

	require.Nil(suite.T(), err)
	order_test := order
	order.FromDomain(suite.mockOrderDomain)
	require.Equal(suite.T(), order_test, order)

}

func (suite *OrderUseCaseTestSuite) TestGetOne_ShouldThrowAnErrorWhenRepoCannotFetchOneOrderByUserID() {
	suite.orderMock.On("GetByUser", suite.userID).Return(suite.mockOrdersDomain, nil).Once()

	orders, err := suite.orderService.GetByUser(suite.userID)

	require.Nil(suite.T(), err)
	orders_test := orders
	require.Equal(suite.T(), orders_test, orders)

}

func (suite *OrderUseCaseTestSuite) SetupTest() {
	suite.orderMock = &mocks.OrderRepository{}
	suite.userMock = &userMocks.UserRepository{}
	newDto := dto.CreateOrderRequestDto{}
	newDto.UserID = 1
	newDto.Direction = 1
	newDto.OrderType = 1
	newDto.Amount = 3.2
	suite.orderID = 1
	suite.userID = 1
	suite.mockOrderCreateRequestDto = newDto
	suite.mockOrderDomain = *order.NewOrder(0, suite.userID, "", 3.2, order.DirectionBuy, order.OrderTypeMarket)
	suite.mockOrdersDomain = order.Orders{
		{
			ID:        1,
			UserID:    1,
			Pair:      "",
			Amount:    9.1,
			Direction: order.DirectionSell,
			OrderType: order.OrderTypeLimit,
		},
		{
			ID:        2,
			UserID:    1,
			Pair:      "",
			Amount:    84.9,
			Direction: order.DirectionBuy,
			OrderType: order.OrderTypeMarket,
		},
	}
	suite.mockUserDomain = user.User{
		Name:        "Fillipe",
		Email:       "fillipe.dev@gmail.com",
		PhoneNumber: "00000000000",
	}

	ordersByUserResponseDto := dto.GetByUserResponseDto{}

	ordersByUserResponseDto.Orders = []dto.OrderResponseDto{
		{
			ID:        1,
			UserID:    1,
			Pair:      "",
			Amount:    9.1,
			Direction: 1,
			OrderType: 1,
		},
		{
			ID:        2,
			UserID:    1,
			Pair:      "",
			Amount:    84.9,
			Direction: 2,
			OrderType: 2,
		},
	}
	suite.mockOrdersByUserResponseDto = ordersByUserResponseDto
	suite.orderService = usecases.NewOrderUseCase(suite.orderMock, suite.userMock)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderUseCaseTestSuite))
}
