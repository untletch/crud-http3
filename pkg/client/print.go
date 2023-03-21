package client

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fatih/color"
)

func PrintResponseRequestHeaders(req *http.Request, resp *http.Response) {
	blue := color.New(color.FgBlue).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	fmt.Printf("%s\n", red("Response Headers"))
	fmt.Printf("%s %s\n", blue(resp.Proto), blue(resp.Status))

	for key, values := range resp.Header {
		fmt.Printf("%s: ", cyan(key))
		for _, value := range values {
			fmt.Printf("%s", yellow(value))
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Printf("%s\n", red("Request Headers"))
	fmt.Printf("%s\n", blue(req.Proto))

	for key, values := range req.Header {
		fmt.Printf("%s: ", cyan(key))
		for _, value := range values {
			fmt.Printf("%s", yellow(value))
		}
		fmt.Println()
	}
	fmt.Println()
}

func PrintResponseBody(respBody io.ReadCloser) {
	_, err := io.Copy(os.Stdout, respBody)
	if err != nil {
		fmt.Println("error copying body to stdout:", err)
		return
	}
	fmt.Println()
}
