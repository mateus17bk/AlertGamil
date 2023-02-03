/**
Ao executar esse projeto no terminal passe o export do da variável GMAIL_PASSWORD
que é a chave de autenticação do google email para receber a mensagem tem que habilitar o acesso a app menos seguro
**/

package main

import "alertmanager/email"

func main() {
	email.SendEmail([]string{"<SEU_EMAIL>"},
		"Alerta de servidor down", "Google", "Erro ao conectar ao servidor", "06/05/2016 19:19", "./email/template.html")

}
