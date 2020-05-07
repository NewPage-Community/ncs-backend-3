package conf

import (
	"github.com/spf13/viper"
)

func Load(v interface{}) {
	viper.SetConfigName("ncs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/ncs/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.SafeWriteConfig()
		} else {
			panic(err)
		}
	}
	err = viper.Unmarshal(v)
	if err != nil {
		panic(err)
	}
}
