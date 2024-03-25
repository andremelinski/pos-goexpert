package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

// conf nao esta exportado, entao vc nao pode importar, precisa fazer um load para exportar o struct
type conf struct {
	DBDrive string `mapstructure:"DBDrive"` 
	DBHost string `mapstructure:"DBHost"`
	DBPort string `mapstructure:"DBPort"`
	DBUser string `mapstructure:"DBUser"`
	DBPassword string `mapstructure:"DBPassword"`
	DBName string `mapstructure:"DBName"`
	WebServerPort string `mapstructure:"WebServerPort"`
	JwTSecret string `mapstructure:"JwTSecret"`
	JwtExpiresIn int `mapstructure:"JwtExpiresIn"`
	TokenAuth *jwtauth.JWTAuth
}

// passa o path de um arquivo que tem as configs 
func LoadConfig(path string) (*conf, error){
	cfg := conf{}

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	// se o .env nao existe ele pega das variaveis de ambiente
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err) 
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err) 
	}
	// inicia o struct jwtauth por onde vai ser realizado o encode e decode  do jwt
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JwTSecret), nil)
	return &cfg, err
}