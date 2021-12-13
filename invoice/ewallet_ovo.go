package invoice

import (
	"fmt"
	"log"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/ewallet"
)

// create e wallet charge > pay with shopeepay return the id of created invoice
func CreateOvoCharge(writekey string) string {
	xendit.Opt.SecretKey = writekey

	// optional cart items
	// ewalletBasketItem := xendit.EWalletBasketItem{
	// 	ReferenceID: "basket-product-ref-id",
	// 	Name:        "product name",
	// 	Category:    "mechanics",
	// 	Currency:    "IDR",
	// 	Price:       50000,
	// 	Quantity:    5,
	// 	Type:        "type",
	// 	SubCategory: "subcategory",
	// 	Metadata: map[string]interface{}{
	// 		"meta": "data",
	// 	},
	// }

	data := ewallet.CreateEWalletChargeParams{
		ReferenceID:    "test-reference-id",
		Currency:       "IDR",
		Amount:         1688,
		CheckoutMethod: "ONE_TIME_PAYMENT",
		ChannelCode:    "ID_OVO",
		ChannelProperties: map[string]string{
			// home page url
			// "success_redirect_url": "https://example.com",
			"mobile_number": "+6281239812420",
		},
		// Basket: []xendit.EWalletBasketItem{
		// 	ewalletBasketItem,
		// },
		Metadata: map[string]interface{}{
			"meta": "data",
		},
	}

	charge, chargeErr := ewallet.CreateEWalletCharge(&data)
	if chargeErr != nil {
		log.Fatal(chargeErr)
	}
	fmt.Printf("created e-wallet charge: %+v\n", charge)
	return charge.ID
}
