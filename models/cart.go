package models

type Cart struct {
	ID            int                 `json:"id" gorm:"primary_key:auto_increment"`
	ProductID     int                 `json:"product_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product       ProductResponse     `json:"product"`
	TransactionID int                 `json:"transaction_id"`
	Transaction   TransactionResponse `json:"transaction"`
	ToppingID     []int               `json:"topping_id" form:"topping_id" gorm:"-"`
	Topping       []Topping           `json:"topping" gorm:"many2many:cart-topping;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Qty           int                 `json:"qty"`
	SubAmount     int                 `json:"sub_amount" gorm:"type:int"`
}

type CartResponse struct {
	ID            int                 `json:"id"`
	ProductID     int                 `json:"product_id"`
	Product       Product             `json:"product"`
	TransactionID int                 `json:"-"`
	Transaction   TransactionResponse `json:"transaction" gorm:"foreignKey:TransactionID"`
	ToppingID     []int               `json:"topping_id" form:"topping_id" gorm:"-"`
	Topping       []Topping           `json:"topping" gorm:"many2many:cart-topping;"`
	SubAmount     int                 `json:"sub_amount"`
}

func (CartResponse) TableName() string {
	return "carts"
}
