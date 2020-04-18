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
	"sort"

	corev1 "k8s.io/api/core/v1"
)

type Pod struct {
	Name      string
	Namespace string
	Node      Node
}

func NewPod(pod corev1.Pod, node Node) Pod {
	return Pod{
		Name:      pod.Name,
		Namespace: pod.Namespace,
		Node:      node,
	}
}

type PodList []Pod

func (l PodList) Headers() string {
	return "NAMESPACE\tNAME\tNODE\tREGION\tZONE\n"
}

func (l PodList) Items() []string {
	sort.SliceStable(l, func(i, j int) bool {
		return l[i].Namespace < l[j].Namespace && l[i].Name < l[j].Name && l[i].Node.Name < l[j].Node.Name
	})
	r := make([]string, 0, len(l))
	for _, ll := range l {
		r = append(r, ll.Namespace+"\t"+ll.Name+"\t"+ll.Node.Name+"\t"+ll.Node.Region+"\t"+ll.Node.Zone+"\n")
	}
	return r
}

func (l PodList) Length() int {
	return len(l)
}
