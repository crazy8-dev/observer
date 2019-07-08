//
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package internal

import (
	"context"

	"github.com/insolar/insolar/configuration"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/instrumentation/instracer"
)

// Jaeger is a default insolar tracer preset.
func Jaeger(
	ctx context.Context,
	cfg configuration.JaegerConfig,
	traceID, nodeRef, nodeRole string,
) (context.Context, func()) {
	inslogger.FromContext(ctx).Infof(
		"Tracing enabled. Agent endpoint: '%s', collector endpoint: '%s'\n",
		cfg.AgentEndpoint,
		cfg.CollectorEndpoint,
	)
	flush := instracer.ShouldRegisterJaeger(
		ctx,
		nodeRole,
		nodeRef,
		cfg.AgentEndpoint,
		cfg.CollectorEndpoint,
		cfg.ProbabilityRate,
	)
	ctx = instracer.SetBaggage(ctx, instracer.Entry{Key: "traceid", Value: traceID})
	return ctx, flush
}
