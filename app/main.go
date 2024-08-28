package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/push"
)

var (
	g = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "latency_api",
		Help:    "Sample metric for latency api",
		Buckets: prometheus.DefBuckets,
	}, []string{"method", "endpoint", "duration"})
)

func init() {
	prometheus.Register(g)
}

func main() {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		time.Sleep(time.Duration(rd.Intn(1000)) * time.Millisecond)
		w.WriteHeader(http.StatusOK)

		b, _ := json.Marshal("OK")
		w.Write(b)

		end := time.Since(now)
		g.WithLabelValues(r.Method, r.URL.Path, strconv.FormatFloat(end.Seconds(), 'E', -1, 64)).Observe(end.Seconds())
	})

	http.Handle("/metrics", promhttp.Handler())

	err := push.New("http://pushgateway:9091", "calculate_job").Collector(g).Add()
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}

	fmt.Println("Server is running on port: 8081")
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
