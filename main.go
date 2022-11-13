// package main

// import (
// 	"github.com/haad/px_dd_exporter/cmd"
// )

// func main() {
// 	cmd.Execute()
// }

package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cfgListen      = ":2112"
	cfgFilePath    = "/tmp/test.yaml"
	cfgMetricsPath = "/metrics"
	cfgDebug       = false
	cfgBaseDir     = "/Users/haad/Devel/go/src/github.com/haad/px_dd_expoter/tests"
)

// func loadConfigFile() []string {
// 	return [""]
// }

func visit(path string, di fs.DirEntry, err error) error {
	r, _ := regexp.Compile("client-data.*.sqlite")

	if r.MatchString(path) {
		fmt.Printf("Visited: %s\n", filepath.Base(path))
		clienDefinitionMetrics(prometheus.Labels{"game": "ts2", "origin": filepath.Base(filepath.Dir(filepath.Dir(path))), "version": strings.ReplaceAll(filepath.Base(path), ".sqlite", "")})
	}
	return nil
}

func clienDefinitionMetrics(labels map[string]string) {
	if err := prometheus.Register(prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Namespace:   "client",
			Name:        "definition_data_version",
			Help:        "Client definition data version.",
			ConstLabels: labels,
		},
		func() float64 { return float64(1) },
	)); err == nil {
		fmt.Printf("GaugeFunc 'definition_data_version' registered with labels: %s \n", labels)
	}
}

func main() {
	// clienDefinitionMetrics(prometheus.Labels{"game": "ts2", "version": "203"})
	// clienDefinitionMetrics(prometheus.Labels{"game": "cinc", "version": "203"})
	// clienDefinitionMetrics(prometheus.Labels{"game": "da2", "version": "203"})
	// clienDefinitionMetrics(prometheus.Labels{"game": "sp2", "version": "203"})
	// log.InitLogger(cfgDebug)

	err := filepath.WalkDir(cfgBaseDir, visit)
	fmt.Printf("filepath.WalkDir() returned %v\n", err)

	http.Handle(cfgMetricsPath, promhttp.Handler())
	http.ListenAndServe(cfgListen, nil)
}
