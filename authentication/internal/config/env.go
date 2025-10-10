package config

import (
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		DbName   string `mapstructure:"dbname"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
	} `mapstructure:"database"`

	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Mail struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Email    string `mapstructure:"email"`
		Password string `mapstructure:"password"`
	} `mapstructure:"mail"`

	Debug     bool   `mapstructure:"debug"`
	SecretKey string `mapstructure:"secret_key"`
}

var Cfg = Config{}
var once sync.Once

func _setConfigs() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	Cfg = cfg

	return nil
}

func NewConfig() error {
	var err error
	once.Do(func() {
		err = _setConfigs()
	})

	return err
}
