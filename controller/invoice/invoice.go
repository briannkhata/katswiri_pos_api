package controller

import (
	"errors"
	"fmt"
	"time"

	"github.com/briannkhata/katswiri_api/database"
	helper "github.com/briannkhata/katswiri_api/helper"
	invoice "github.com/briannkhata/katswiri_api/model"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

var payload struct {
	UserID        int       `json:"user_id"`
	ProductIDs    []int     `json:"product_ids"`
	Qtys          []int     `json:"qtys"`
	VATs          []float64 `json:"vats"`
	Prices        []float64 `json:"prices"`
	VAT           float64   `json:"vat"`
	SubTotal      float64   `json:"sub_total"`
	Total         float64   `json:"total"`
	Tendered      float64   `json:"tendered"`
	PaymentTypeID int       `json:"payment_type_id"`
	Details       string    `json:"details"`
	SaleType      string    `json:"sale_type"`
	ShopId        int       `json:"shop_id"`
	ClientID      int       `json:"client_id"`
}

func SubmitInvoice(c *fiber.Ctx) error {
	db := database.DBConn

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"msg":     "Invalid input",
			"data":    "json: cannot unmarshal string into Go struct field payload.exampleField of type int",
		})
	}

	sale := invoice.Sale{
		UserID:        payload.UserID,
		SaleDate:      time.Now(),
		Tendered:      payload.Tendered,
		Change:        payload.Tendered - payload.Total,
		VAT:           payload.VAT,
		Total:         payload.Total,
		SubTotal:      payload.SubTotal,
		PaymentTypeID: helper.IntPointer(payload.PaymentTypeID),
		ClientID:      helper.IntPointer(payload.ClientID),
		Details:       helper.StringPointer(payload.Details),
		SaleType:      helper.StringPointer(payload.SaleType),
		ShopID:        helper.IntPointer(payload.ShopId),
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&sale).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"msg":     "Failed to create sale",
				"data":    err.Error(),
			})
		}

		for i, productID := range payload.ProductIDs {
			saleDetail := invoice.SaleDetail{
				ProductID:     productID,
				Qty:           payload.Qtys[i],
				VAT:           payload.VATs[i],
				Price:         payload.Prices[i],
				SaleID:        helper.IntPointer(sale.ID),
				SaleDate:      &sale.SaleDate,
				SaleType:      sale.SaleType,
				PaymentTypeID: sale.PaymentTypeID,
				ShopID:        sale.ShopID,
				UserID:        &sale.UserID,
				ClientID:      sale.ClientID,
			}

			if err := tx.Create(&saleDetail).Error; err != nil {
				return err
			}

			if err := tx.Model(&invoice.Quantity{}).
				Where("product_id = ? AND shop_id = ?", productID, payload.ShopId).
				UpdateColumn("qty", gorm.Expr("qty - ?", payload.Qtys[i])).Error; err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"success": false,
					"msg":     "Failed to update quantity",
					"data":    err.Error(),
				})
			}

		}

		return nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"msg":     "Transaction failed",
			"data":    err.Error(),
		})
	}

	saleDetails, err := GetSaleDetail(c, sale.ID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":     "sale created successfully",
		"success": true,
		"data":    saleDetails,
	})

}

func GetSaleDetail(c *fiber.Ctx, saleId int) ([]invoice.SaleDetail, error) {

	db := database.DBConn
	var saleDetail []invoice.SaleDetail

	if err := db.Where("sale_id = ?", saleId).Find(&saleDetail).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Sale not found",
			})
		}
		return nil, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to retrieve sale details",
		})
	}

	if len(saleDetail) == 0 {
		return nil, c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Sale details not found",
		})
	}

	return saleDetail, nil
}

func GetProductPrice(db *gorm.DB, productId int) (float64, error) {
	var price float64
	query := `
		SELECT 
			price 
		FROM 
			tbl_sale_details 
		WHERE 
			product_id = ?  
		LIMIT 1 
	`

	err := db.Raw(query, productId).Scan(&price).Error
	if err != nil {
		return 0, err
	}

	if price == 0 {
		return 0, fmt.Errorf("price not found for the product with ID %d", productId)
	}

	return price, nil
}
