package config

type SMTPConfig struct {
	Host   string
	Port   int
	Sender string
	Pass   string
	To     string
}
