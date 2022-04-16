package config

type Config struct {
	Token     string
	BotPrefix string
}

func (config Config) GetToke() string {
	return config.Token
}

func (config Config) GetBotPrefix() string {
	return config.BotPrefix
}
