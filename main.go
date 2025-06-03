package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var counter = promauto.NewCounter(prometheus.CounterOpts{
    Name: "api_calls_total",
    Help: "The total number of processed API calls",
})

type Simple struct {
    Name        string
    Description string
    Url         string
}

func SimpleFactory(host string) Simple {
    return Simple{"Hello", "Dear Students!!!", host}
}

func handler(w http.ResponseWriter, r *http.Request) {
    // simple := Simple{"Hello", "Dear Students!", r.Host}
    simple := SimpleFactory(r.Host)

    jsonOutput, _ := json.Marshal(simple)

    if r.URL.Path == "/" {
        counter.Inc()
    }

    fmt.Fprintln(w, string(jsonOutput))
}

func main() {
    fmt.Println("Server started on port 4444")
    http.Handle("/metrics", promhttp.Handler())
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":4444", nil))
}
