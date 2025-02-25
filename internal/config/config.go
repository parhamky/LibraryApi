package config

import "os"

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type HttpConfig struct {
	Url  string
	Port string
}

type CacheConfig struct {
	Host     string
	Port     string
	Password string
	DbName   string
}

func IsTest() string {
	return os.Getenv("IS_TEST")
}

func LoadDBConfig() DBConfig {
	return DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

func LoadCacheConfig() CacheConfig {
	return CacheConfig{
		Host:     os.Getenv("CACHE_HOST"),
		Port:     os.Getenv("CACHE_PORT"),
		Password: os.Getenv("CACHE_PASSWORD"),
		DbName:   os.Getenv("CACHE_DB"),
	}
}

func LoadHttpConfig() HttpConfig {
	return HttpConfig{
		Url:  os.Getenv("HTTP_URL"),
		Port: os.Getenv("HTTP_PORT"),
	}
}
