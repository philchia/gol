package gol

import (
	"bytes"
	"testing"

	"github.com/philchia/gol/adapter"
)

func Test_gollog_msgPump(t *testing.T) {
	type fields struct {
		level    LogLevel
		option   LogOption
		adapters map[string]adapter.Adapter
		logChan  chan string
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gollog{
				level:    tt.fields.level,
				option:   tt.fields.option,
				adapters: tt.fields.adapters,
				logChan:  tt.fields.logChan,
			}
			l.msgPump()
		})
	}
}

func Test_gollog_put(t *testing.T) {
	type fields struct {
		level    LogLevel
		option   LogOption
		adapters map[string]adapter.Adapter
		logChan  chan string
	}
	type args struct {
		msg string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gollog{
				level:    tt.fields.level,
				option:   tt.fields.option,
				adapters: tt.fields.adapters,
				logChan:  tt.fields.logChan,
			}
			l.put(tt.args.msg)
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
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			itoa(tt.args.buf, tt.args.i, tt.args.wid)
		})
	}
}

func Test_gollog_generatePrefix(t *testing.T) {
	type fields struct {
		level    LogLevel
		option   LogOption
		adapters map[string]adapter.Adapter
		logChan  chan string
	}
	type args struct {
		callDepth int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gollog{
				level:    tt.fields.level,
				option:   tt.fields.option,
				adapters: tt.fields.adapters,
				logChan:  tt.fields.logChan,
			}
			if got := l.generatePrefix(tt.args.callDepth); got != tt.want {
				t.Errorf("gollog.generatePrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gollog_generateLog(t *testing.T) {
	type fields struct {
		level    LogLevel
		option   LogOption
		adapters map[string]adapter.Adapter
		logChan  chan string
	}
	type args struct {
		callDepth int
		level     LogLevel
		msg       string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gollog{
				level:    tt.fields.level,
				option:   tt.fields.option,
				adapters: tt.fields.adapters,
				logChan:  tt.fields.logChan,
			}
			if got := l.generateLog(tt.args.callDepth, tt.args.level, tt.args.msg); got != tt.want {
				t.Errorf("gollog.generateLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gollog_Debug(t *testing.T) {
	type fields struct {
		level    LogLevel
		option   LogOption
		adapters map[string]adapter.Adapter
		logChan  chan string
	}
	type args struct {
		i []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gollog{
				level:    tt.fields.level,
				option:   tt.fields.option,
				adapters: tt.fields.adapters,
				logChan:  tt.fields.logChan,
			}
			l.Debug(tt.args.i...)
		})
	}
}

func Test_gollog_Debugf(t *testing.T) {
	type fields struct {
		level    LogLevel
		option   LogOption
		adapters map[string]adapter.Adapter
		logChan  chan string
	}
	type args struct {
		format string
		i      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gollog{
				level:    tt.fields.level,
				option:   tt.fields.option,
				adapters: tt.fields.adapters,
				logChan:  tt.fields.logChan,
			}
			l.Debugf(tt.args.format, tt.args.i...)
		})
	}
}

func Test_gollog_Info(t *testing.T) {
	type fields struct {
		level    LogLevel
		option   LogOption
		adapters map[string]adapter.Adapter
		logChan  chan string
	}
	type args struct {
		i []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gollog{
				level:    tt.fields.level,
				option:   tt.fields.option,
				adapters: tt.fields.adapters,
				logChan:  tt.fields.logChan,
			}
			l.Info(tt.args.i...)
		})
	}
}

func Test_gollog_Infof(t *testing.T) {
	type fields struct {
		level    LogLevel
		option   LogOption
		adapters map[string]adapter.Adapter
		logChan  chan string
	}
	type args struct {
		format string
		i      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gollog{
				level:    tt.fields.level,
				option:   tt.fields.option,
				adapters: tt.fields.adapters,
				logChan:  tt.fields.logChan,
			}
			l.Infof(tt.args.format, tt.args.i...)
		})
	}
}

