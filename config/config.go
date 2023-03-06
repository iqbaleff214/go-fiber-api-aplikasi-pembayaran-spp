package config

import (
	"os"
)

var instance config

type Config interface {
	DatabaseURI() string
	SecretJWT() string
	Port() string
}

type config struct {
	port        string
	databaseURI string
	secretJWT   string
}

func init() {
	instance = config{
		databaseURI: os.Getenv("DATABASE_URI"),
		secretJWT:   os.Getenv("JWT_SECRET_KEY"),
		port:        os.Getenv("PORT"),
	}

	if instance.databaseURI == "" {
		instance.databaseURI = "root:root@tcp(127.0.0.1:3306)/aplikasi_pembayaran_spp?parseTime=true"
	}

	if instance.secretJWT == "" {
		instance.secretJWT = "iqbaleff214"
	}

	if instance.port == "" {
		instance.port = ":8000"
	}
}

func NewConfig() Config {
	return instance
}

func (c config) DatabaseURI() string {
	return c.databaseURI
}

func (c config) SecretJWT() string {
	return c.secretJWT
}

func (c config) Port() string {
	return c.port
}
