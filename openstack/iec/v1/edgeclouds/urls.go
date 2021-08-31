package edgeclouds

import (
	"github.com/chnsz/golangsdk"
)

func GetURL(c *golangsdk.ServiceClient, edgeCloudID string) string {
	return c.ServiceURL("edgeclouds", edgeCloudID)
}

func DeleteURL(c *golangsdk.ServiceClient, edgeCloudID string) string {
	return c.ServiceURL("edgeclouds", edgeCloudID)
}
