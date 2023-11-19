package main

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var onlineUsers = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "goappOnlineUsers",
	Help: "Online Users",
	ConstLabels: map[string]string{
		"course": "fullcycle",
	},
})

func main() {
	r := prometheus.NewRegistry()

	r.MustRegister(onlineUsers)

	go func(){
		onlineUsers.Set(float64(rand.Intn(2000)))
	}()

	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":8181", nil))
}
