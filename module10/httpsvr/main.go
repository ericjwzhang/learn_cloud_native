package main

import (
	"flag"
	"fmt"
	"github.com/ericjwzhang/learn_cloud_native/module10/metrics"
	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"net/http/pprof"
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

func healthz(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "it's working!")
	if err != nil {
		return
	}
}

func printLog(w http.ResponseWriter, r *http.Request) {
	flag.Set("v", "4")
	glog.V(2).Info("This is my test log.")

}

func getCurrentIP(r *http.Request) string {
	// 先尝试从 X-Forwarded-For 请求头的一个值作为用户的IP
	ip := r.Header.Get("X-Forwarded-For")
	//log.Printf(r.RemoteAddr)
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}

func images(w http.ResponseWriter, r *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	randInt := rand.Intn(2000)
	time.Sleep(time.Millisecond * time.Duration(randInt))
	w.Write([]byte(fmt.Sprintf("<h1>%d<h1>", randInt)))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/logs", printLog)
	mux.HandleFunc("/images", images)
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("start http server faild, error: %s\n", err.Error())
	}
}
