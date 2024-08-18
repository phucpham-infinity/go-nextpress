package configs

import (
	"fmt"

	"github.com/phucpham-infinity/go-nextpress/app/global"
	"github.com/spf13/viper"
)

func InitEnv() {
	viper := viper.New()
	viper.AddConfigPath("./env")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	if err := viper.Unmarshal(&global.Env); err != nil {
		fmt.Printf("Unable to decode into struct, %v\n", err)
	}
}
