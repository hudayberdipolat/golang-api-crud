package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	DbConfig     dbConfig     `json:"db_config"`
	ServerConfig serverConfig `json:"server_config"`
	PublicPath   string       `json:"public_path" env:"PUBLIC_PATH"`
	RootPath     string       `json:"root_path" env:"ROOT_PATH"`
}

type dbConfig struct {
	DbHost     string `json:"db_host" env:"DB_HOST" `
	DbPort     string `json:"db_port" env:"DB_PORT"`
	DbUser     string `json:"db_user" env:"DB_USER"`
	DbPassword string `json:"db_password" env:"DB_PASSWORD"`
	DbName     string `json:"db_name" env:"DB_NAME"`
	DbSllMode  string `json:"db_sll_mode" env:"DB_SLL_MODE"`
}

type serverConfig struct {
	ServerHost string `json:"server_host" env:"SERVER_HOST"`
	ServerPort string `json:"server_port" env:"SERVER_PORT"`
	AppName    string `json:"app_name" env:"APP_NAME"`
	AppHeader  string `json:"app_header" env:"APP_HEADER"`
}

func GetAppConfig() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig("../../.env", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, err
}
