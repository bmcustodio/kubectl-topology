// Copyright 2020 bmcustodio
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
	corev1 "k8s.io/api/core/v1"

	"github.com/bmcustodio/kubectl-topology/internal/util"
)

func init() {
	rootCmd.AddCommand(podCmd)
	podCmd.PersistentFlags().BoolP("all-namespaces", "A", false, "List pods across all namespaces.")
	podCmd.PersistentFlags().StringP("selector", "l", "", "A label selector to filter out pods.")
}

var podCmd = &cobra.Command{
	Args:  cobra.NoArgs,
	Use:   "pod",
	Short: "Provides insight into the distribution of pods per region or zone.",
	RunE: func(cmd *cobra.Command, args []string) error {
		a, err := cmd.Flags().GetBool("all-namespaces")
		if err != nil {
			return err
		}
		n, _ := cmd.Flags().GetString("namespace")
		if a {
			n = corev1.NamespaceAll
		}
		r, _ := cmd.Flags().GetString("region")
		z, _ := cmd.Flags().GetString("zone")
		l, err := cmd.Flags().GetString("selector")
		if err != nil {
			return err
		}
		o, err := util.NewTopologyOptions(r, z, n)
		if err != nil {
			return err
		}
		p, err := util.ListPods(kubeClient, o, l)
		if err != nil {
			return err
		}
		return util.PrintResult(p, false)
	},
}
