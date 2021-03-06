package certificate

import "github.com/chnsz/golangsdk"

const (
	rootPath     = "elbaas"
	resourcePath = "certificate"
)

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(c.ProjectID, rootPath, resourcePath)
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(c.ProjectID, rootPath, resourcePath, id)
}
