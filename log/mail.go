package log

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"
)

var (
	from     = "120735581@qq.com"
	pass     = "peerdmnoqirqbiaa"
	smtpHost = "smtp.qq.com"
)

func sendMail(subject, body string, to []string) error {
	auth := smtp.PlainAuth("", from, pass, smtpHost)

	conn, err := tls.Dial("tcp", smtpHost+":465", nil)
	if err != nil {
		return err
	}
	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return err
	}
	if err = client.Auth(auth); err != nil {
		return err
	}
	if err = client.Mail(from); err != nil {
		return err
	}

	contentType := "Content-Type: text/plain;charset=UTF-8"
	msgStr := fmt.Sprint(
		"To:", strings.Join(to, ";"),
		"\r\nFrom:", from,
		"\r\nSubject:", subject,
		"\r\n", contentType,
		"\r\n\r\n", body,
	)
	msg := []byte(msgStr)
	for _, addr := range to {
		if err := client.Rcpt(addr); err != nil {
			return err
		}
	}
	writer, err := client.Data()
	if err != nil {
		return err
	}
	_, err = writer.Write(msg)
	if err != nil {
		return err
	}
	writer.Close()
	return nil
}