func Test_gollog_Warn(t *testing.T) {
	type fields struct {
		level    LogLevel
		option   LogOption
		adapters map[string]adapter.Adapter
		logChan  chan string
	}
	type args struct {
		i []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gollog{
				level:    tt.fields.level,
				option:   tt.fields.option,
				adapters: tt.fields.adapters,
				logChan:  tt.fields.logChan,
			}
			l.Warn(tt.args.i...)
		})
	}
}

func Test_gollog_Warnf(t *testing.T) {
	type fields struct {
		level    LogLevel
		option   LogOption
		adapters map[string]adapter.Adapter
		logChan  chan string
	}
	type args struct {
		format string
		i      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gollog{
				level:    tt.fields.level,
				option:   tt.fields.option,
				adapters: tt.fields.adapters,
				logChan:  tt.fields.logChan,
			}
			l.Warnf(tt.args.format, tt.args.i...)
		})
	}
}

func Test_gollog_Error(t *testing.T) {
	type fields struct {
		level    LogLevel
		option   LogOption
		adapters map[string]adapter.Adapter
		logChan  chan string
	}
	type args struct {
		i []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gollog{
				level:    tt.fields.level,
				option:   tt.fields.option,
				adapters: tt.fields.adapters,
				logChan:  tt.fields.logChan,
			}
			l.Error(tt.args.i...)
		})
	}
}

func Test_gollog_Errorf(t *testing.T) {
	type fields struct {
		level    LogLevel
		option   LogOption
		adapters map[string]adapter.Adapter
		logChan  chan string
	}
	type args struct {
		format string
		i      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gollog{
				level:    tt.fields.level,
				option:   tt.fields.option,
				adapters: tt.fields.adapters,
				logChan:  tt.fields.logChan,
			}
			l.Errorf(tt.args.format, tt.args.i...)
		})
	}
}

func Test_gollog_Critical(t *testing.T) {
	type fields struct {
		level    LogLevel
		option   LogOption
		adapters map[string]adapter.Adapter
		logChan  chan string
	}
	type args struct {
		i []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gollog{
				level:    tt.fields.level,
				option:   tt.fields.option,
				adapters: tt.fields.adapters,
				logChan:  tt.fields.logChan,
			}
			l.Critical(tt.args.i...)
		})
	}
}

func Test_gollog_Criticalf(t *testing.T) {
	type fields struct {
		level    LogLevel
		option   LogOption
		adapters map[string]adapter.Adapter
		logChan  chan string
	}
	type args struct {
		format string
		i      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gollog{
				level:    tt.fields.level,
				option:   tt.fields.option,
				adapters: tt.fields.adapters,
				logChan:  tt.fields.logChan,
			}
			l.Criticalf(tt.args.format, tt.args.i...)
		})
	}
}

func Test_gollog_SetLevel(t *testing.T) {
	type fields struct {
		level    LogLevel
		option   LogOption
		adapters map[string]adapter.Adapter
		logChan  chan string
	}
	type args struct {
		level LogLevel
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gollog{
				level:    tt.fields.level,
				option:   tt.fields.option,
				adapters: tt.fields.adapters,
				logChan:  tt.fields.logChan,
			}
			l.SetLevel(tt.args.level)
		})
	}
}

func Test_gollog_SetOption(t *testing.T) {
	type fields struct {
		level    LogLevel
		option   LogOption
		adapters map[string]adapter.Adapter
		logChan  chan string
	}
	type args struct {
		option LogOption
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gollog{
				level:    tt.fields.level,
				option:   tt.fields.option,
				adapters: tt.fields.adapters,
				logChan:  tt.fields.logChan,
			}
			l.SetOption(tt.args.option)
		})
	}
}

func Test_gollog_AddLogAdapter(t *testing.T) {
	type fields struct {
		level    LogLevel
		option   LogOption
		adapters map[string]adapter.Adapter
		logChan  chan string
	}
	type args struct {
		a    adapter.Adapter
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gollog{
				level:    tt.fields.level,
				option:   tt.fields.option,
				adapters: tt.fields.adapters,
				logChan:  tt.fields.logChan,
			}
			l.AddLogAdapter(tt.args.name, tt.args.a)
		})
	}
}
