package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"context"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if p, ok := w.(http.Pusher); ok {
		p.Push("/node_modules/todomvc-app-css/index.css", nil)
		p.Push("/dist/bundle.js", nil)
		p.Push("/node_modules/todomvc-common/base.js", nil)
		p.Push("/learn.json", nil)
	}

	contents := map[string]string{
		"css": "text/css",
		"js": "text/js",
		"json": "application/json",
		"html": "text/html",
	}

	req := strings.Split(r.URL.Path, ".")
	format := req[len(req)-1]

	file := "public/vanilla-es6"
	if (r.URL.Path == "/") {
		file += "/index.html"
	} else {
		file += r.URL.Path
	}

	w.Header().Set("content-type", contents[format])

	data, _ := ioutil.ReadFile(file)
	fmt.Fprintln(w, string(data))
}

func main() {
	server := http.Server{Addr: ":3001"}

	http.HandleFunc("/", rootHandler)

	go func () {
		time.Sleep(time.Second * 20)
		if err := server.Shutdown(context.Background()); err != nil {
			fmt.Println(err)
		}
	}()

	err := server.ListenAndServeTLS("cert.pem", "key.pem")
	if err != http.ErrServerClosed {
		fmt.Println(err)
	}

	fmt.Println("Server gracefully stopped")
}