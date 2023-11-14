package bootstrap

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	ContextTimeout  int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost          string `mapstructure:"DB_HOST_MONGO"`
	DBPort          string `mapstructure:"DB_PORT_MONGO"`
	DBUser          string `mapstructure:"DB_USER_MONGO"`
	DBPass          string `mapstructure:"DB_PASS_MONGO"`
	DBName          string `mapstructure:"DB_NAME_MONGO"`
	NumberOfBlocks  int32  `mapstructure:"NUMBER_OF_BLOCKS"`
	NumberOfWallets int32  `mapstructure:"NUMBER_OF_WALLETS"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}
