package tags

import "github.com/chnsz/golangsdk"

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(id, "tags")
}
