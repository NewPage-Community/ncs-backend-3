package log

import (
	"testing"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name string
		args Config
	}{
		{
			"pro",
			Config{Debug: false},
		},
		{
			"debug",
			Config{Debug: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init(&tt.args)
			defer func() {
				Close()
				Warn(recover())
			}()

			Info("This is a info message")
			Warn("This is a warn message")
			Error("This a error message")
			Debug("This a debug message")
			Panic("This a panic message")
		})
	}
}
