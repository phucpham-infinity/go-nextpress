package configs

import (
	"fmt"

	"github.com/phucpham-infinity/go-nextpress/app/global"
)

func Run() {

	InitEnv()
	fmt.Println("Init env", global.Env.MySql.Host)

}
