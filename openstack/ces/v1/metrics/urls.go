package metrics

import "github.com/chnsz/golangsdk"

func getMetricsURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("metrics")
}
