//go:build !oracle

package build

import (
	"github.com/grafana/alloy/internal/component/discovery"
	"github.com/grafana/alloy/internal/static/integrations/oracledb_exporter"
)

func (b *ConfigBuilder) appendOracledbExporter(config *oracledb_exporter.Config, instanceKey *string) discovery.Exports {
	// Oracle DB exporter is not available without the 'oracle' build tag
	return discovery.Exports{}
}
