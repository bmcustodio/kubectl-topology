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

package util

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListNodes(kubeClient kubernetes.Interface, o *TopologyOptions) (NodeList, error) {
	for _, p := range []string{
		TopologyLabelPrefix,
		Pre117TopologyLabelPrefix,
	} {
		n, err := kubeClient.CoreV1().Nodes().List(metav1.ListOptions{
			LabelSelector: o.GetLabelSelector(p),
		})
		if err != nil {
			return nil, err
		}
		if len(n.Items) == 0 {
			continue
		}
		r := make(NodeList, 0, len(n.Items))
		for _, nn := range n.Items {
			r = append(r, NewNode(nn))
		}
		return r, nil
	}
	return nil, nil
}
