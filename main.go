// package main

// import (
// 	"github.com/haad/px_dd_exporter/cmd"
// )

// func main() {
// 	cmd.Execute()
// }

package main

import (
	"net/http"
	"time"

	log "github.com/haad/px_dd_exporter/log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cfgListen      = ":2112"
	cfgFilePath    = "/tmp/test.yaml"
	cfgMetricsPath = "/metrics"
	cfgDebug       = false
)

func loadConfigFile() []string {
}

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "definition_data_version",
		Help: "Client definition data version",
	})
)

func main() {
	recordMetrics()
	log.InitLogger(cfgDebug)

	http.Handle(cfgMetricsPath, promhttp.Handler())
	http.ListenAndServe(cfgListen, nil)
}
