package main

import (
	"encoding/json"
	"fmt"
	"log"

	"xendit-test/config"
	inv "xendit-test/invoice"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

func main() {
	// load environment variables
	config, err := config.LoadEnv(".")
	if err != nil {
		log.Fatal(err)
	}

	// shopeepay charge invoice
	fmt.Printf("\nSHOPEEPAY INVOICE\n")
	id := inv.CreateShopeepayCharge(config.WriteKey)
	inv.GetEwalletCharge(id, config.ReadKey)

	// ovo charge invoice
	fmt.Printf("\nOVO INVOICE\n")
	id = inv.CreateOvoCharge(config.WriteKey)
	inv.GetEwalletCharge(id, config.ReadKey)

	// initiate echo
	e := echo.New()

	// middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// get callback from xendit
	// put callback url with your root url + below path in xendit callbacks url dashboard
	e.POST("/callbacks/ewallet", func(c echo.Context) error {
		response := map[string]interface{}{}
		// decode request body from xendit
		err := json.NewDecoder(c.Request().Body).Decode(&response)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("response callbacks =", response)

		return nil
	})

	e.Start("127.0.0.1:8000")
}

func createInvoice() {
	xendit.Opt.SecretKey = "xnd_development_FpeM3C727hIk1e1XxjiwaTBDrogSpOEVklHYMr4tEoTcY8JhQgwtuxUb1yFDLZ"
	data := invoice.CreateParams{
		ExternalID: "demo_1475801962607",
		Amount:     50000,
	}

	resp, err := invoice.Create(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created invoice: %+v\n", resp)
}

func getInvoice(id string) {
	xendit.Opt.SecretKey = "xnd_development_P4qDfOss0OCpl8RtKrROHjaQYNCk9dN5lSfk+R1l9Wbe+rSiCwZ3jw=="

	data := invoice.GetParams{
		ID: id,
	}

	resp, err := invoice.Get(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved invoice: %+v\n", resp)
}
