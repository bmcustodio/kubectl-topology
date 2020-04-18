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
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

type Result interface {
	Headers() string
	Items() []string
	Length() int
}

func PrintResult(r Result, noHeaders bool) error {
	if r.Length() == 0 {
		fmt.Println(NoResourcesFoundMessage)
		return nil
	}
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 3, ' ', 0)
	if !noHeaders {
		fmt.Fprint(w, r.Headers())
	}
	d := r.Items()
	sort.SliceStable(d, func(i, j int) bool {
		return d[i] < d[j]
	})
	for _, rr := range d {
		fmt.Fprint(w, rr)
	}
	return w.Flush()
}
