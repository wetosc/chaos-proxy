package main

import (
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
)

const (
	Normal         = iota
	Err404         = iota
	Err500         = iota
	CustomResponse = iota
)

type ProxyHandler struct {
	proxy *httputil.ReverseProxy
}

func responseType(c *ResponseConfig) int {
	if !c.CheckConfig() {
		panic("Invalid configuration")
	}
	x := float32(0)
	r := rand.Float32()

	if (0 < r) && (r < x+c.Per404) {
		return Err404
	}
	x += c.Per404
	if (0 < r) && (r < x+c.Per500) {
		return Err500
	}
	x += c.Per500
	if (0 < r) && (r < x+c.PerCustom) {
		return CustomResponse
	}

	return Normal
}

func (handler *ProxyHandler) ServeHTTP(writter http.ResponseWriter, req *http.Request) {
	rspType := responseType(&config)
	log.Println(rspType)

	switch rspType {
	case Normal:
		handler.proxy.ServeHTTP(writter, req)
	case Err404:
		http.NotFound(writter, req)
	case Err500:
		http.Error(writter, "Internal server error", 500)
	}

}

func errorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
