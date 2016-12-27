package mail

import "testing"

func Test_mailAdapter_Write(t *testing.T) {
	type fields struct {
		host      string
		account   string
		password  string
		subject   string
		receivers []string
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
			m := &mailAdapter{
				host:      tt.fields.host,
				account:   tt.fields.account,
				password:  tt.fields.password,
				subject:   tt.fields.subject,
				receivers: tt.fields.receivers,
			}
			got, err := m.Write(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("mailAdapter.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mailAdapter.Write() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mailAdapter_Close(t *testing.T) {
	type fields struct {
		host      string
		account   string
		password  string
		subject   string
		receivers []string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{{
		"case1",
		fields{
			"smtp.gmail.com:25",
			"test@gmail.com",
			"password",
			"test subject",
			[]string{"test@hotmail.com"},
		},
		false,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mailAdapter{
				host:      tt.fields.host,
				account:   tt.fields.account,
				password:  tt.fields.password,
				subject:   tt.fields.subject,
				receivers: tt.fields.receivers,
			}
			if err := m.Close(); (err != nil) != tt.wantErr {
				t.Errorf("mailAdapter.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewAdapter(t *testing.T) {
	type args struct {
		host      string
		account   string
		password  string
		subject   string
		receivers []string
	}
	tests := []struct {
		name    string
		args    args
		wantNil bool
	}{
		{
			"case1",
			args{
				"smtp.gmail.com:25",
				"test@gmail.com",
				"password",
				"test subject",
				[]string{"test@hotmail.com"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAdapter(tt.args.host, tt.args.account, tt.args.password, tt.args.subject, tt.args.receivers...); (got == nil) != tt.wantNil {
				t.Error("NewAdapter() got != want")
			}
		})
	}
}
