// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package exportercreator // import "github.com/astronomer/exportercreator"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer"
)

// Observer endpoint kinds used by k8s_observer, jsonfile_observer, etc. Older tagged
// github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer modules
// may not export named constants for every kind; these values match the strings returned
// by real observer implementations.
const (
	ObserverEndpointTypeK8sCRD   = observer.EndpointType("k8s.crd")
	ObserverEndpointTypeJSONFile = observer.EndpointType("jsonfile")
)
