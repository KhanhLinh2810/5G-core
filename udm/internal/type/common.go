package types

import "github.com/caarlos0/env/v6"

type SNssai struct {
	sst int    `json:"sst"`
	sd  string `json:"sd"`
}

type Config struct {
	Environment string `env:"ENVIRONMENT,required"`
	Version     int    `env:"VERSION,required"`
}