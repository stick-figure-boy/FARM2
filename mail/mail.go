package mail

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

func sendEmail(subject, body string, receiver []string) error {
	smtpServer := fmt.Sprintf("%s:%s", os.Getenv("MAIL_HOST"), os.Getenv("MAIL_PORT"))
	auth := smtp.CRAMMD5Auth(os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASSWORD"))
	msg := []byte(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(receiver, ","), subject, body))

	if err := smtp.SendMail(smtpServer, auth, os.Getenv("MAIL_SENDAR"), receiver, msg); err != nil {
		return err
	}
	return nil
}
