package ipapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
)

// IPAPI represents a response from ip-api.com .
type IPAPI struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	ZIP         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	Query       string  `json:"query"`
}

// Get returns information about a given IP address from ip-api.com.
// If the specified IP address is nil, ip-api.com will use the IP address that is seen on their end.
// if the proxy is nil, no proxy will be used for the request.
func Get(ip net.IP, proxy *url.URL) (*IPAPI, error) {
	u := "http://ip-api.com/json"
	if ip != nil {
		u += "/" + ip.String()
	}
	var client *http.Client
	if proxy != nil {
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: func(*http.Request) (*url.URL, error) {
					return proxy, nil
				},
			},
		}
	} else {
		client = http.DefaultClient
	}
	resp, err := client.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ipapi IPAPI
	if err := json.Unmarshal(body, &ipapi); err != nil {
		return nil, err
	}
	if ipapi.Status != "success" {
		return nil, fmt.Errorf("ip-api query failed: %s", ipapi.Status)
	}
	return &ipapi, nil
}
