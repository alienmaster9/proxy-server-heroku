package main

import (
    "github.com/AaronO/gogo-proxy"
    "net/http"
    "os"
)

func main() {
    p, _ := proxy.New(proxy.ProxyOptions{
        Balancer: func(req *http.Request) (string, error) {
            req.Header.Set("X-Forwarded-For", "14.1.32.0")
            req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")
            return os.Getenv("UPSTREAM_SERVER"), nil
        },
    })

    http.ListenAndServe(":"+os.Getenv("PORT"), p)
}
