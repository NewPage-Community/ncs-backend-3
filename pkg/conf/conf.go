package conf

import (
	"backend/pkg/log"
	"github.com/spf13/viper"
)

func Load(v interface{}) {
	viper.SetConfigName("ncs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/ncs/")
	viper.AddConfigPath("/Users/gunslinger/ncs-test/")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err = viper.SafeWriteConfig(); err != nil {
				log.Error(err)
			}
		} else {
			panic(err)
		}
	}
	if err := viper.Unmarshal(v); err != nil {
		panic(err)
	}
}
