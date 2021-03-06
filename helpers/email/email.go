package email

import (
	"errors"
	gomail "gopkg.in/mail.v2"
	"context"
)

func SendEmail(c context.Context, to string, subject string, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "daffashafwan.dev@gmail.com")
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)
	dial := gomail.NewDialer("smtp.gmail.com", 587, "daffashafwan.dev@gmail.com", "daffashafwan.dev.2021")
	err := dial.DialAndSend(msg)
	if err != nil {
		return errors.New("Gagal Mengirim Email")
	}

	return errors.New("Berhasil Kirim Email")
}
