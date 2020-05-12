package conf

import (
	"github.com/spf13/viper"
)

func Load(v interface{}) {
	viper.SetConfigName("ncs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/ncs/")
	viper.AddConfigPath("/Users/gunslinger/ncs-test/")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.SafeWriteConfig()
		} else {
			panic(err)
		}
	}
	if err := viper.Unmarshal(v); err != nil {
		panic(err)
	}
}
