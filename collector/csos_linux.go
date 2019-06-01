// Copyright 2019 Critical stack
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

var csosDesc = prometheus.NewDesc(
	prometheus.BuildFQName(namespace, "csos", "entries"),
	"Current test run.",
	[]string{
		"test_reference",
		"exec_params",
	},
	nil,
)

type csosCollector struct {
}
type csosperfdata struct {
	TestReference string
	ExecParams    string
}

func init() {
	registerCollector("csos", defaultEnabled, newCsosCollector)
}

// NewCsosCollector returns new csosCollector.
func newCsosCollector() (Collector, error) {
	return &csosCollector{}, nil
}

func getCsosPerfTag() (csosperfdata, error) {
	myData := csosperfdata{TestReference: "Sludge", ExecParams: "Blah"}
	return myData, error(nil)
}

func (c *csosCollector) Update(ch chan<- prometheus.Metric) error {
	csos, err := getCsosPerfTag()
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(csosDesc, prometheus.GaugeValue, 1,
		csos.TestReference,
		csos.ExecParams,
	)

	return nil
}
