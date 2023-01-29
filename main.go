package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

func main() {
	from := "mateusguedes17@gmail.com"
	password := os.Getenv("GMAIL_PASSWORD")
	if password == "" {
		panic("GMAIL_PASSWORD não foi configurado")
	}
	to := []string{
		"mateusguedes17@gmail.com",
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	t, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-Version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Alerta:: Servidor down \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Server  string
		Error   string
		Horario string
	}{
		Server:  "Google",
		Error:   "Erro ao acessar o servidor",
		Horario: "28/01/2023 22:00",
	})

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println("Erro ao enviar o email: %s", err)
	}
	fmt.Println("Email enviado com sucesso!")
}
