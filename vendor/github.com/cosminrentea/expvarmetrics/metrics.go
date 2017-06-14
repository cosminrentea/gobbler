// Package metrics implements simple general counter-metrics.
// Metrics are enabled by default. If you want to disable metrics, build with:
// go build -tags noexpvarmetrics
package metrics

import (
	log "github.com/Sirupsen/logrus"

	"expvar"
	"fmt"
	"io"
	"net/http"
	"runtime"
)

var (
	logger        = log.WithField("module", "metrics")
	numGoroutines = expvar.NewInt("num_goroutines")
)

const (
	// DefaultAverageLatencyJSONValue is the default value shown in JSON when no value can be computed.
	DefaultAverageLatencyJSONValue = "\"\""
	MilliPerNano                   = 1000000
)

// HttpHandler is a HTTP handler writing the current metrics to the http.ResponseWriter
func HttpHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	writeMetrics(rw)
}

func writeMetrics(w io.Writer) {
	numGoroutines.Set(int64(runtime.NumGoroutine()))
	fmt.Fprint(w, "{\n")
	first := true
	expvar.Do(func(kv expvar.KeyValue) {
		if !first {
			fmt.Fprint(w, ",\n")
		}
		first = false
		fmt.Fprintf(w, "%q: %s", kv.Key, kv.Value)
	})
	fmt.Fprint(w, "\n}\n")
}

// LogOnDebugLevel logs all the current metrics, if logging is on Debug level.
func LogOnDebugLevel() {
	if log.GetLevel() == log.DebugLevel {
		fields := log.Fields{}
		expvar.Do(func(kv expvar.KeyValue) {
			fields[kv.Key] = kv.Value
		})
		logger.WithFields(fields).Debug("current values of metrics")
	}
}
