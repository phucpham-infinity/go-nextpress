package logger

import "github.com/spf13/viper"

func init() {
	viper := viper.New()
	viper.AddConfigPath("")
}
