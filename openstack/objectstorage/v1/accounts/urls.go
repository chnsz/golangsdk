package accounts

import "github.com/chnsz/golangsdk"

func getURL(c *golangsdk.ServiceClient) string {
	return c.Endpoint
}

func updateURL(c *golangsdk.ServiceClient) string {
	return getURL(c)
}
