package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"crud-http3/pkg/client"

	"github.com/quic-go/quic-go/http3"
)

func request(req *http.Request, verify, userAgentVar bool) (*http.Response, *http3.RoundTripper) {
	roundTripper := &http3.RoundTripper{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: verify,
		},
	}

	client_req := http.Client{
		Transport: roundTripper,
		Timeout:   30 * time.Second,
	}

	// set request headers
	accept := "text/html,application/xhtml+xml," +
		"application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"
	req.Header.Add("Accept", accept)
	req.Header.Add("Host", req.Host)
	if userAgentVar {
		req.Header.Add("User-Agent", client.UserAgent())
	}

	resp, err := client_req.Do(req)
	if err != nil {
		log.Fatal("client.Do error:", err)
	}

	return resp, roundTripper
}

func main() {
	var req *http.Request

	urlStr := flag.String("url", "", "http3 request 'url'")
	method := flag.String("method", "GET", "http method")
	verbose := flag.Bool("verbose", true, "print request and response headers")
	skipVerify := flag.Bool("skipverify", true, "skip certificate verification")
	userAgentVar := flag.Bool("useragent", false, "use random user agent (default false)")

	flag.Parse()

	if *urlStr == "" {
		fmt.Println("missing required -url parameter (http3 url)")
		return
	}

	switch *method {
	case "GET":
		req = client.Get(*urlStr)
	}

	resp, rt := request(req, *skipVerify, *userAgentVar)

	if *verbose {
		client.PrintResponseRequestHeaders(req, resp)
	}

	client.PrintResponseBody(resp.Body)

	resp.Body.Close()
	rt.Close()
}
