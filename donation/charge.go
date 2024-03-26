package donation

import (
	"time"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

func createCharge(cfg Config, donate Donation) (err error) {
	client, _ := omise.NewClient(
		cfg.Webhhok,
		cfg.Secret,
	)

	token, create := &omise.Token{}, &operations.CreateToken{
		Name:            donate.Name,
		Number:          donate.CCNumber,
		ExpirationMonth: time.Month(donate.ExpMonth),
		ExpirationYear:  donate.ExpYear,
		SecurityCode:    donate.CVV,
	}
	if err = client.Do(token, create); err != nil {
		return
	}

	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   donate.AmountSubunits,
		Currency: "thb",
		Card:     token.ID,
	}

	if err = client.Do(charge, createCharge); err != nil {
		return
	}
	return
}
