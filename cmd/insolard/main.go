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
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"

	"github.com/insolar/insolar/log"

	"github.com/insolar/observer/server"
)

type inputParams struct {
	configPath        string
	genesisConfigPath string
	traceEnabled      bool
}

func parseInputParams() inputParams {
	var rootCmd = &cobra.Command{Use: "insolard"}
	var result inputParams
	rootCmd.Flags().StringVarP(&result.configPath, "config", "c", "", "path to config file")
	rootCmd.Flags().StringVarP(&result.genesisConfigPath, "heavy-genesis", "", "", "path to genesis config for heavy node")
	rootCmd.Flags().BoolVarP(&result.traceEnabled, "trace", "t", false, "enable tracing")
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal("Wrong input params:", err)
	}

	return result
}

func main() {
	params := parseInputParams()
	jww.SetStdoutThreshold(jww.LevelDebug)

	s := server.NewObserverServer(params.configPath, params.traceEnabled)
	s.Serve()
}
