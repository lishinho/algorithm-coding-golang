package main

import (
	"crypto/tls"
	"go-common/library/log"
	"gopkg.in/gomail.v2"
)

func sendMail() {
	m := gomail.NewMessage()
	m.SetHeader("From", "lizhenbang0924@outlook.com")
	m.SetHeader("To", "lishinho95@proton.me")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!还有个附件给你")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	m.Attach("/Users/lishinho/Downloads/Union.png")

	d := gomail.NewDialer("smtp.office365.com", 587, "lizhenbang0924@outlook.com", "*****")
	//d.SSL = true
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	log.Info("success!")
}
