package conf

import (
	"testing"
)

type Config struct {
	Test struct {
		Port   int
		Enable bool
	}
}

func TestLoad(t *testing.T) {
	c := &Config{}
	Load(c)
	t.Log(c)
	if c.Test.Port != 2 && !c.Test.Enable {
		t.Errorf("Load(): Test config v[%v, %v], want[2, true]", c.Test.Port, c.Test.Enable)
	}
}
