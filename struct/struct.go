package main

import "github.com/xendit/xendit-go"

type CreateEWalletChargeParams struct {
	ForUserID         string                     `json:"-"`
	WithFeeRule       string                     `json:"-"`
	ReferenceID       string                     `json:"reference_id" validate:"required"`
	Currency          string                     `json:"currency" validate:"required"`
	Amount            float64                    `json:"amount" validate:"required"`
	CheckoutMethod    string                     `json:"checkout_method" validate:"required"`
	ChannelCode       string                     `json:"channel_code,omitempty"`
	ChannelProperties map[string]string          `json:"channel_properties,omitempty"`
	CustomerID        string                     `json:"customer_id,omitempty"`
	Basket            []xendit.EWalletBasketItem `json:"basket,omitempty"`
	Metadata          map[string]interface{}     `json:"metadata,omitempty"`
}
