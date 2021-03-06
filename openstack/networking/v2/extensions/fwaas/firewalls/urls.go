package firewalls

import "github.com/chnsz/golangsdk"

const (
	rootPath     = "fw"
	resourcePath = "firewalls"
)

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}
