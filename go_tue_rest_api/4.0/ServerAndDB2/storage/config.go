package storage

type Config struct {
	//Строка подключения к бд
	DatabaseURI string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}
