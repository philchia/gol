package slack

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/philchia/gol/adapter"
	"github.com/philchia/gol/internal"
	"github.com/philchia/gol/level"
)

var _ adapter.Adapter = (*slackWriter)(nil)

type slackWriter struct {
	webhookURL string
	logLevel   level.LogLevel
}

// NewAdapter create a slack adapter
func NewAdapter(webhook string, l ...level.LogLevel) adapter.Adapter {
	s := &slackWriter{
		webhookURL: webhook,
	}
	if len(l) > 0 {
		s.logLevel = l[0]
	}
	return s
}

func (s *slackWriter) Write(b []byte) (int, error) {
	msg := fmt.Sprintf("{\"test\":\"%s\"}", internal.Bytes2str(b))
	form := url.Values{}
	form.Add("payload", msg)

	resp, err := http.PostForm(s.webhookURL, form)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Post webhook failed %s %d", resp.Status, resp.StatusCode)
	}
	resp.Body.Close()
	return len(b), nil
}

func (s *slackWriter) Close() error {
	return nil
}

func (s *slackWriter) Level() level.LogLevel {
	return s.logLevel
}
