package main

type Config struct {
	Server struct {
		Port int
	}
	App struct {
		RateBookingDuration  int    `config:"rate-booking-duration"`
		DefaultBaseCurrency  string `config:"default-base-currency"`
		DefaultAdditionalPip int    `config:"default-additional-pip"`
		ForexRateApiUrl      string `config:"forex-rate-api-url"`
	}
	Db struct {
		Url      string
		User     string
		Password string
	}
}

type FlatConfig struct {
	ServerPort           int    `config:"server.port"`
	RateBookingDuration  int    `config:"app.rate-booking-duration"`
	DefaultBaseCurrency  string `config:"app.default-base-currency"`
	DefaultAdditionalPip int    `config:"app.default-additional-pip"`
	ForexRateApiUrl      string `config:"app.forex-rate-api-url"`
	DbUrl                string `config:"db.url"`
	DbUser               string `config:"db.user"`
	DbPassword           string `config:"db.password"`
}
