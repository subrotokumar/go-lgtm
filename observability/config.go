package observability

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Config struct {
	ServiceName    string `env:"SERVICE_NAME"      envDefault:"obs-service"`
	ServiceVersion string `env:"SERVICE_VERSION"   envDefault:"0.0.1"`
	Enabled        bool   `env:"TELEMETRY_ENABLED" envDefault:"true"`
}

func NewConfigFromEnv() (Config, error) {
	telem := Config{}
	if err := env.Parse(&telem); err != nil {
		return Config{}, fmt.Errorf("failed to parse telemetry config: %w", err)
	}

	return telem, nil
}
