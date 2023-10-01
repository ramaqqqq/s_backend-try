package models

type User struct {
	UserID   int     `gorm:"primary_key;not null" json:"user_id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Role     string  `json:"role"`
	Balance  float64 `json:"balance"`

	Shoppingcarts []Shoppingcart
	Payments      []Payment
	BaseTime
}
