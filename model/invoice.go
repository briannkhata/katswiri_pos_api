package model

import "time"

type Sale struct {
	ID            int       `gorm:"primary_key;autoIncrement" json:"sale_id"`
	UserID        int       `gorm:"not null" json:"user_id"`
	SaleDate      time.Time `gorm:"not null" json:"sale_date"`
	Tendered      float64   `gorm:"type:double" json:"tendered"`
	Change        float64   `gorm:"type:double" json:"change"`
	VAT           float64   `gorm:"type:double" json:"vat"`
	Total         float64   `gorm:"type:double" json:"total"`
	ShopID        *int      `gorm:"type:int" json:"shop_id"`
	SubTotal      float64   `gorm:"type:double" json:"sub_total"`
	ClientID      *int      `gorm:"type:int" json:"client_id"`
	PaymentTypeID *int      `gorm:"type:int" json:"payment_type_id"`
	Details       *string   `gorm:"type:varchar(300)" json:"details"`
	Balance       float64   `gorm:"type:double" json:"balance"`
	SaleType      *string   `gorm:"type:varchar(100)" json:"sale_type"`
}

type SaleDetail struct {
	ID            int        `gorm:"primary_key;autoIncrement" json:"sale_detail_id"`
	ProductID     int        `gorm:"not null" json:"product_id"`
	Price         float64    `gorm:"type:double" json:"price"`
	Qty           int        `gorm:"type:int" json:"qty"`
	VAT           float64    `gorm:"type:double" json:"vat"`
	Total         float64    `gorm:"type:double" json:"total"`
	SaleID        *int       `gorm:"type:int" json:"sale_id"`
	SaleDate      *time.Time `gorm:"type:datetime" json:"sale_date"`
	SubTotal      float64    `gorm:"type:double" json:"sub_total"`
	UserID        *int       `gorm:"type:int" json:"user_id"`
	ShopID        *int       `gorm:"type:int" json:"shop_id"`
	ClientID      *int       `gorm:"type:int" json:"client_id"`
	SaleType      *string    `gorm:"type:varchar(100)" json:"sale_type"`
	PaymentTypeID *int       `gorm:"type:int" json:"payment_type_id"`
}

type Quantity struct {
	ProductID int `gorm:"not null" json:"product_id"`
	Qty       int `gorm:"not null" json:"qty"`
	ShopID    int `gorm:"not null" json:"shop_id"`
}

func (Sale) TableName() string {
	return "tbl_sales"
}

func (SaleDetail) TableName() string {
	return "tbl_sale_details"
}

func (Quantity) TableName() string {
	return "tbl_quantities"
}
