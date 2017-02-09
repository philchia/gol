package mail

import (
	"net/smtp"
	"strings"

	"fmt"

	"github.com/philchia/gol/adapter"
	"github.com/philchia/gol/level"
)

var _ adapter.Adapter = (*mailAdapter)(nil)

type mailAdapter struct {
	host      string
	account   string
	password  string
	subject   string
	receivers []string
	logLevel  level.LogLevel
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
func NewAdapter(host, account, password, subject string, receivers []string, l ...level.LogLevel) adapter.Adapter {
	adapter := &mailAdapter{
		host:      host,
		account:   account,
		password:  password,
		subject:   subject,
		receivers: receivers,
	}
	if len(l) > 0 {
		adapter.logLevel = l[0]
	}
	return adapter
}

func (m *mailAdapter) Level() level.LogLevel {
	return m.logLevel
}
