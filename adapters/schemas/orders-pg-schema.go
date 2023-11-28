package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        int `gorm:"primaryKey"`
	UserID    int
	Pair      string `gorm:"unique"`
	Amount    float32
	Direction int
	OrderType int
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}
