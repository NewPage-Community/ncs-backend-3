package log

import "testing"

func TestInit(t *testing.T) {
	type args struct {
		debug bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"pro",
			args{false},
		},
		{
			"debug",
			args{true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init(tt.args.debug)
			defer Close()
			Info("This is a info message")
			Warn("This is a warn message")
			Error("This a error message")
			Debug("This a debug message")
		})
	}
}