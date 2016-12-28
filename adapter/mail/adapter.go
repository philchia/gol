package mail

import (
	"net/smtp"
	"strings"

	"fmt"

	"github.com/philchia/gol/adapter"
)

var _ adapter.Adapter = (*mailAdapter)(nil)

type mailAdapter struct {
	host      string
	account   string
	password  string
	subject   string
	receivers []string
}

func (m *mailAdapter) Write(b []byte) (int, error) {

	hp := strings.Split(m.host, ":")
	auth := smtp.PlainAuth("", m.account, m.password, hp[0])
	err := smtp.SendMail(m.host, auth, m.account, m.receivers, b)
	if err != nil {
		fmt.Println(err)
	}
	return len(b), nil
}

func (m *mailAdapter) Close() error {
	return nil
}

// NewAdapter create a mail adapter
func NewAdapter(host, account, password, subject string, receivers ...string) adapter.Adapter {
	adapter := &mailAdapter{
		host:      host,
		account:   account,
		password:  password,
		subject:   subject,
		receivers: receivers,
	}
	return adapter
}
