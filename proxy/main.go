package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var config = LoadConfig("config.json")

type App struct{}

func (app *App) Run(config ResponseConfig) {
	fmt.Println(config)
	remote, err := url.Parse(config.RemoteHost)
	errorPanic(err)

	proxy := httputil.NewSingleHostReverseProxy(remote)
	handler := &ProxyHandler{proxy}

	err = http.ListenAndServe(config.LocalAddress, handler)
	errorPanic(err)
}

func main() {
	app := App{}
	app.Run(config)
}
