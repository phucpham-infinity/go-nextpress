package global

type ENV struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	MySql struct {
		Host            string `mapstructure:"host"`
		Port            int    `mapstructure:"port"`
		Username        string `mapstructure:"username"`
		Password        string `mapstructure:"password"`
		DbName          string `mapstructure:"dbName"`
		MaxIdeConns     int    `mapstructure:"maxIdeConns"`
		MaxOpenConns    int    `mapstructure:"maxOpenConns"`
		ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
	} `mapstructure:"mysql"`
}

var (
	Env ENV
)
