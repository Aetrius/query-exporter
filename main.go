package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	//Create a new instance of the foocollector and
	//register it with the prometheus client.
	foo := defaultTestCollector()
	prometheus.MustRegister(foo)
	log.Info("logging foo" + foo.pingMetric.String())

	//This section will start the HTTP server and expose
	//any metrics on the /metrics endpoint.
	http.Handle("/metrics", promhttp.Handler())
	log.Info("Beginning to serve on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func head(s string) bool {
	r, e := http.Head(s)
	return e == nil && r.StatusCode == 200
}
