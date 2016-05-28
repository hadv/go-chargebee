package chargebee

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	log "github.com/Innovatube/log4go"
)

// defaultHTTPTimeout is the default timeout on the http.Client used by the library.
// This is chosen to be consistent with the other Stripe language libraries and
// to coordinate with other timeouts configured in the Stripe infrastructure.
const defaultHTTPTimeout = 80 * time.Second

// Backend is an interface for making calls against a chargebee service.
// This interface exists to enable mocking for during testing if needed.
type Backend interface {
	Call(method, path, key string, body *url.Values, v interface{}) error
}

// BackendConfiguration is the internal implementation for making HTTP calls to chargebee.
type BackendConfiguration struct {
	URL        string
	HTTPClient *http.Client
}

const (
	CHARSET     string = "UTF-8"
	API_VERSION string = "v2"
)

// Key is the Chargebee API key used globally in the binding.
var Key string

// SiteName is the Chargebee API sitename used globally in the binding.
var SiteName string

var httpClient = &http.Client{Timeout: defaultHTTPTimeout}

func NewBackend(siteName string) Backend {
	apiBaseUrl := fmt.Sprintf("https://%v.chargebee.com/api/%v", siteName, API_VERSION)
	return BackendConfiguration{
		URL:        apiBaseUrl,
		HTTPClient: httpClient,
	}
}

// Call is the Backend.Call implementation for invoking Chargebee APIs.
func (s BackendConfiguration) Call(method, path, key string, form *url.Values, v interface{}) error {
	var body io.Reader
	if form != nil && len(*form) > 0 {
		data := form.Encode()
		if strings.ToUpper(method) == "GET" {
			path += "?" + data
		} else {
			body = bytes.NewBufferString(data)
		}
	}

	req, err := s.NewRequest(method, path, key, "application/x-www-form-urlencoded", body)
	if err != nil {
		return err
	}

	if err := s.Do(req, v); err != nil {
		return err
	}

	return nil
}

// NewRequest is used by Call to generate an http.Request. It handles encoding
// parameters and attaching the appropriate headers.
func (s *BackendConfiguration) NewRequest(method, path, key, contentType string, body io.Reader) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = s.URL + path

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept-Charset", CHARSET)
	req.SetBasicAuth(key, "")
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Accept", "application/json")

	return req, nil
}

// Do is used by Call to execute an API request and parse the response. It uses
// the backend's HTTP client to execute the request and unmarshals the response
// into v. It also handles unmarshaling errors returned by the API.
func (s *BackendConfiguration) Do(req *http.Request, v interface{}) error {
	fmt.Println("Requesting %v %v%v\n", req.Method, req.URL.Host, req.URL.Path)
	res, err := s.HTTPClient.Do(req)

	if err != nil {
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		log.Error("chargebee error: %v", string(resBody))
		return errors.New("An error happen, we cannot fulfill your request at the moment.")
	}

	if v != nil {
		return json.Unmarshal(resBody, v)
	}

	return nil
}
