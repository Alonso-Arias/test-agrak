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
	e.POST("/api/v1/product", productPost)
	e.GET("/api/v1/products/findAll", findAllProductsGet)
	e.GET("/api/v1/product/:sku", productGet)
	e.PUT("/api/v1/product", productPut)
	e.DELETE("/api/v1/product/:sku", productDelete)
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

// get product
// @Summary get product by sku
// @tags Product
// @Description obtiene producto por sku
// @ID productGet
// @Accept  json
// @Produce  json
// @Param sku path string true "Sku"
// @Success 200  {object} product.GetProductResponse
// @Failure 404 {object}  errors.CustomError
// @Failure 500 {object}  errors.CustomError
// @Router /product/{sku} [get]
func productGet(c echo.Context) error {

	req := product.GetProductRequest{
		Sku: c.Param("sku"),
	}

	res, err := product.ProductService{}.GetProduct(context.TODO(), req)
	if ce, ok := err.(errs.CustomError); ok {
		return c.JSON(ce.Code, err)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}

// delete product
// @Summary delete product by sku
// @tags Product
// @Description elimina un producto
// @ID productDelete
// @Accept  json
// @Produce  json
// @Param sku path string true "Sku"
// @Success 200  {object} product.DeleteProductResponse
// @Failure 404 {object}  errors.CustomError
// @Failure 500 {object}  errors.CustomError
// @Router /product/{sku} [delete]
func productDelete(c echo.Context) error {

	req := product.DeleteProductRequest{
		Sku: c.Param("sku"),
	}

	res, err := product.ProductService{}.DeleteProduct(context.TODO(), req)
	if ce, ok := err.(errs.CustomError); ok {
		return c.JSON(ce.Code, err)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}

// update product
// @Summary update product by sku
// @tags Product
// @Description actualiza un producto
// @ID productPut
// @Accept  json
// @Produce  json
// @Param UpdateProductRequest body product.UpdateProductRequest true "Product"
// @Success 200  {object} product.UpdateProductResponse
// @Failure 404 {object}  errors.CustomError
// @Failure 500 {object}  errors.CustomError
// @Router /product [put]
func productPut(c echo.Context) error {

	log := loggerf.WithField("func", "productPut")

	req := product.UpdateProductRequest{}

	if err := c.Bind(req); err != nil {
		log.WithError(err).Error("Binding error")
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := product.ProductService{}.UpdateProduct(context.TODO(), req)
	if ce, ok := err.(errs.CustomError); ok {
		return c.JSON(ce.Code, err)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}

// save product
// @Summary save product
// @tags Product
// @Description guarda un producto
// @ID productPost
// @Accept  json
// @Produce  json
// @Param SaveProductRequest body product.SaveProductRequest true "Product"
// @Success 200  {object} product.SaveProductResponse
// @Failure 404 {object}  errors.CustomError
// @Failure 500 {object}  errors.CustomError
// @Router /product [post]
func productPost(c echo.Context) error {

	log := loggerf.WithField("func", "productPost")

	req := product.SaveProductRequest{}

	if err := c.Bind(req); err != nil {
		log.WithError(err).Error("Binding error")
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := product.ProductService{}.SaveProduct(context.TODO(), req)
	if ce, ok := err.(errs.CustomError); ok {
		return c.JSON(ce.Code, err)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
