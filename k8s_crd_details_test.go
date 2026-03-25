// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package exportercreator

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer"
)

// stubK8sCRDDetails is a test double for observer.K8sCRD. The observer module used in
// go.mod may not define that type yet; behavior matches contrib's K8sCRD Env/Type.
type stubK8sCRDDetails struct {
	Name        string
	UID         string
	Labels      map[string]string
	Annotations map[string]string
	Namespace   string
	Group       string
	Version     string
	Kind        string
	Spec        map[string]any
	Status      map[string]any
}

func (c *stubK8sCRDDetails) Env() observer.EndpointEnv {
	return map[string]any{
		"uid":         c.UID,
		"name":        c.Name,
		"labels":      c.Labels,
		"annotations": c.Annotations,
		"namespace":   c.Namespace,
		"group":       c.Group,
		"version":     c.Version,
		"kind":        c.Kind,
		"spec":        c.Spec,
		"status":      c.Status,
	}
}

func (*stubK8sCRDDetails) Type() observer.EndpointType {
	return ObserverEndpointTypeK8sCRD
}
