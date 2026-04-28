package notice

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strings"
)

type EmailSet struct {
	SMTPHost string   `yaml:"smtpHost"`
	SMTPPort int      `yaml:"smtpPort"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
	From     string   `yaml:"from"`
	To       []string `yaml:"to"`
	Subject  string   `yaml:"subject"`
}

func EmailPushMessage(emailSet EmailSet, message string) error {
	if len(emailSet.To) == 0 {
		return nil
	}
	if emailSet.SMTPPort == 0 {
		emailSet.SMTPPort = 587
	}
	if emailSet.From == "" {
		emailSet.From = emailSet.Username
	}
	if emailSet.Subject == "" {
		emailSet.Subject = "山姆库存提醒"
	}
	if message == "" {
		message = emailSet.Subject
	}

	addr := fmt.Sprintf("%s:%d", emailSet.SMTPHost, emailSet.SMTPPort)
	auth := smtp.PlainAuth("", emailSet.Username, emailSet.Password, emailSet.SMTPHost)
	content := buildEmailContent(emailSet, message)
	return smtp.SendMail(addr, auth, emailSet.From, emailSet.To, content)
}

func buildEmailContent(emailSet EmailSet, message string) []byte {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("From: %s\r\n", emailSet.From))
	buf.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(emailSet.To, ",")))
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", emailSet.Subject))
	buf.WriteString("MIME-Version: 1.0\r\n")
	buf.WriteString("Content-Type: text/plain; charset=UTF-8\r\n")
	buf.WriteString("\r\n")
	buf.WriteString(message)
	return buf.Bytes()
}
