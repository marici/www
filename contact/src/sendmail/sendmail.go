package sendmail

import (
	"fmt"
	"log"
	"bytes"
	"net/smtp"
	"text/template"
)

type SMTPServer struct {
	Host string
	Port int
	UserName string
	Password string
}

type Email struct {
	From string
	To string
	Subject string
	Data interface{}
}

func BuildBody(tmplname string, email Email) (buffer *bytes.Buffer) {
	tmpl, _ := template.ParseFiles(tmplname)
	buffer = new(bytes.Buffer)
	tmpl.Execute(buffer, email)
	return buffer
}

func (server SMTPServer) SendMail(
		from string, to string, subject string, tmplname string,
		data interface{}) (err error) {
	c, err := smtp.Dial(fmt.Sprintf("%s:%d", server.Host, server.Port))
	if err != nil {
		log.Fatal(err)
	}
	c.Mail(from)
	c.Rcpt(to)
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()

	email := Email{from, to, subject, data}
	buffer := BuildBody(tmplname, email)
	_, err = buffer.WriteTo(wc)
	c.Quit()

	return err
}

func (server SMTPServer) SendMailWithAuth(
		from string, to string, subject string, tmplname string,
		data interface{}) (err error) {
	email := Email{from, to, subject, data}
	buffer := BuildBody(tmplname, email)
	auth := smtp.PlainAuth("", server.UserName, server.Password, server.Host)
	connectTo := fmt.Sprintf("%s:%d", server.Host, server.Port)
	err = smtp.SendMail(
		connectTo,
		auth,
		server.UserName,
		[]string{to},
		buffer.Bytes())
	return err
}
