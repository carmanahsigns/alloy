//go:build !oracle

package oracledb_exporter

import (
	"context"
	"net/http"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/grafana/alloy/internal/static/integrations"
	"github.com/grafana/alloy/internal/static/integrations/config"
	integrations_v2 "github.com/grafana/alloy/internal/static/integrations/v2"
	"github.com/grafana/alloy/internal/static/integrations/v2/metricsutils"
	config_util "github.com/prometheus/common/config"
)

// DefaultConfig is the default config for the oracledb v2 integration
var DefaultConfig = Config{
	ConnectionString: config_util.Secret(os.Getenv("DATA_SOURCE_NAME")),
	MaxOpenConns:     10,
	MaxIdleConns:     0,
	QueryTimeout:     5,
	CustomMetrics:    []string{},
}

// Config is the configuration for the oracledb v2 integration
type Config struct {
	ConnectionString config_util.Secret `yaml:"connection_string"`
	MaxIdleConns     int                `yaml:"max_idle_connections"`
	MaxOpenConns     int                `yaml:"max_open_connections"`
	QueryTimeout     int                `yaml:"query_timeout"`
	DefaultMetrics   string             `yaml:"default_metrics,omitempty"`
	CustomMetrics    []string           `yaml:"custom_metrics,omitempty"`
	Username         string             `yaml:"username,omitempty"`
	Password         config_util.Secret `yaml:"password,omitempty"`
}

// UnmarshalYAML implements yaml.Unmarshaler for Config
func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	*c = DefaultConfig

	type plain Config
	return unmarshal((*plain)(c))
}

// Name returns the integration name this config is associated with.
func (c *Config) Name() string {
	return "oracledb"
}

// InstanceKey returns the addr of the oracle instance.
func (c *Config) InstanceKey(agentKey string) (string, error) {
	return "", nil
}

// NewIntegration returns a stub OracleDB Integration
func (c *Config) NewIntegration(logger log.Logger) (integrations.Integration, error) {
	level.Warn(logger).Log("msg", "the oracledb integration requires the 'oracle' build tag; enabling it will do nothing")
	return &stubIntegration{}, nil
}

func init() {
	integrations.RegisterIntegration(&Config{})
	integrations_v2.RegisterLegacy(&Config{}, integrations_v2.TypeMultiplex, metricsutils.NewNamedShim("oracledb"))
}

// stubIntegration implements a no-op integration
type stubIntegration struct{}

// MetricsHandler returns an http.NotFoundHandler to satisfy the Integration interface
func (i *stubIntegration) MetricsHandler() (http.Handler, error) {
	return http.NotFoundHandler(), nil
}

// ScrapeConfigs returns an empty list of scrape configs
func (i *stubIntegration) ScrapeConfigs() []config.ScrapeConfig {
	return []config.ScrapeConfig{}
}

// Run just waits for the context to finish
func (i *stubIntegration) Run(ctx context.Context) error {
	<-ctx.Done()
	return ctx.Err()
}
