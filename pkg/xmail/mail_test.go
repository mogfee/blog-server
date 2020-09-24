package xmail

import (
	"fmt"
	"github.com/xinliangnote/go-util/mail"
	"testing"
)

func TestSendmail(t *testing.T) {
	options := &mail.Options{
		MailHost: "smtp.126.com",
		MailPort: 465,
		MailUser: "mogfee@126.com",
		MailPass: "JMIWMUZXDXRUJKRO",
		MailTo:   "mogfee@126.com",
		Subject:  "subject",
		Body:     "body",
	}
	fmt.Println(mail.Send(options))
}
