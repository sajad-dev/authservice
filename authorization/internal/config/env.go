package config

import (
	"os"
	"reflect"
	"sync"
)

var instanse = AppConfig{}
var once sync.Once

func _setConfigs() AppConfig {
	cfg := NewAppConfig()

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

	return *cfg
}

func NewConfig() AppConfig {
	once.Do(func() {
		instanse = _setConfigs()
	})

	return instanse
}
