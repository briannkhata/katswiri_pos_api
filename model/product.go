package model

import "time"

type Product struct {
	ProductID         int       `json:"product_id" gorm:"column:product_id"`
	ProductName       string    `json:"product_name" gorm:"column:product_name"`
	Barcode           string    `json:"barcode" gorm:"column:barcode"`
	CategoryName      string    `json:"category_name" gorm:"column:category_name"`
	BrandName         string    `json:"brand_name" gorm:"column:brand_name"`
	SellingPrice      float64   `json:"selling_price" gorm:"column:selling_price"`
	PromoPrice        float64   `json:"promo_price" gorm:"column:promo_price"`
	OnPromotion       bool      `json:"on_promotion" gorm:"column:on_promotion"`
	ExpiryDate        string    `json:"expiry_date" gorm:"column:expiry_date"`
	DateAdded         time.Time `json:"date_added" gorm:"column:date_added"`
	AddedBy           int       `json:"added_by" gorm:"column:added_by"`
	ProductImage      string    `json:"product_image" gorm:"column:product_image"`
	Unit              string    `json:"unit" gorm:"column:unit"`
	UnitQty           float64   `json:"unit_qty" gorm:"column:unit_qty"`
	AvailableQuantity int       `json:"available_quantity" gorm:"column:available_quantity"`
	ShopID            int       `json:"shop_id" gorm:"column:shop_id"`
}
