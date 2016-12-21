package internal

import (
	"reflect"
	"testing"
)

func TestJoinStrings(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"case1",
			args{
				[]string{"Hello", " world", "!!!!"},
			},
			"Hello world!!!!",
		},
		{
			"case2",
			args{
				[]string{"Hello", " from", " gol"},
			},
			"Hello from gol",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinStrings(tt.args.strs...); got != tt.want {
				t.Errorf("JoinStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStr2bytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"case1",
			args{
				"Hello",
			},
			[]byte("Hello"),
		},
		{
			"case2",
			args{
				"Hello world",
			},
			[]byte("Hello world"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Str2bytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Str2bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes2str(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"case1",
			args{
				[]byte("Hello"),
			},
			"Hello",
		},
		{
			"case2",
			args{
				[]byte("Hello world"),
			},
			"Hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bytes2str(tt.args.b); got != tt.want {
				t.Errorf("Bytes2str() = %v, want %v", got, tt.want)
			}
		})
	}
}
