package router

import (
	invoice "github.com/briannkhata/katswiri_pos_api/controller/invoice"
	product "github.com/briannkhata/katswiri_pos_api/controller/product"
	auth "github.com/briannkhata/katswiri_pos_api/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {

	app.Get("/api/v1/product", auth.AuthenticateRoutes(), product.GetProducts)
	app.Get("/api/v1/product/:id", auth.AuthenticateRoutes(), product.GetProduct)
	app.Get("/api/v1/product/:id", auth.AuthenticateRoutes(), product.GetProductBarcode)
	app.Post("/api/v1/invoice/", auth.AuthenticateRoutes(), invoice.SubmitInvoice)

}
