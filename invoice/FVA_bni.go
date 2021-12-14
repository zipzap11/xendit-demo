package invoice

import (
	"fmt"
	"log"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/virtualaccount"
)

func CreateVA(bankcode string, writekey string) string {
	xendit.Opt.SecretKey = writekey

	data := virtualaccount.CreateFixedVAParams{
		ExternalID: "demo-1475804036622",
		BankCode:   bankcode,
		Name:       "Rika Sutanto",
	}

	resp, err := virtualaccount.CreateFixedVA(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created fixed va: %+v\n", resp)
	return resp.ID
}

func GetVA(VAid string, readkey string) {
	xendit.Opt.SecretKey = readkey

	data := virtualaccount.GetFixedVAParams{
		ID: "59e03a976fab8b1850fdf347",
	}

	resp, err := virtualaccount.GetFixedVA(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved fixed va: %+v\n", resp)
}
