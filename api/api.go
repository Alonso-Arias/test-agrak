package main

import (
	"context"
	"net/http"

	_ "github.com/Alonso-Arias/test-agrak/api/docs"

	errs "github.com/Alonso-Arias/test-agrak/errors"
	"github.com/Alonso-Arias/test-agrak/log"
	"github.com/Alonso-Arias/test-agrak/services/product"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var loggerf = log.LoggerJSON().WithField("package", "main")

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /api/v1
func main() {
	e := echo.New()
	e.POST("/api/v1/product", findAllProductsGet)
	e.GET("/api/v1/products/findAll", findAllProductsGet)
	e.GET("/api/v1/product/:sku", findAllProductsGet)
	e.PUT("/api/v1/product/:sku", findAllProductsGet)
	e.DELETE("/api/v1/product/:sku", findAllProductsGet)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":1323"))

}

// find all products
// @Summary Find all products
// @tags Products
// @Description obtiene todos los productos
// @ID findAllProductsGet
// @Accept  json
// @Produce  json
// @Success 200  {object} product.FindAllProductsResponse
// @Failure 404 {object}  errors.CustomError
// @Failure 500 {object}  errors.CustomError
// @Router /products/findAll [get]
func findAllProductsGet(c echo.Context) error {

	res, err := product.ProductService{}.FindAllProducts(context.TODO())
	if ce, ok := err.(errs.CustomError); ok {
		return c.JSON(ce.Code, err)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
