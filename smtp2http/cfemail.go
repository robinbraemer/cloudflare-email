package main

import (
	"context"
	"fmt"
	"github.com/rsjethani/rsling"
	"net/http"
)

type Contact struct {
	Address string `json:"address"`
	Name    string `json:"name,omitempty"`
}

type Email struct {
	To      []Contact `json:"to"`
	ReplyTo []Contact `json:"replyTo,omitempty"`
	CC      []Contact `json:"cc,omitempty"`
	BCC     []Contact `json:"bcc,omitempty"`
	From    Contact   `json:"from"`
	Subject string    `json:"subject"`
	Text    string    `json:"text,omitempty"`
	HTML    string    `json:"html,omitempty"`
}

func SendEmail(ctx context.Context, email Email, url, token string) error {
	req, err := rsling.New().
		Base(url).Post("").
		Set("Authorization", token).
		BodyJSON(email).
		RequestWithContext(ctx)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d %s", res.StatusCode, res.Status)
	}
	return nil
}
