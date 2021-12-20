package email

import (
	"net/http"

	gomail "gopkg.in/mail.v2"

	"github.com/daffashafwan/pointcuan/helpers/response"
	"github.com/labstack/echo/v4"
)

func SendEmail(c echo.Context, to string, subject string, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "daffashafwan.dev@gmail.com")
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)
	dial := gomail.NewDialer("smtp.gmail.com", 587, "daffashafwan.dev@gmail.com", "daffashafwan.dev.2021")
	err := dial.DialAndSend(msg)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	// // sender data
	// from := "daffashafwan.dev@gmail.com"
	// password := "daffashafwan.dev.2021"   // ex: "ieiemcjdkejspqz"
	// // receiver address
	// toEmail := to // ex: "Jane.Smith@yahoo.com"
	// recipient := []string{toEmail}
	// // smtp - Simple Mail Transfer Protocol
	// host := "smtp.gmail.com"
	// port := "587"
	// address := host + ":" + port
	// message := []byte(subject + body)
	// // athentication data
	// // func PlainAuth(identity, username, password, host string) Auth
	// auth := smtp.PlainAuth("", from, password, host)
	// // send mail
	// // func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	// err := smtp.SendMail(address, auth, from, recipient, message)

	return response.SuccessResponse(c, http.StatusOK, "Sukses Kirim Email")
}
