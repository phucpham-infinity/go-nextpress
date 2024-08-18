package configs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/phucpham-infinity/go-nextpress/app/global"
	"github.com/spf13/viper"
)

func InitEnv() {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if strings.Contains(pwd, "/cmd/server") {
		pwd = filepath.Dir(pwd) // lùi lại 1 cấp
		pwd = filepath.Dir(pwd) // lùi lại 1 cấp nữa
	}

	fmt.Println(pwd)

	viper := viper.New()
	viper.AddConfigPath(pwd + "/env")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&global.Env); err != nil {
		fmt.Printf("Unable to decode into struct, %v\n", err)
	}
}
