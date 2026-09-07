package main

import (
	"context"
	"fmt"

	"github.com/resend/resend-go/v4"
)

type SendEmail struct {
	Name        string
	Email       string
	Phone       string
	Details     string
	PlanChoices string
}

func sendEmail() {
	ctx := context.TODO()
	client := resend.NewClient("re_xxxxxxxxx")

	params := &resend.SendEmailRequest{
		From:    "Acme <onboarding@resend.dev>",
		To:      []string{"delivered@resend.dev"},
		Subject: "hello world",
		Html:    "<p>it works!</p>",
		ReplyTo: "onboarding@resend.dev",
	}

	sent, err := client.Emails.SendWithContext(ctx, params)

	if err != nil {
		panic(err)
	}
	fmt.Println(sent.Id)
}
