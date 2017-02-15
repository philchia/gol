package gol

import (
	"testing"

	"github.com/philchia/gol/adapter"
	"github.com/philchia/gol/adapter/fake"
	"github.com/philchia/gol/adapter/fakeSync"
	"github.com/philchia/gol/level"

	"log"
	"reflect"
)

var _adapter = fakeSync.NewAdapter()

func init() {
	logger.AddLogAdapter("fake", _adapter)
	logger.RemoveAdapter(CONSOLELOGGER)
	logger.SetOption(LogOption(0))
}

func TestDebug(t *testing.T) {
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
			Debug(tt.args.i...)
			if got := _adapter.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Debug() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func TestDebugf(t *testing.T) {
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
			Debugf(tt.args.format, tt.args.i...)
			if got := _adapter.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Debugf() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func TestInfo(t *testing.T) {
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
			Info(tt.args.i...)
			if got := _adapter.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Info() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func TestInfof(t *testing.T) {
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
			Infof(tt.args.format, tt.args.i...)
			if got := _adapter.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Infof() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func TestWarn(t *testing.T) {
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
			Warn(tt.args.i...)
			if got := _adapter.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Warn() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func TestWarnf(t *testing.T) {
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
			Warnf(tt.args.format, tt.args.i...)
			if got := _adapter.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Warnf() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func TestError(t *testing.T) {
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
			Error(tt.args.i...)
			if got := _adapter.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func TestErrorf(t *testing.T) {
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
			Errorf(tt.args.format, tt.args.i...)
			if got := _adapter.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Errorf() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func TestCritical(t *testing.T) {
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
			Critical(tt.args.i...)
			if got := _adapter.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Critical() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func TestCriticalf(t *testing.T) {
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
			Criticalf(tt.args.format, tt.args.i...)
			if got := _adapter.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Criticalf() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func TestSetLevel(t *testing.T) {
	type args struct {
		level level.LogLevel
	}
	tests := []struct {
		name string
		args args
		want level.LogLevel
	}{
		{
			"case1",
			args{
				level.DEBUG,
			},
			level.DEBUG,
		},
		{
			"case2",
			args{
				level.WARN,
			},
			level.WARN,
		},
		{
			"case2",
			args{
				level.ERROR,
			},
			level.ERROR,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLevel(tt.args.level)
			if got := logger.(*gollog).level; got != tt.want {
				t.Errorf("SetLevel(level) = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestSetOption(t *testing.T) {
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
			SetOption(tt.args.option)
			if got := logger.(*gollog).option; got != tt.want {
				t.Errorf("SetOption(option) = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestAddLogAdapter(t *testing.T) {
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
			if err := AddLogAdapter(tt.args.name, tt.args.adapter); (err != nil) != tt.wantErr {
				t.Errorf("AddLogAdapter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRemoveAdapter(t *testing.T) {
	defer AddLogAdapter("fake", _adapter)

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
			if err := RemoveAdapter(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("RemoveAdapter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFlush(t *testing.T) {
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
			Flush()
		})
	}
}

func BenchmarkLog(b *testing.B) {
	l := log.New(fake.NewAdapter(), "\033[32m[DEBUG]\033[0m ", log.LstdFlags)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l.Println("Hello")
	}
}

func BenchmarkGol(b *testing.B) {
	g := NewLogger(level.DEBUG)
	g.RemoveAdapter(CONSOLELOGGER)
	// g.SetOption(0)
	g.AddLogAdapter("fake", fake.NewAdapter(level.DEBUG))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Debug("Hello")
	}
}

func BenchmarkMultiThreadLog(b *testing.B) {
	l := log.New(fakeSync.NewAdapter(), "\033[32m[DEBUG]\033[0m ", log.LstdFlags)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			go l.Println("Hello")
		}
	}
}

func BenchmarkMultiThreadGol(b *testing.B) {
	g := NewLogger(level.DEBUG)
	g.RemoveAdapter(CONSOLELOGGER)
	g.AddLogAdapter("fake", fakeSync.NewAdapter(level.DEBUG))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			go g.Debug("Hello")
		}
	}
}
