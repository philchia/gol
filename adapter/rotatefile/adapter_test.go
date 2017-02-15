package rotatefile

import "testing"

func TestNewAdapter(t *testing.T) {
	type args struct {
		name            string
		maxFileNum      int
		maxBytesPerFile ByteSize
	}
	tests := []struct {
		name    string
		args    args
		wantNil bool
	}{
		{
			"case1",
			args{
				`../../assets/test.log`,
				2,
				MB,
			},
			false,
		},
		{
			"case2",
			args{
				`../../assets/test.log`,
				4,
				MB,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAdapter(tt.args.name, tt.args.maxFileNum, tt.args.maxBytesPerFile); (got == nil) != tt.wantNil {
				t.Error("NewAdapter(), got != want")
			}
		})
	}
}
