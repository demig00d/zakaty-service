package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"gopkg.in/yaml.v3"
	"unicode/utf8"
)

type (
	// Config -.
	Config struct {
		Token string `yaml:"token env:"TOKEN""`
		Sheet `yaml:"sheet"`
		HTTP  `yaml:"http"`
		Log   `yaml:"logger"`
	}

	// Sheet -.
	Sheet struct {
		SpreadsheetId string  `yaml:"spreadsheet-id"`
		Range         string  `yaml:"range"`
		Columns       Columns `yaml:"columns"`
	}

	// Columns -.
	Columns struct {
		TelegramId Rune `yaml:"telegram-id"`
		Firstname  Rune `yaml:"firstname"`
		Lastname   Rune `yaml:"lastname"`
		Patronym   Rune `yaml:"patronym"`
		Sum        Rune `yaml:"sum"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}
)

// NewConfig returns app config.
func NewConfig() (Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./config/config.yml", cfg)

	return *cfg, err
}

type Rune rune

func (r *Rune) UnmarshalYAML(n *yaml.Node) error {
	var s string
	if err := n.Decode(&s); err != nil {
		return err
	}

	rn, _ := utf8.DecodeRune([]byte(s))
	*r = Rune(rn)
	return nil
}
