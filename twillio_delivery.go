package delivery

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

type TwillioDelivery struct {
	accountSid string
	authToken  string
	fromPhone  string
}

func NewTwillioDelivery(accountSid string, authToken string, fromPhone string) *TwillioDelivery {
	return &TwillioDelivery{accountSid: accountSid, authToken: authToken, fromPhone: fromPhone}
}

func (delivery *TwillioDelivery) Send(ctx context.Context, destination string, message string)error {
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + delivery.accountSid + "/Messages.json"

	// Build out the data for our message
	v := url.Values{}
	v.Set("To", destination)
	v.Set("From", delivery.fromPhone)
	body := message
	v.Set("Body", body)

	rb := *strings.NewReader(v.Encode())

	// Create client
	client := &http.Client{}


	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(delivery.accountSid, delivery.authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	_, err := client.Do(req.WithContext(ctx))
	return err
}
