package ruletypes

import "github.com/chnsz/golangsdk"

func listRuleTypesURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("qos", "rule-types")
}
