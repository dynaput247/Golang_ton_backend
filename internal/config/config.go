package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	BotToken    string `yaml:"BOT_TOKEN" env-required:"true" env:"BOT_TOKEN"`
	BotID       string `yaml:"BOT_ID" env:"BOT_ID"`
	ChannelID   int64  `yaml:"CHANNEL_ID" env-required:"true" env:"CHANNEL_ID"`
	AdminUserID int64  `yaml:"ADMIN_USER_ID" env-required:"true" env:"ADMIN_USER_ID"`
	TmaURL      string `yaml:"TMA_URL" env-required:"true" env:"TMA_URL"`
}

const (
	configPath = "cfg.yaml"
)

func loadConfigFromFile(path string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, fmt.Errorf("cannot read config from file: %w", err)
	}
	return &cfg, nil
}

func loadConfigFromEnv() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("cannot read config from environment: %w", err)
	}
	return &cfg, nil
}

func MustLoad() * Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	cfg, err := loadConfigFromEnv()
	if err != nil {
		fmt.Printf("Failed to get environment variables: %v. Trying to load from file %s\n", err, configPath)
		cfg, err = loadConfigFromFile(configPath)
		if err != nil {
			log.Fatalf("cannot load config: %v", err)
		}
	}

	return cfg
}