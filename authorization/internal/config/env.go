package config

import (
	"fmt"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

var Config = &AppConfig{}

func BootConfig(address string, opts ...AppConfigOption) *AppConfig {
	err := godotenv.Load(address)
	if err != nil {
		panic(fmt.Sprintf("Config error: %s", err))
	}

	cfg := NewAppConfig(opts...)

	v := reflect.ValueOf(cfg).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		envValue := os.Getenv(fieldType.Name)
		if envValue != "" {
			field.SetString(envValue)
		}
	}

	Config = cfg
	return cfg
}
