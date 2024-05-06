package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/mnako/letters"
	"io"
	"log"
	"time"

	"github.com/emersion/go-smtp"
)

type (
	Args struct {
		Addr      string `arg:"--addr,env:ADDR" default:":1587"`
		User      string `arg:"required,--user,env:USER"`
		Pass      string `arg:"required,--pass,env:PASS"`
		PostURL   string `arg:"required,--post-url,env:POST_URL"`
		PostToken string `arg:"required,--post-token,env:POST_TOKEN"`
	}
)

func main() {
	var args Args
	arg.MustParse(&args)

	be := &Backend{
		Auther: func(username, password string) error {
			if username != args.User || password != args.Pass {
				return errors.New("invalid username or password")
			}
			return nil
		},
		Forward: func(r io.Reader) error {
			// Parse email
			email, err := letters.ParseEmail(r)
			if err != nil {
				return fmt.Errorf("error parsing email: %w", err)
			}

			// Convert email to web service format
			e, err := ConvertEmail(email)
			if err != nil {
				return fmt.Errorf("error converting email: %w", err)
			}

			ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
			defer cancel()

			err = SendEmail(ctx, *e, args.PostURL, args.PostToken)
			if err != nil {
				return fmt.Errorf("error sending email: %w", err)
			}
			return nil
		},
	}

	s := smtp.NewServer(be)

	s.Addr = args.Addr
	s.Domain = "localhost"
	s.WriteTimeout = 10 * time.Second
	s.ReadTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	log.Println("Starting server at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
