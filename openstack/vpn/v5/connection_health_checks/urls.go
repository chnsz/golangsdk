package connection_health_checks

import (
	"github.com/chnsz/golangsdk"
)

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("connection-monitors")
}
