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

func ListPods(kubeClient kubernetes.Interface, o *TopologyOptions, s string) (PodList, error) {
	n, err := ListNodes(kubeClient, o)
	if err != nil {
		return nil, err
	}
	r := make(PodList, 0)
	for _, nn := range n {
		p, err := kubeClient.CoreV1().Pods(o.Namespace).List(metav1.ListOptions{
			FieldSelector: "spec.nodeName=" + nn.Name,
			LabelSelector: s,
		})
		if err != nil {
			return nil, err
		}
		for _, pp := range p.Items {
			r = append(r, NewPod(pp, nn))
		}
	}
	return r, nil
}
