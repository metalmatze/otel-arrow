// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/component"
)

var (
	Type      = component.MustNewType("concurrentbatch")
	ScopeName = "github.com/open-telemetry/otel-arrow/collector/processor/concurrentbatchprocessor"
)

const (
	TracesStability  = component.StabilityLevelBeta
	MetricsStability = component.StabilityLevelBeta
	LogsStability    = component.StabilityLevelBeta
)
