package stackevents

import "github.com/chnsz/golangsdk"

func listURL(c *golangsdk.ServiceClient, stackName, stackID string) string {
	return c.ServiceURL("stacks", stackName, stackID, "events")
}
