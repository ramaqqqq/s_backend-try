package models

import "time"

type Item struct {
	ItemID   int     `gorm:"primary_key" json:"item_id"`
	UserID   int     `json:"user_id"`
	ItemName string  `json:"item_name"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
}

type Shoppingcart struct {
	PurchaseID   int       `gorm:"primary_key" json:"purchase_id"`
	UserID       int       `json:"user_id"`
	ItemID       int       `json:"item_id"`
	Quantity     int       `json:"quantity"`
	TotalPrice   float64   `json:"total_price"`
	PurchaseDate time.Time `json:"purchase_date"`
	Payment      Payment   `gorm:"foreignkey:PurchaseID"`
}

type Payment struct {
	PaymentID   int       `gorm:"primary_key" json:"payment_id"`
	UserID      int       `json:"user_id"`
	Amount      float64   `json:"amount"`
	Status      string    `json:"status"`
	OrderID     string    `json:"order_id"`
	SnapUrl     string    `json:"snap_url"`
	PaymentDate time.Time `json:"payment_date"`
	PurchaseID  int       `json:"purchase_id"`
}
