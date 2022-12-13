package base

import (
	"log"
	"time"

	"testing"

	m "github.com/Alonso-Arias/test-agrak/db/model"
)

func TestGetConnection(t *testing.T) {

	dbc := GetDB()

	result := m.Product{}

	dbc.Raw("SELECT * FROM PRODUCT").Scan(&result)

}

func TestGetTime(t *testing.T) {

	loc, _ := time.LoadLocation("Europe/Monaco")
	//set timezone,
	savetrxTime := time.Now().In(loc)

	log.Println("Hora  : ", savetrxTime)

}
