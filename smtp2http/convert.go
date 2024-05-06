package main

import (
	"errors"
	"github.com/mnako/letters"
	"net/mail"
)

func ConvertEmail(e letters.Email) (*Email, error) {
	if len(e.Headers.From) == 0 {
		return nil, errors.New("email has no 'From' address")
	}
	return &Email{
		To:      Map(e.Headers.To, AddressToContact),
		ReplyTo: Map(e.Headers.ReplyTo, AddressToContact),
		CC:      Map(e.Headers.Cc, AddressToContact),
		BCC:     Map(e.Headers.Bcc, AddressToContact),
		From:    AddressToContact(e.Headers.From[0]), // Assuming there's always at least one "From" contact
		Subject: e.Headers.Subject,
		Text:    e.Text,
		HTML:    e.HTML,
	}, nil
}

func AddressToContact(addr *mail.Address) Contact {
	return Contact{
		Address: addr.Address,
		Name:    addr.Name,
	}
}

func Map[I, O any](input []I, f func(I) O) []O {
	output := make([]O, len(input))
	for i, v := range input {
		output[i] = f(v)
	}
	return output
}
