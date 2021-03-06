package volumes

import "github.com/chnsz/golangsdk"

func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("volumes")
}

func listURL(c *golangsdk.ServiceClient) string {
	return createURL(c)
}

func deleteURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id)
}

func getURL(c *golangsdk.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func updateURL(c *golangsdk.ServiceClient, id string) string {
	return deleteURL(c, id)
}
