// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package exportercreator

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/exporter/exportertest"

	"github.com/stuart23/exportercreator/internal/metadata"
)

func TestCreateExporter_SameInstancePerConfig(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	params := exportertest.NewNopSettings(metadata.Type)
	ctx := t.Context()

	lExp, err := factory.CreateLogs(ctx, params, cfg)
	require.NoError(t, err)
	mExp, err := factory.CreateMetrics(ctx, params, cfg)
	require.NoError(t, err)
	tExp, err := factory.CreateTraces(ctx, params, cfg)
	require.NoError(t, err)

	require.Same(t, lExp, mExp)
	require.Same(t, lExp, tExp)
}
