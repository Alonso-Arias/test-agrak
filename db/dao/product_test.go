package dao

import (
	"context"
	"testing"

	"github.com/Alonso-Arias/test-agrak/db/model"
	"github.com/stretchr/testify/assert"
)

var productDao = NewProductDAO()

func TestFindAll_OK(t *testing.T) {

	result, err := productDao.FindAll(context.TODO())

	if err != nil {
		assert.FailNowf(t, "fails", "fails to gets products: %v", err)
	}

	t.Logf("Result : %v", result)

}

func TestGetBySku_OK(t *testing.T) {

	result, err := productDao.Get(context.TODO(), "FAL-8406270")

	if err != nil {
		assert.FailNowf(t, "fails", "fails to gets products: %v", err)
	}

	t.Logf("Result : %v", result)

}

func TestUpdate_OK(t *testing.T) {

	err := productDao.Update(context.TODO(), model.Product{Sku: "FAL-8406270", Name: "Alonso"})

	if err != nil {
		assert.FailNowf(t, "fails", "fails to update product: %v", err)
	}

}

func TestSave_OK(t *testing.T) {

	err := productDao.Save(context.TODO(), model.Product{Sku: "FAL-7777", Name: "Alonso", Brand: "Test", Size: "MM", Price: 1000, PrincipalImage: "url"})

	if err != nil {
		assert.FailNowf(t, "fails", "fails to update product: %v", err)
	}

}
