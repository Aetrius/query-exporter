package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

// Define a struct for you collector that contains pointers
// to prometheus descriptors for each metric you wish to expose.
// Note you can also include fields of other types if they provide utility
// but we just won't be exposing them as metrics.
type testCollector struct {
	pingMetric *prometheus.Desc
	upMetric   *prometheus.Desc
}

// You must create a constructor for you collector that
// initializes every descriptor and returns a pointer to the collector
func defaultTestCollector() *testCollector {
	return &testCollector{
		pingMetric: prometheus.NewDesc("ping_metric",
			"Ping Metric",
			nil, nil,
		),
		upMetric: prometheus.NewDesc("up_metric",
			"Up Metric",
			nil, nil,
		),
	}
}

// Each and every collector must implement the Describe function.
// It essentially writes all descriptors to the prometheus desc channel.
func (collector *testCollector) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
	ch <- collector.pingMetric
	ch <- collector.upMetric
}

// Collect implements required collect function for all prometheus collectors
func (collector *testCollector) Collect(ch chan<- prometheus.Metric) {

	//Implement logic here to determine proper metric value to return to prometheus
	//for each descriptor or call other functions that do so.
	var metricValue float64
	metricValue = float64(collectData())

	//Write latest value for each metric in the prometheus metric channel.
	//Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
	ch <- prometheus.MustNewConstMetric(collector.pingMetric, prometheus.CounterValue, metricValue)
	ch <- prometheus.MustNewConstMetric(collector.upMetric, prometheus.CounterValue, metricValue)

}

func collectData() int {

	fmt.Println("Running collect data")
	return 1
}
