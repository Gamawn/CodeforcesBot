package bot

type Config struct {
	BotToken string `toml:"bot_token"`
}

func NewConfig() *Config {
	return &Config{}
}
