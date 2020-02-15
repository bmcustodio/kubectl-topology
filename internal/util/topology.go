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
	"fmt"
)

type TopologyOptions struct {
	Namespace string
	Region    string
	Zone      string
}

func NewTopologyOptions(region, zone, namespace string) (*TopologyOptions, error) {
	if len(region)*len(zone) != 0 {
		return nil, fmt.Errorf("region and zone are mutually exclusive")
	}
	return &TopologyOptions{
		Namespace: namespace,
		Region:    region,
		Zone:      zone,
	}, nil
}

func (o *TopologyOptions) GetLabelSelector(prefix string) string {
	switch {
	case len(o.Region) > 0:
		return fmt.Sprintf("%s/%s=%s", prefix, RegionSuffix, o.Region)
	case len(o.Zone) > 0:
		return fmt.Sprintf("%s/%s=%s", prefix, ZoneSuffix, o.Zone)
	default:
		return ""
	}
}
