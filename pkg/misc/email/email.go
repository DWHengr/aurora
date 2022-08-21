package email

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

type EmailConfig struct {
	Host     string `yaml:"host""`
	Port     string `yaml:"port""`
	Username string `yaml:"username""`
	Password string `yaml:"password""`
	Thread   int    `yaml:"thread""`
}

var emailSender *email.Pool

func NewEmail(c *EmailConfig) {
	var err error
	emailSender, err = email.NewPool(
		c.Host+":"+c.Port,
		c.Thread,
		smtp.PlainAuth("", c.Username, c.Password, c.Host),
	)
	if err != nil {
		panic(err)
	}
}

func GetEmailSender() *email.Pool {
	return emailSender
}
