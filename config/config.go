package config

import (
	"github.com/caarlos0/env/v10"
	"log"
	"strconv"
)

type Config struct {
	Hostname   string `env:"HOSTNAME" envDefault:"localhost"`
	Port       uint16 `env:"PORT" envDefault:"8123"`
	VolumePath string `env:"VOLUME_PATH,expand" envDefault:"volume/"`
}

func LoadConfig() *Config {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	log.Printf("Loaded Config: %+v\n", cfg)

	return cfg
}

func (cfg *Config) GetHostWithPort() string {
	return cfg.Hostname + ":" + strconv.FormatUint(uint64(cfg.Port), 10)
}
