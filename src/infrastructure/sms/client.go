package sms

import (
	"context"
	"log"
)

type Client interface {
	SendMessage(ctx context.Context, phoneNumber, msg string) error
}

type clientImpl struct {
	apiKey string
}

func NewClient(apiKey string) Client {
	return &clientImpl{
		apiKey: apiKey,
	}
}

func (c *clientImpl) SendMessage(ctx context.Context, phoneNumber, msg string) error {
	log.Println("sending sms to", phoneNumber, "with message ", msg)
	return nil
}
