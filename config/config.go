package config

type Config struct {
	Server struct {
		Port int
	}
	App struct {
		RateBookingDuration  int    `mapstructure:"rate-booking-duration"`
		DefaultBaseCurrency  string `mapstructure:"default-base-currency"`
		DefaultAdditionalPip int    `mapstructure:"default-additional-pip"`
		ForexRateApiUrl      string `mapstructure:"forex-rate-api-url"`
	}
	Db struct {
		Url      string
		User     string
		Password string
	}
}
