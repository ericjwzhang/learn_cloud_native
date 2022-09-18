package main

import (
	"fmt"
	"github.com/ericjwzhang/module10/httpsvr/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 设置环境变量 Version
	err := os.Setenv("VERSION", "v1.0.0")
	if err != nil {
		return
	}
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Printf("os version: %s\n", version)

	//将request的header设置到response中
	for k, v := range r.Header {
		for _, vv := range v {
			//fmt.Printf("Header key: %s, Header value: %s", k, vv)
			w.Header().Set(k, vv)
		}
	}
	ip := getCurrentIP(r)
	log.Printf("Success! Response code: %d", http.StatusOK)
	log.Printf("Success! clientip: %s", ip)
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "welcome to my http server index!")
	if err != nil {
		return
	}
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func health(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "it's working!")
	if err != nil {
		return
	}
}

func getCurrentIP(r *http.Request) string {
	// 先尝试从 X-Forwarded-For 请求头的一个值作为用户的IP
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}

func images(w http.ResponseWriter, r *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	delay := randInt(10, 2000)
	time.Sleep(time.Millisecond * time.Duration(delay))
	w.Write([]byte(fmt.Sprintf("<h1>%d<h1>", delay)))
	log.Printf("Respond in %d ms", delay)
}

func main() {
	log.Printf("Starting http server...")
	metrics.Register()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", health)
	mux.HandleFunc("/images", images)
	mux.HandleFunc("/", index)
	mux.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("start http server faild, error: %s\n", err.Error())
	}
}
