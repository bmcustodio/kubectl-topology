// Copyright 2020 bmcstdio
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

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/bmcstdio/kubectl-topology/internal/util"
)

func init() {
	rootCmd.AddCommand(nodeCmd)
}

var nodeCmd = &cobra.Command{
	Args:  cobra.NoArgs,
	Use:   "node",
	Short: "Provides insight into the distribution of nodes per region or zone.",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, _ := cmd.Flags().GetString("region")
		z, _ := cmd.Flags().GetString("zone")
		o, err := util.NewTopologyOptions(r, z, "")
		if err != nil {
			return err
		}
		n, err := util.ListNodes(kubeClient, o)
		if err != nil {
			return err
		}
		return util.PrintResult(n, false)
	},
}
