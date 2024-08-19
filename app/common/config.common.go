package common

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	MySql  MySqlConfig  `mapstructure:"mysql"`
	Logger LoggerConfig `mapstructure:"logger"`
	Redis  RedisConfig  `mapstructure:"redis"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type MySqlConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	DbName          string `mapstructure:"dbName"`
	MaxIdeConns     int    `mapstructure:"maxIdeConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type LoggerConfig struct {
	Level    string `mapstructure:"level"`
	Filename string `mapstructure:"filename"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
