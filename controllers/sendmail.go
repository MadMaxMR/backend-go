package controllers

import (
	"bytes"      //para escribir informacion
	"crypto/tls" //para configurar las opciones de seguridad de tls
	"fmt"
	"html/template" // para tomar el archivo html y enviarlo
	"log"
	"net/mail" //para traer la estructura del un correo electrónico
	"net/smtp" //para enviar el correo
)

type Dest struct {
	Url   string
	Email string
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
func url(token string, email string) string {
	return fmt.Sprintf("http://localhost:8000/reset?tk=%s&ml=%s", token, email)
}

func SendMail(email string, token string) {
	from := mail.Address{Name: "Umachay restaurar contraseña", Address: "mad.pruebas.max@gmail.com"}
	to := mail.Address{Name: "User Email", Address: email}
	subject := "Restablecer contraseña"
	dest := Dest{Url: url(token, email), Email: to.Address}

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["Content-Type"] = `text/html; charset="UTF-8"`

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	t, err := template.ParseFiles("template.html")
	checkErr(err)

	buf := new(bytes.Buffer)
	err = t.Execute(buf, dest)
	checkErr(err)

	message += buf.String()

	servername := "smtp.gmail.com:465"
	host := "smtp.gmail.com"

	auth := smtp.PlainAuth("", "mad.pruebas.max@gmail.com", "981781002", host)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tlsConfig)
	checkErr(err)

	client, err := smtp.NewClient(conn, host)
	checkErr(err)

	err = client.Auth(auth)
	checkErr(err)

	err = client.Mail(from.Address)
	checkErr(err)

	err = client.Rcpt(to.Address)
	checkErr(err)

	w, err := client.Data()
	checkErr(err)

	_, err = w.Write([]byte(message))
	checkErr(err)

	err = w.Close()
	checkErr(err)

	client.Quit()
}
