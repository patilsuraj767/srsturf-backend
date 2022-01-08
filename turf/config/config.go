package config

import (
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "15432")
	viper.SetDefault("DB_NAME", "postgres")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "postgres")
}
