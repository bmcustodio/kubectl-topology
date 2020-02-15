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
	"os"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/bmcstdio/kubectl-topology/internal/version"
)

func init() {
	rootCmd.PersistentFlags().StringP("region", "r", "", "The region to filter resources by. Mutually exclusive with '--zone'.")
	rootCmd.PersistentFlags().StringP("zone", "z", "", "The zone to filter resources by. Mutually exclusive with '--region'.")
	rootCmd.SetVersionTemplate("kubectl-topology " + version.Version)
}

var (
	kubeClient kubernetes.Interface
)

var rootCmd = &cobra.Command{
	Version: version.Version,
	Args:    cobra.NoArgs,
	Use:     "kubectl-topology",
	Short:   "Provides insight into the topology of a Kubernetes cluster.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		r := clientcmd.NewDefaultClientConfigLoadingRules()
		o := &clientcmd.ConfigOverrides{}
		c, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(r, o).ClientConfig()
		if err != nil {
			return err
		}
		if kubeClient, err = kubernetes.NewForConfig(c); err != nil {
			return err
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
