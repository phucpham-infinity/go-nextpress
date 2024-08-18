package structs

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
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
	Logger LoggerConfig `mapstructure:"logger"`
	Redis  RedisConfig  `mapstructure:"redis"`
}

type LoggerConfig struct {
	Level    string `mapstructure:"level"`
	Filename string `mapstructure:"filename"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
