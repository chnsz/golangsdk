package schedulerstats

import "github.com/chnsz/golangsdk"

func storagePoolsListURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("scheduler-stats", "get_pools")
}
