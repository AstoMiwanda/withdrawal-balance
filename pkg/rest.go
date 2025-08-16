package pkg

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	err        error
	httpClient *http.Client
)

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	PATCH  Method = "PATCH"
	DELETE Method = "DELETE"
)

type Request struct {
	Method      Method
	URL         string
	Headers     map[string]string
	QueryParams map[string]string
	Body        []byte
}

func getHttpClient() *http.Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: time.Duration(150) * time.Second,
		}
		httpClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}
	return httpClient
}

func SendRequestRaw(request Request, id string, isDump bool, functionName string) (*http.Response, error) {
	client := getHttpClient()
	if len(request.QueryParams) != 0 {
		request.URL = fmt.Sprintf("%s%s", request.URL, "?")
		params := url.Values{}
		for key, value := range request.QueryParams {
			params.Add(key, value)
		}
		request.URL = fmt.Sprintf("%s%s", request.URL, params.Encode())
	}

	req, err := http.NewRequest(string(request.Method), request.URL, strings.NewReader(string(request.Body)))
	if err != nil {
		logrus.Info(fmt.Sprintf("Error Request  %v: %v", functionName, err.Error()))
		return nil, err
	}

	if len(request.Headers) != 0 {

		for k, v := range request.Headers {
			req.Header[k] = []string{v}
		}
	}

	reqDump, err := httputil.DumpRequest(req, isDump)
	if err != nil {
		logrus.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Info(fmt.Sprintf("Error  Get Response %v: %v", functionName, err.Error()))
		return nil, err
	}

	respDump, err := httputil.DumpResponse(resp, isDump)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info(fmt.Sprintf("URL | %v |ID | %v |Request : %v", request.URL, id, string(reqDump)))
	logrus.Info(fmt.Sprintf("URL | %v |ID | %v |Response : %v", request.URL, id, string(respDump)))
	return resp, nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func SetBasicAuth(username, password string) string {
	return "Basic " + basicAuth(username, password)
}
