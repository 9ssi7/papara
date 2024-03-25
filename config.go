package papara

type Config struct {
	ApiKey string `env:"PAPARA_API_KEY"`
	ApiUrl string `env:"PAPARA_API_URL"`
}
