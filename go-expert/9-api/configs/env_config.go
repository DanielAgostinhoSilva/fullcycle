package configs

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/spf13/viper"
	"log"
)

type EnvConfig struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	DBDsn         string `mapstructure:"DB_DSN"`
	MigrationDir  string `mapstructure:"MIGRATION_DIR"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JwtSecret     string `mapstructure:"JWT_SECRET"`
	JwExpiresIn   int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig(filePath string) *EnvConfig {
	var cfg EnvConfig
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.SetConfigFile(filePath)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	log.Println(".env file loaded")
	return &cfg
}
