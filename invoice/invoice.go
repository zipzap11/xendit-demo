package invoice

import (
	"fmt"
	"log"

	"xendit-test/obj"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

func CreateInvoice(writekey string, inv obj.InvoiceObject) string {
	xendit.Opt.SecretKey = writekey
	shouldSendEmail := true
	data := invoice.CreateParams{
		ExternalID:         inv.ID,
		Amount:             inv.Amount,
		SuccessRedirectURL: inv.SuccessUrl,
		FailureRedirectURL: inv.FailUrl,
		Currency:           inv.Currency,
		PayerEmail:         inv.Email,
		ShouldSendEmail:    &shouldSendEmail,
		Description:        inv.Description,
		CustomerNotificationPreference: xendit.InvoiceCustomerNotificationPreference{
			InvoiceCreated:  []string{"email"},
			InvoiceReminder: []string{"email"},
			InvoicePaid:     []string{"email"},
			InvoiceExpired:  []string{"email"},
		},
		Customer: xendit.InvoiceCustomer{
			GivenNames: inv.Name,
			Email:      inv.Email,
		},
	}

	resp, err := invoice.Create(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created invoice: %+v\n", resp)
	return resp.ID
}

func GetInvoice(readkey string, id string) {
	xendit.Opt.SecretKey = readkey

	data := invoice.GetParams{
		ID: id,
	}

	resp, err := invoice.Get(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved invoice: %+v\n", resp)
}
