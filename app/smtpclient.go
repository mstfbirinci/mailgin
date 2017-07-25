package app

import gomail "gopkg.in/gomail.v2"

//SendMail sends the mail
func SendMail(msg *Message) {
	dialer := CreateDialer()
	message := gomail.NewMessage()

	message.SetHeader("From", msg.From)
	message.SetHeader("To", msg.Recipients...)
	message.SetHeader("Subject", msg.Subject)
	message.SetBody(msg.ContentType, msg.Content)

	if err := dialer.DialAndSend(message); err != nil {
		panic(err)
	}

}

//CreateDialer create an SMTP client instance.
func CreateDialer() *gomail.Dialer {
	smtp := ConfigurationInstance().Smtp

	dialer := gomail.NewDialer(smtp.SmtpServer, smtp.Port, smtp.User, smtp.Pass)

	return dialer
}
