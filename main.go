package main

import (
	"encoding/json"
	"fmt"
	"log"

	"xendit-test/config"
	inv "xendit-test/invoice"
	"xendit-test/obj"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// load environment variables
	config, err := config.LoadEnv(".")
	if err != nil {
		log.Fatal(err)
	}

	// create general invoices
	// every payment method are included as default
	fmt.Printf("\nCREATE GENERAL INVOICE\n")
	invoiceObject := obj.InvoiceObject{
		ID:          "1",
		SuccessUrl:  "https://google.com",
		FailUrl:     "https://example.com",
		Amount:      12000,
		Name:        "Francisco",
		Email:       "slashschtye252@gmail.com",
		Description: "Pulsa 10.000",
		Currency:    "IDR",
	}

	id := inv.CreateInvoice(config.WriteKey, invoiceObject)
	fmt.Println("SUCCESS CREATE INVOICE")

	inv.GetInvoice(config.ReadKey, id)

	// initiate echo
	e := echo.New()

	// middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// get callback from xendit for updated payment status with ewallet
	// put callback url with your root url + below path in xendit callbacks url dashboard
	e.POST("/callbacks/ewallet", func(c echo.Context) error {
		response := map[string]interface{}{}
		// decode request body from xendit
		err := json.NewDecoder(c.Request().Body).Decode(&response)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("response callbacks update status ewallet =", response)

		return nil
	})

	// get callback from xendit for created FIXED VIRTUAL ACCOUNT
	// put callback url with your root url + below path in xendit callbacks url dashboard
	e.POST("/callbacks/virtual-created", func(c echo.Context) error {
		response := map[string]interface{}{}
		// decode request body from xendit
		err := json.NewDecoder(c.Request().Body).Decode(&response)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("response callbacks created virtual account =", response)

		return nil
	})

	// get callback from xendit for paid FIXED VIRTUAL ACCOUNT
	// put callback url with your root url + below path in xendit callbacks url dashboard
	e.POST("/callbacks/virtual-paid", func(c echo.Context) error {
		response := map[string]interface{}{}
		// decode request body from xendit
		err := json.NewDecoder(c.Request().Body).Decode(&response)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("response callbacks paid virtual account =", response)

		return nil
	})

	// get callback from xendit for paid general invoice
	// put callback url with your root url + below path in xendit callbacks url dashboard
	e.POST("/callbacks/invoice", func(c echo.Context) error {
		response := map[string]interface{}{}
		// decode request body from xendit
		err := json.NewDecoder(c.Request().Body).Decode(&response)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("response callbacks paid invoice =", response)

		return nil
	})

	e.Start("127.0.0.1:8000")
}
