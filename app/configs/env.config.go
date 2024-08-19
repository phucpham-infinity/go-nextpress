package configs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/phucpham-infinity/go-nextpress/app/common"
	"github.com/phucpham-infinity/go-nextpress/app/context"
	"github.com/spf13/viper"
)

func InitEnv() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if strings.Contains(pwd, "/cmd/server") {
		pwd = filepath.Dir(pwd)
		pwd = filepath.Dir(pwd)
	}
	viper := viper.New()
	viper.AddConfigPath(pwd + "/env")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file , %s", err)
		os.Exit(1)
	}

	var config common.Config

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode into struct, %v\n", err)
		os.Exit(1)
	}
	context.AppContext().SetConfig(&config)
}
