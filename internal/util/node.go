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

package util

import (
	corev1 "k8s.io/api/core/v1"
)

type Node struct {
	Name   string
	Region string
	Zone   string
}

func NewNode(node corev1.Node) Node {
	r := Node{
		Name: node.Name,
	}
	if v, exists := node.Labels[RegionLabel]; exists && v != "" {
		r.Region = v
	} else {
		r.Region = node.Labels[Pre117RegionLabel]
	}
	if v, exists := node.Labels[ZoneLabel]; exists && v != "" {
		r.Zone = v
	} else {
		r.Zone = node.Labels[Pre117ZoneLabel]
	}
	return r
}

type NodeList []Node

func (l NodeList) Headers() string {
	return "NAME\tREGION\tZONE\n"
}

func (l NodeList) Items() []string {
	r := make([]string, 0, len(l))
	for _, ll := range l {
		r = append(r, ll.Name+"\t"+ll.Region+"\t"+ll.Zone+"\n")
	}
	return r
}

func (l NodeList) Length() int {
	return len(l)
}
