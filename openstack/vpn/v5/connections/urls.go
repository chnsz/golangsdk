package connections

import (
	"github.com/chnsz/golangsdk"
)

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("vpn-connection")
}
