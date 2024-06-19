package types

type Config struct {
	APIPort    int
	SMTPPort   int
	SMTPHost   string
	SMTPUser   string
	SMTPPass   string
	SMTPUseTLS bool
}
