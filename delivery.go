package delivery

import "context"

type PlainTextDelivery interface {
	Send(ctx context.Context, destination, message string)error
}


type EmailTemplateDelivery interface {
	SendTemplate(ctx context.Context, destination, subject, template string, variables map[string]string)error
}

