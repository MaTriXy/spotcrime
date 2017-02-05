package spotcrime

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.95 Safari/537.36"
	baseURL   = "http://api.spotcrime.com/crimes.json"

	// in case no radius is provided (only lat,lon), use this as the default radius
	defaultRadius = 0.01
)

// Client is the Spotcrime client. It contains all the different resources available.
type Client struct {
	key   string
	Debug bool
}

// New creates a new Spotcrime client with the appropriate secret key
func New(key string) (*Client, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("must provide key")
	}

	return &Client{
		key: key,
	}, nil
}

// Request contains information to sent to the api endpoint
type Request struct {
	Lat    float64
	Lon    float64
	Radius float64
	Proxy  string
}

// Response contains Results from the API request
type Response struct {
	Results Results `json:"crimes"`
}

// Results is a slice of Result
type Results []Result

// Result directly corresponds to the JSON returned by the API
type Result struct {
	CDID int     `json:"cdid"`
	Type string  `json:"type"`
	Date string  `json:"date"`
	Link string  `json:"link"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}

// GetCrimes fetches crimes from Spotcrime API
func (c *Client) GetCrimes(r *Request) (*Response, error) {
	if r.Lat == 0 || r.Lon == 0 {
		return nil, fmt.Errorf("must provide lat and lon")
	}

	if r.Radius == 0 {
		r.Radius = defaultRadius
	}

	sURL := fmt.Sprintf("%s?key=%s&lat=%f&lon=%f&radius=%f", baseURL, c.key, r.Lat, r.Lon, r.Radius)

	if c.Debug {
		log.Printf("sURL: %s", sURL)
	}

	client := &http.Client{}

	if len(r.Proxy) > 0 {
		_, err := url.Parse(r.Proxy)
		if err == nil {
			os.Setenv("HTTP_PROXY", r.Proxy)
		}
	}

	req, err := http.NewRequest("GET", sURL, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	var response *Response
	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
