package main

type Config struct {
	Server struct {
		Port int
	}
	App struct {
		RateBookingDuration  int    `koanf:"rate-booking-duration"`
		DefaultBaseCurrency  string `koanf:"default-base-currency"`
		DefaultAdditionalPip int    `koanf:"default-additional-pip"`
		ForexRateApiUrl      string `koanf:"forex-rate-api-url"`
	}
	Db struct {
		Url      string
		User     string
		Password string
	}
}
