package products

import (
	"github.com/chnsz/golangsdk"
)

// Get products
func Get(client *golangsdk.ServiceClient, engine string) (r GetResult) {
	url := getURL(client)
	if engine != "" {
		url = url + "?engine=" + engine
	}
	_, r.Err = client.Get(url, &r.Body, nil)
	return
}
