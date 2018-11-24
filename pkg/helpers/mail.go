package helper

import (
	"fmt"
	"net/smtp"
)

func SendMail(c chan string, email string, body string) string {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "jainatishay.j@gmail.com", "tqmswbmnzntftwpb", "smtp.gmail.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{email}
	msg := []byte("To: " + email + "\r\n" +
		"Subject: Go-Social!\r\n" +
		"\r\n" +
		"" + body + ".\r\n")
	err := smtp.SendMail("smtp.gmail.com:25", auth, "jainatishay.j@gmail.com", to, msg)
	if err != nil {
		fmt.Println(err)
		return "Mail not sent!"
	}
	return "Mail sent."
}
