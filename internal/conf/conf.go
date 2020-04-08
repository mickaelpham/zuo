package conf

import (
	"log"
	"syscall"
)

type Conf struct {
	ClientId     string
	ClientSecret string
	BaseUrl      string
}

func FromEnv() *Conf {
	return &Conf{
		ClientId:     env("ZUO_CLIENT_ID"),
		ClientSecret: env("ZUO_CLIENT_SECRET"),
		BaseUrl:      env("ZUO_BASE_URL"),
	}
}

func env(name string) (value string) {
	value, found := syscall.Getenv(name)
	if !found {
		log.Fatalf("Missing %q environment variable\n", name)
	}
	return
}
