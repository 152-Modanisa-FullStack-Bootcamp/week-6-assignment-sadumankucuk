package config

type Config struct {
	InitialBalanceAmount int `json:"initialBalanceAmount"`
	MinimumBalanceAmount int `json:"minimumBalanceAmount"`
}

var c = &Config{}

func Get() *Config {
	return c
}
