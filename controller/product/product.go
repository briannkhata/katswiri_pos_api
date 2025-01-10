package controller

import (
	"github.com/briannkhata/katswiri_pos_api/database"
	product "github.com/briannkhata/katswiri_pos_api/model"
	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	db := database.DBConn
	var products []product.Product
	query := `
		SELECT 
			p.product_id,
			p.name AS product_name,
			p.barcode,
			c.category AS category_name,
			b.brand_name,
			p.selling_price,
			p.promo_price,
			p.on_promotion,
			p.expiry_date,
			p.date_added,
			p.added_by,
			p.image AS product_image,
			u.unit_type AS unit,
			u.qty AS unit_qty,
			q.qty AS available_quantity,
			q.shop_id
		FROM 
			tbl_products p
		LEFT JOIN 
			tbl_category c ON p.category_id = c.category_id
		LEFT JOIN 
			tbl_brands b ON p.brand_id = b.brand_id
		LEFT JOIN 
			tbl_units u ON p.unit_id = u.unit_id
		LEFT JOIN 
			tbl_quantities q ON p.product_id = q.product_id
	`
	if err := db.Raw(query).Scan(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to retrieve products",
			"data":    nil,
		})
	}

	if len(products) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": "No Product Found",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Products retrieved successfully",
		"data":    products,
	})
}

func GetProduct(c *fiber.Ctx) error {
	db := database.DBConn
	var product product.Product
	productID := c.Params("id")

	query := `
		SELECT 
			p.product_id,
			p.name AS product_name,
			p.barcode,
			c.category AS category_name,
			b.brand_name,
			p.selling_price,
			p.promo_price,
			p.on_promotion,
			p.expiry_date,
			p.date_added,
			p.added_by,
			p.image AS product_image,
			u.unit_type AS unit,
			u.qty AS unit_qty,
			q.qty AS available_quantity,
			q.shop_id
		FROM 
			tbl_products p
		LEFT JOIN 
			tbl_category c ON p.category_id = c.category_id
		LEFT JOIN 
			tbl_brands b ON p.brand_id = b.brand_id
		LEFT JOIN 
			tbl_units u ON p.unit_id = u.unit_id
		LEFT JOIN 
			tbl_quantities q ON p.product_id = q.product_id
		WHERE 
			p.product_id = ?
	`
	if err := db.Raw(query, productID).Scan(&product).Error; err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": "Product not found",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Product retrieved successfully",
		"data":    product,
	})
}

func GetProductBarcode(c *fiber.Ctx) error {
	db := database.DBConn
	var product product.Product
	productID := c.Params("id")

	query := `
		SELECT 
			p.product_id,
			p.barcode,
			p.selling_price,
			p.promo_price,
			p.on_promotion,
			p.expiry_date,
			u.qty AS unit_qty,
			q.qty AS available_quantity,
			q.shop_id
		FROM 
			tbl_products p
		LEFT JOIN 
			tbl_category c ON p.category_id = c.category_id
		LEFT JOIN 
			tbl_brands b ON p.brand_id = b.brand_id
		LEFT JOIN 
			tbl_units u ON p.unit_id = u.unit_id
		LEFT JOIN 
			tbl_quantities q ON p.product_id = q.product_id
		WHERE 
			p.product_id = ?
	`
	if err := db.Raw(query, productID).Scan(&product).Error; err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": "Product not found",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Product retrieved successfully",
		"data":    product,
	})
}
