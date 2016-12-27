package gol

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/philchia/gol/adapter"
	"github.com/philchia/gol/adapter/fakeSync"
)

var _logger = NewLogger(DEBUG)
var _loggerAdapter = fakeSync.NewAdapter()

func init() {

	_logger.AddLogAdapter("fake", _loggerAdapter)
	_logger.RemoveAdapter(CONSOLELOGGER)
	_logger.SetOption(LogOption(0))
}

func Test_gollog_Debug(t *testing.T) {
	type args struct {
		i []interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"case1",
			args{
				[]interface{}{"Hello world"},
			},
			[]byte("\033[32m[DEBUG]\033[0m Hello world\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_logger.Debug(tt.args.i...)
			if got := _loggerAdapter.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Debug() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func Test_gollog_Debugf(t *testing.T) {
	type args struct {
		format string
		i      []interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"case1",
			args{
				"%s",
				[]interface{}{"Hello world"},
			},
			[]byte("\033[32m[DEBUG]\033[0m Hello world\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_logger.Debugf(tt.args.format, tt.args.i...)
			if got := _loggerAdapter.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Debugf() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func Test_gollog_Info(t *testing.T) {
	type args struct {
		i []interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"case1",
			args{
				[]interface{}{"Hello world"},
			},
			[]byte("\033[34m[INFO]\033[0m Hello world\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_logger.Info(tt.args.i...)
			if got := _loggerAdapter.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Info() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func Test_gollog_Infof(t *testing.T) {
	type args struct {
		format string
		i      []interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"case1",
			args{
				"%s",
				[]interface{}{"Hello world"},
			},
			[]byte("\033[34m[INFO]\033[0m Hello world\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_logger.Infof(tt.args.format, tt.args.i...)
			if got := _loggerAdapter.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Infof() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func Test_gollog_Warn(t *testing.T) {
	type args struct {
		i []interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"case1",
			args{
				[]interface{}{"Hello world"},
			},
			[]byte("\033[33m[WARN]\033[0m Hello world\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_logger.Warn(tt.args.i...)
			if got := _loggerAdapter.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Warn() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func Test_gollog_Warnf(t *testing.T) {
	type args struct {
		format string
		i      []interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"case1",
			args{
				"%s",
				[]interface{}{"Hello world"},
			},
			[]byte("\033[33m[WARN]\033[0m Hello world\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_logger.Warnf(tt.args.format, tt.args.i...)
			if got := _loggerAdapter.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Warnf() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func Test_gollog_Error(t *testing.T) {
	type args struct {
		i []interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"case1",
			args{
				[]interface{}{"Hello world"},
			},
			[]byte("\033[31m[ERROR]\033[0m Hello world\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_logger.Error(tt.args.i...)
			if got := _loggerAdapter.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func Test_gollog_Errorf(t *testing.T) {
	type args struct {
		format string
		i      []interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"case1",
			args{
				"%s",
				[]interface{}{"Hello world"},
			},
			[]byte("\033[31m[ERROR]\033[0m Hello world\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_logger.Errorf(tt.args.format, tt.args.i...)
			if got := _loggerAdapter.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Errorf() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func Test_gollog_Critical(t *testing.T) {
	type args struct {
		i []interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"case1",
			args{
				[]interface{}{"Hello world"},
			},
			[]byte("\033[35m[CRITICAL]\033[0m Hello world\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_logger.Critical(tt.args.i...)
			if got := _loggerAdapter.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Critical() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func Test_gollog_Criticalf(t *testing.T) {
	type args struct {
		format string
		i      []interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"case1",
			args{
				"%s",
				[]interface{}{"Hello world"},
			},
			[]byte("\033[35m[CRITICAL]\033[0m Hello world\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_logger.Criticalf(tt.args.format, tt.args.i...)
			if got := _loggerAdapter.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Criticalf() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func Test_gollog_SetLevel(t *testing.T) {
	type args struct {
		level LogLevel
	}
	tests := []struct {
		name string
		args args
		want LogLevel
	}{
		{
			"case1",
			args{
				DEBUG,
			},
			DEBUG,
		},
		{
			"case2",
			args{
				WARN,
			},
			WARN,
		},
		{
			"case2",
			args{
				ERROR,
			},
			ERROR,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_logger.SetLevel(tt.args.level)
			if got := _logger.(*gollog).level; got != tt.want {
				t.Errorf("SetLevel(level) = %d, want %d", got, tt.want)
			}
		})
	}
}

func Test_gollog_SetOption(t *testing.T) {
	type args struct {
		option LogOption
	}
	tests := []struct {
		name string
		args args
		want LogOption
	}{
		{
			"case1",
			args{
				Lmicroseconds,
			},
			Lmicroseconds,
		},
		{
			"case2",
			args{
				Ltime,
			},
			Ltime,
		},
		{
			"case3",
			args{
				Ltime | Ldate,
			},
			Ltime | Ldate,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_logger.SetOption(tt.args.option)
			if got := _logger.(*gollog).option; got != tt.want {
				t.Errorf("SetOption(option) = %d, want %d", got, tt.want)
			}
		})
	}
}

func Test_gollog_AddLogAdapter(t *testing.T) {
	type args struct {
		name    string
		adapter adapter.Adapter
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"case1",
			args{
				"fake1",
				fakeSync.NewAdapter(),
			},
			false,
		},
		{
			"case2",
			args{
				"fake",
				fakeSync.NewAdapter(),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := _logger.AddLogAdapter(tt.args.name, tt.args.adapter); (err != nil) != tt.wantErr {
				t.Errorf("AddLogAdapter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gollog_RemoveAdapter(t *testing.T) {
	defer AddLogAdapter("fake", _loggerAdapter)

	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"case1",
			args{
				"fake2",
			},
			true,
		},
		{
			"case2",
			args{
				"fake",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := _logger.RemoveAdapter(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("RemoveAdapter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gollog_Flush(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			"case1",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_logger.Flush()
		})
	}
}

func Test_itoa(t *testing.T) {
	type args struct {
		buf *bytes.Buffer
		i   int
		wid int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"case0",
			args{
				bytes.NewBuffer(nil),
				5,
				1,
			},
		},
		{
			"case1",
			args{
				bytes.NewBuffer(nil),
				15,
				2,
			},
		},
		{
			"case2",
			args{
				bytes.NewBuffer(nil),
				155,
				3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			itoa(tt.args.buf, tt.args.i, tt.args.wid)
			if got := tt.args.buf.String(); got != fmt.Sprintf("%d", tt.args.i) {
				t.Errorf("itoa() got %s, want %s", got, fmt.Sprintf("%d", tt.args.i))
			}
		})
	}
}

func Test_gollog_generatePrefix(t *testing.T) {
	type fields struct {
		level    LogLevel
		option   LogOption
		adapters map[string]adapter.Adapter
		logChan  chan *bytes.Buffer
		doneChan chan struct{}
	}
	type args struct {
		buf       *bytes.Buffer
		callDepth int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"case1",
			fields{
				DEBUG,
				LstdFlags | Lmicroseconds | LUTC,
				map[string]adapter.Adapter{},
				make(chan *bytes.Buffer, 10),
				make(chan struct{}),
			},
			args{
				bytes.NewBuffer(nil),
				2,
			},
		},
		{
			"case2",
			fields{
				WARN,
				LstdFlags | Lmicroseconds | LUTC,
				map[string]adapter.Adapter{},
				make(chan *bytes.Buffer, 10),
				make(chan struct{}),
			},
			args{
				bytes.NewBuffer(nil),
				20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gollog{
				level:    tt.fields.level,
				option:   tt.fields.option,
				adapters: tt.fields.adapters,
				logChan:  tt.fields.logChan,
				doneChan: tt.fields.doneChan,
			}
			l.generatePrefix(tt.args.buf, tt.args.callDepth)
			if got := tt.args.buf.String(); got == "" {
				t.Error(`generatePrefix() got = "",want not nil`)
			}
		})
	}
}
