package configs

import "github.com/spf13/viper"

type conf struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBName            string `mapstructure:"DB_NAME"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            int `mapstructure:"DB_PORT"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	WebServerPort     int `mapstructure:"WEB_SERVER_PORT"`
	IPMaxRequests	int `mapstructure:"ID_MAX_REQUEST"`
	TokenMaxRequests	int `mapstructure:"TOKEN_MAX_REQUEST"`
	OperatingWindowMs int `mapstructure:"TIME_WINDOW_MS"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}