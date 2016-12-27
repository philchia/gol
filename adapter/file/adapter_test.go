package file

import "testing"

func TestNewAdapter(t *testing.T) {
	type args struct {
		pathToFile string
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
			},
			false,
		},
		{
			"case3",
			args{
				`$#@@^&*&\|||\\\\///assets/test.log`,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAdapter(tt.args.pathToFile); (got == nil) != tt.wantNil {
				t.Errorf("NewFileAdapter() , got != want")
			}
		})
	}
}
