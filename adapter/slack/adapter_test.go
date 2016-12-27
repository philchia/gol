package slack

import "testing"

func TestNewAdapter(t *testing.T) {
	type args struct {
		webhook string
	}
	tests := []struct {
		name    string
		args    args
		wantNil bool
	}{
		{
			"case1",
			args{
				"http://slack.com",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAdapter(tt.args.webhook); (got == nil) != tt.wantNil {
				t.Error("NewAdapter() got != want")
			}
		})
	}
}

func Test_slackWriter_Write(t *testing.T) {
	type fields struct {
		webhookURL string
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &slackWriter{
				webhookURL: tt.fields.webhookURL,
			}
			got, err := s.Write(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("slackWriter.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("slackWriter.Write() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_slackWriter_Close(t *testing.T) {
	type fields struct {
		webhookURL string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"case1",
			fields{
				"http://slack.com",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &slackWriter{
				webhookURL: tt.fields.webhookURL,
			}
			if err := s.Close(); (err != nil) != tt.wantErr {
				t.Errorf("slackWriter.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
