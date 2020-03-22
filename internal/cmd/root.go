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
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"

	"github.com/bmcstdio/kubectl-topology/internal/version"
)

func init() {
	configFlags = genericclioptions.NewConfigFlags(true)
	configFlags.AddFlags(rootCmd.PersistentFlags())
	rootCmd.PersistentFlags().StringP("region", "r", "", "The region to filter resources by. Mutually exclusive with '--zone'.")
	rootCmd.PersistentFlags().StringP("zone", "z", "", "The zone to filter resources by. Mutually exclusive with '--region'.")
	rootCmd.SetVersionTemplate("kubectl-topology " + version.Version)
}

var (
	configFlags *genericclioptions.ConfigFlags
	kubeClient  kubernetes.Interface
)

var rootCmd = &cobra.Command{
	Version:      version.Version,
	Args:         cobra.NoArgs,
	Use:          "kubectl-topology",
	SilenceUsage: true,
	Short:        "Provides insight into the topology of a Kubernetes cluster.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		kubeConfig, err := configFlags.ToRESTConfig()
		if err != nil {
			return err
		}
		kubeClient, err = kubernetes.NewForConfig(kubeConfig)
		if err != nil {
			return err
		}
		return nil
	},
}

func Execute() {
	pflag.CommandLine = pflag.NewFlagSet("kubectl-topology", pflag.ExitOnError)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
