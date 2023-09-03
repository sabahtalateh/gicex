package config

import (
	"fmt"
	"github.com/sabahtalateh/gic"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type App struct {
	StartTimeout time.Duration `yaml:"start_timeout"`
	StopTimeout  time.Duration `yaml:"stop_timeout"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type Config struct {
	App App `yaml:"app"`
	DB  DB  `yaml:"db"`
}

func init() {
	gic.Add[Config](
		gic.WithInitE(func() (Config, error) {
			var (
				c Config
			)

			file := os.Getenv("CONFIG_FILE")
			if file == "" {
				return c, fmt.Errorf("provide CONFIG_FILE env var")
			}

			bb, err := os.ReadFile(file)
			if err != nil {
				return c, err
			}

			if err = yaml.Unmarshal(bb, &c); err != nil {
				return c, err
			}

			return c, nil
		}),
	)
}
