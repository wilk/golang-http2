# HTTP/2 Demo

This is a demo of [HTTP/2 Push](https://http2.github.io/) with Go 1.8 performed in a live coding session during the [release party](https://github.com/golang/go/wiki/Go-1.8-Release-Party).

## Requirements

- [Go 1.8](https://github.com/golang/go/releases/tag/go1.8)
- OpenSSL

## Running the Demo

First of all, generate the certificate for the TLS connection:

```bash
$ ./gen_cert.sh
```

Now, launch the following command:

```bash
$ go run demo.go
```

Then, open your favourite browser and type `https://localhost:3001` in the address bar.

## Serving the Slideshow

Serving the slideshow can be achived with [present](https://godoc.org/golang.org/x/tools/present):

```bash
$ present
```

Then, again open your browser and type `http://127.0.0.1:3999`.
There you go!