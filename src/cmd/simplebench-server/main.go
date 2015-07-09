package main

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const headerForwardedFor = "X-Forwarded-For"

type (
	response struct {
		Time     int64
		Hash     string
		ClientIP string
	}
)

func main() {
	http.HandleFunc("/time", handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Unix()
	res := response{
		Time:     currentTime,
		Hash:     fmt.Sprintf("%x", sha512.Sum512([]byte(trueRemoteAddr(r)))),
		ClientIP: trueRemoteAddr(r),
	}
	resj, _ := json.Marshal(res)
	etag := fmt.Sprintf("\"%x\"", sha512.Sum512(resj))
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Etag", etag)
	w.Write(resj)
}

func trueRemoteAddr(req *http.Request) string {
	forwardedFor := req.Header.Get(headerForwardedFor)
	if forwardedFor != "" {
		return strings.Split(forwardedFor, ",")[0]
	}
	return req.RemoteAddr
}
