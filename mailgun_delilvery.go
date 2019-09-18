package delivery

import (
	"context"
	"github.com/mailgun/mailgun-go/v3"
)


type MailgunDelivery struct {
	impl  *mailgun.MailgunImpl
	from  string
}

func NewMailgunDelivery(domain string, apiKey string, from string) *MailgunDelivery {
	mg := mailgun.NewMailgun(domain, apiKey)
	mg.SetAPIBase("https://api.eu.mailgun.net/v3")
	return &MailgunDelivery{impl:mg, from:from}
}


func (m *MailgunDelivery) Send(ctx context.Context, destination, message string) error {
	msg := m.impl.NewMessage(m.from, "", message, destination)
	_, _, err := m.impl.Send(ctx, msg)
	return err
}


func (m *MailgunDelivery)SendTemplate(ctx context.Context, destination, subject, template string, variables map[string]string)error{
	msg := m.impl.NewMessage(m.from, subject, "", destination)
	msg.SetTemplate(template)
	for key, value := range variables {
		err := msg.AddVariable(key, value);
		if err != nil{
			panic(err)
		}
	}
	_, _, err := m.impl.Send(ctx, msg)
	return err
}



