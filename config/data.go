package config

import (
	"github.com/bemobi/envconfig"
)

// Data contains parameters that controls the execution of the server.
// The struct is filled from environment variables.
var Data struct {
	Server struct {
		Addr string `envconfig:"default=0.0.0.0:8080"`
	}
	Redis struct {
		Addr     string `envconfig:"optional"`
		Password string `envconfig:"optional"`
		DB       int    `envconfig:"default=0"`
	}
}

func init() {
	if err := envconfig.Init(&Data); err != nil {
		panic(err)
	}
}
