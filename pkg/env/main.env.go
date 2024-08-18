package env

import "github.com/spf13/viper"

type Env struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
}

func ENV() Env {
	viper := viper.New()
	viper.AddConfigPath("./env")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	var env Env
	if err := viper.Unmarshal(&env); err != nil {
		return env
	}
	return env
}
