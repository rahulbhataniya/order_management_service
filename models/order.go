package models

// Order represents an order in the system
type Order struct {
	OrderID     int64   `gorm:"column:order_id;primaryKey;autoIncrement" json:"order_id"`
	UserID      int64   `gorm:"column:user_id;not null" json:"user_id" binding:"required"`
	ItemIDs     string  `gorm:"column:item_ids;not null" json:"item_ids" binding:"required"`
	TotalAmount float64 `gorm:"column:total_amount;not null" json:"total_amount" binding:"required"`
	Status      string  `gorm:"column:status;default:'pending'" json:"status"`
	CreatedAt   int64   `gorm:"autoCreateTime"` // Auto timestamp
	UpdatedAt   int64   `gorm:"autoUpdateTime"` // Auto timestamp
}
