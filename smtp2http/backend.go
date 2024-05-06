package main

import (
	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"io"
	"log"
)

// The Backend implements SMTP server methods.
type Backend struct {
	Forward func(r io.Reader) error
	Auther  func(username, password string) error
}

// NewSession is called after client greeting (EHLO, HELO).
func (bkd *Backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &Session{b: bkd}, nil
}

// A Session is returned after successful login.
type Session struct {
	b *Backend
}

// AuthMechanisms returns a slice of available auth mechanisms; only PLAIN is
// supported in this example.
func (s *Session) AuthMechanisms() []string {
	return []string{sasl.Plain}
}

// Auth is the handler for supported authenticators.
func (s *Session) Auth(mech string) (sasl.Server, error) {
	return sasl.NewPlainServer(func(identity, username, password string) error {
		if s.b.Auther != nil {
			return s.b.Auther(username, password)
		}
		return nil
	}), nil
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	log.Println("Mail from:", from)
	return nil
}

func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {
	log.Println("Rcpt to:", to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	return s.b.Forward(r)
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}
