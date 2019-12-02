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

package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/log"

	insconf "github.com/insolar/insolar/configuration"

	"github.com/insolar/observer/component"
	"github.com/insolar/observer/configuration"
)

var stop = make(chan os.Signal, 1)

func main() {
	cfg := configuration.Load()
	logger, err := log.NewLog(insconf.Log{
		Level:      cfg.Log.Level,
		Formatter:  cfg.Log.Format,
		Adapter:    "zerolog",
		OutputType: "stderr",
		BufferSize: 0,
	})
	if err != nil {
		log.Fatalf("Can't create logger: %s", err.Error())
	}
	ctx := inslogger.SetLogger(context.Background(), logger)

	manager := component.Prepare(ctx, cfg)
	manager.Start()
	graceful(logger, manager.Stop)
}

func graceful(logger insolar.Logger, that func()) {
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	logger.Infof("gracefully stopping...")
	that()
}
