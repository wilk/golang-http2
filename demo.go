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

		// push resources only when the index.html has been requested
		if p, ok := w.(http.Pusher); ok {
			err := p.Push("/node_modules/todomvc-app-css/index.css", nil)
			if err != nil {
				fmt.Println(err)
			}
			err = p.Push("/dist/bundle.js", nil)
			if err != nil {
				fmt.Println(err)
			}
			err = p.Push("/node_modules/todomvc-common/base.js", nil)
			if err != nil {
				fmt.Println(err)
			}
			err = p.Push("/learn.json", nil)
			if err != nil {
				fmt.Println(err)
			}
		}
	} else {
		// otherwise just serve the requested resource
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
		// gracefully shutdown after 20 secs
		if err := server.Shutdown(context.Background()); err != nil {
			fmt.Println(err)
		}
	}()

	// TLS is required for HTTP/2
	err := server.ListenAndServeTLS("cert.pem", "key.pem")
	if err != http.ErrServerClosed {
		fmt.Println(err)
	}

	fmt.Println("Server gracefully stopped")
}