//go:build !oracle

package oracledb

import (
	"context"

	"github.com/grafana/alloy/internal/component"
	"github.com/grafana/alloy/internal/featuregate"
	"github.com/grafana/alloy/internal/runtime/logging/level"
)

func init() {
	component.Register(component.Registration{
		Name:      "prometheus.exporter.oracledb",
		Stability: featuregate.StabilityGenerallyAvailable,
		Args:      StubArguments{},

		Build: func(opts component.Options, args component.Arguments) (component.Component, error) {
			level.Warn(opts.Logger).Log("msg", "prometheus.exporter.oracledb requires the 'oracle' build tag; this component will do nothing")
			return &StubComponent{}, nil
		},
	})
}

// StubArguments is a stub for oracledb arguments when not built with oracle support.
type StubArguments struct{}

// SetToDefault implements syntax.Defaulter.
func (a *StubArguments) SetToDefault() {}

// StubComponent is a no-op component for unsupported platforms.
type StubComponent struct{}

// Run implements component.Component.
func (c *StubComponent) Run(ctx context.Context) error {
	<-ctx.Done()
	return nil
}

// Update implements component.Component.
func (c *StubComponent) Update(args component.Arguments) error {
	return nil
}
