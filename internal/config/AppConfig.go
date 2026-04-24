package config

import (
	"strings"
	"time"

	"netklit/pkg/logger"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type AppConfig struct {
	App struct {
		Name string `mapstructure:"name"`
		Port string `mapstructure:"port"`
	} `mapstructure:"app"`
	Options struct {
		Timeout time.Duration `mapstructure:"timeout"`
		Workers int           `mapstructure:"workers"`
	} `mapstructure:"options"`
	NS    string `mapstructure:"namespace"`
	Owner string `mapstructure:"owner"`
}

var Config AppConfig

func setup() {
	_ = godotenv.Load(".env")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("env")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		logger.Log.Errorf("fatal error config file: %s", err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		logger.Log.Fatalf("Impossibile decodificare la config: %s", err)
	}

	logger.Log.Infof("app name: %s", Config.App.Name)
	logger.Log.Infof("app port: %s", Config.App.Port)
	logger.Log.Infof("namespace name: %s", Config.NS)
	logger.Log.Infof("owner name: %s", Config.Owner)
	logger.Log.Infof("options timeout: %d", Config.Options.Timeout)
	logger.Log.Infof("options workers: %d", Config.Options.Workers)

}
func Execute() {
	setup()
}
