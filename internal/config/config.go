package config

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	DefaultDevConfigPath = "config.yaml"
	DefaultPort          = 8080
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Admin    AdminConfig    `yaml:"admin"`
	Database DatabaseConfig `yaml:"database"`
	Paths    PathsConfig    `yaml:"paths"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type AdminConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type DatabaseConfig struct {
	Path string `yaml:"path"`
}

type PathsConfig struct {
	DataDir string `yaml:"data_dir"`
}

func Load(path string) (Config, error) {
	if path == "" {
		path = defaultConfigPath()
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("reading config file %q: %w", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("parsing config file %q: %w", path, err)
	}

	if cfg.Server.Port == 0 {
		cfg.Server.Port = DefaultPort
	}
	if cfg.Server.Port < 1 || cfg.Server.Port > 65535 {
		return Config{}, errors.New("server.port must be between 1 and 65535")
	}

	if strings.TrimSpace(cfg.Database.Path) == "" {
		cfg.Database.Path = defaultDBPath()
	}
	if strings.TrimSpace(cfg.Paths.DataDir) == "" {
		cfg.Paths.DataDir = defaultDataDir()
	}

	cfg.Admin.Username = strings.TrimSpace(cfg.Admin.Username)
	cfg.Admin.Password = strings.TrimSpace(cfg.Admin.Password)
	if cfg.Admin.Username == "" || cfg.Admin.Password == "" {
		return Config{}, errors.New("admin.username and admin.password must be set")
	}

	return cfg, nil
}

func defaultConfigPath() string {
	if runtime.GOOS == "freebsd" {
		return "/usr/local/etc/trigexmoe.yaml"
	}
	return DefaultDevConfigPath
}

func defaultDBPath() string {
	if runtime.GOOS == "freebsd" {
		return "/var/db/trigexmoe.sqlite"
	}
	return "data/trigexmoe.sqlite"
}

func defaultDataDir() string {
	if runtime.GOOS == "freebsd" {
		return "/usr/local/share/trigexmoe/data"
	}
	return "data"
}
