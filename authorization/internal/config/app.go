package config

type AppConfig struct {
	APP_NAME          string
	DESCRIPTION       string
	AUTHOR            string
	DEBUG             string
	JWT               string
	DATABASE_NAME     string
	DATABASE_USER     string
	DATABASE_PASSWORD string
	DATABASE_PORT     string
	DATABASE_HOST     string
}

type AppConfigOption func(*AppConfig)

func WithAppName(name string) AppConfigOption {
	return func(c *AppConfig) { c.APP_NAME = name }
}

func WithDescription(desc string) AppConfigOption {
	return func(c *AppConfig) { c.DESCRIPTION = desc }
}

func WithAuthor(author string) AppConfigOption {
	return func(c *AppConfig) { c.AUTHOR = author }
}

func WithDebug(debug string) AppConfigOption {
	return func(c *AppConfig) { c.DEBUG = debug }
}

func WithJWT(jwt string) AppConfigOption {
	return func(c *AppConfig) { c.JWT = jwt }
}

func WithDatabaseName(name string) AppConfigOption {
	return func(c *AppConfig) { c.DATABASE_NAME = name }
}

func WithDatabaseUser(user string) AppConfigOption {
	return func(c *AppConfig) { c.DATABASE_USER = user }
}

func WithDatabasePassword(password string) AppConfigOption {
	return func(c *AppConfig) { c.DATABASE_PASSWORD = password }
}

func WithDatabasePort(port string) AppConfigOption {
	return func(c *AppConfig) { c.DATABASE_PORT = port }
}

func WithDatabaseHost(host string) AppConfigOption {
	return func(c *AppConfig) { c.DATABASE_HOST = host }
}

func NewAppConfig(opts ...AppConfigOption) *AppConfig {
	cfg := &AppConfig{
		APP_NAME:          "GOLANG_APP",
		DESCRIPTION:       "",
		AUTHOR:            "Sajad pourajam",
		DEBUG:             "true",
		JWT:               "",
		DATABASE_NAME:     "GOLANG_APP",
		DATABASE_USER:     "root",
		DATABASE_PASSWORD: "root",
		DATABASE_PORT:     "3306",
		DATABASE_HOST:     "127.0.0.1",
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

