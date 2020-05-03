package conf

import (
	"github.com/spf13/viper"
)

// TODO unit test

func Load(v interface{}) {
	viper.SetConfigName("ncs")
	viper.AddConfigPath("/ncs/")
	err := viper.Unmarshal(v)
	if err != nil {
		panic(err)
	}
}
