package cards

import (
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
)

type Card struct {
	Secret   string
	Key      string
	Currency string
}

type Transaction struct {
	TransactionStatusID int
	Amount              int
	Currency            string
	LastFour            string
	BankReturnCode      string
}

func (c *Card) Charge(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	return c.CreatePaymentIntent(currency, amount)
}

func (c *Card) CreatePaymentIntent(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	stripe.Key = c.Secret

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64((int64(amount))),
		Currency: stripe.String(currency),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		msg := ""
		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = cardErrorMessage(stripeErr.Code)
		}
		return nil, msg, err
	}
	return pi, "", err
}

func cardErrorMessage(code stripe.ErrorCode) string {
	var msg = ""

	switch code {
	case stripe.ErrorCodeCardDeclined:
		msg = "Your card was declined"
	case stripe.ErrorCodeExpiredCard:
		msg = "Your carc is expired"
	default:
		msg = "Your Card "
	}
	return msg
}
