package cmd

import (
	"testing"
)

func TestRun(t *testing.T) {
	go Run("test", func() {})
}
