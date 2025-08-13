package config

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Theme        string              `yaml:"theme"`
	AutoCheck    bool                `yaml:"autoCheck"`
	TimerEnabled bool                `yaml:"timerEnabled"`
	Bindings     map[string][]string `yaml:"bindings"`
}

func Default() Config {
	return Config{
		Theme:        "dark",
		AutoCheck:    true,
		TimerEnabled: true,
		Bindings:     map[string][]string{},
	}
}

func path() (string, error) {
	h, err := os.UserHomeDir()
	if err != nil { return "", err }
	return filepath.Join(h, ".punkdoku", "config.yaml"), nil
}

func Load() (Config, error) {
	cfg := Default()
	p, err := path()
	if err != nil { return cfg, err }
	b, err := os.ReadFile(p)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return cfg, nil
		}
		return cfg, err
	}
	if err := yaml.Unmarshal(b, &cfg); err != nil { return cfg, err }
	return cfg, nil
}

func Save(cfg Config) error {
	p, err := path()
	if err != nil { return err }
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil { return err }
	data, err := yaml.Marshal(cfg)
	if err != nil { return err }
	return os.WriteFile(p, data, 0o644)
}
