package main

import (
	"log"
	"net/smtp"
)

func main() {

	//Criamos um slice do tipo string do tamanho máximo de 1 para receber nosso e-mail destinatário.
	recipients := make([]string, 1)
	recipients[0] = "email@dodestinatario.com"

	err := smtp.SendMail(
		/* endereço do servidor de SMTP */ "smtp.gmail.com:25",
		/* mecanismo de autenticação*/ smtp.PlainAuth("", "seuemail@gmail.com", "suasenhagmail", "smtp.gmail.com"),
		/* e-mail de origem */ "seuemail@gmail.com",
		/*Mensagem no RFC 822-style*/ recipients,
		/* Corpo da mensagem */ []byte("Subject:Olá!\n\n Olá Fulano. Tudo de bom com Go!"))
	if err != nil {
		log.Fatal(err)
	}

}
