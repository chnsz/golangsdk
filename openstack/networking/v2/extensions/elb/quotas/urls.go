package quotas

import "github.com/chnsz/golangsdk"

const (
	rootPath     = "elbaas"
	resourcePath = "quotas"
)

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(c.ProjectID, rootPath, resourcePath)
}
