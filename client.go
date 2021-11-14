package ipwhois

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/go-querystring/query"
)

var (
	// Base URL for the whois API
	FreeURL = "https://ipwhois.app/json/"
	ProURL  = "https://ipwhois.pro/json/"
)

// Client is the client for the ipwhois API
type Client struct {
	url        string       // Base URL for API requests. Will default to Free if there is no API Key provided in the client setup
	apiKey     string       // API key for Nomics API. Not requried for the free version.
	HTTPClient *http.Client // HTTP client for API requests.
}

// NewClient returns a new Client for the Free Version of the API.
func NewClient() *Client {
	return &Client{
		url:        FreeURL,
		HTTPClient: &http.Client{},
	}
}

// NewClientPro returns a new Client for the Pro Version of the API.
// You must provide your API key for the Pro Version.
func NewClientPro(apiKey string) *Client {
	return &Client{
		url:        ProURL,
		apiKey:     apiKey,
		HTTPClient: &http.Client{},
	}
}

func (c *Client) unmarshal(body *string, s interface{}) error {
	err := json.Unmarshal([]byte(*body), &s)
	if err != nil {
		return err
	}
	return nil
}

// get makes a GET request to the API. It returns the response body and error.
func (c *Client) get(address *string, queryMap interface{}) (string, error) {
	url := fmt.Sprintf("%s%s", c.url, *address)

	queryParams, err := query.Values(queryMap)
	if err != nil {
		return "", err
	}

	if c.apiKey != "" {
		url = fmt.Sprintf("%s?key=%s&%s", url, c.apiKey, queryParams.Encode())
	} else {
		url = fmt.Sprintf("%s?%s", url, queryParams.Encode())
	}

	fmt.Printf("URL: %s\n", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}

	defer resp.Body.Close()

	if resp.Header.Get("Content-Type") != "application/json; charset=utf-8" {
		return "", fmt.Errorf("unexpected content type: %s", resp.Header.Get("Content-Type"))
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	bodyString := string(bodyBytes)

	if resp.StatusCode == http.StatusOK {
		return bodyString, nil
	} else {
		log.Println(resp.Status)
		log.Println(bodyString)
		return "", fmt.Errorf("%s", resp.Status)
	}
}
