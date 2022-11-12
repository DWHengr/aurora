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
var emailConfig *EmailConfig

func NewEmail(c *EmailConfig) {
	emailConfig = c
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

func GetUsername() string {
	if emailConfig != nil {
		return emailConfig.Username
	}
	return ""
}

func GetEmailSender() *email.Pool {
	return emailSender
}
