package endpoints

import "github.com/chnsz/golangsdk"

func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("endpoints")
}

func endpointURL(client *golangsdk.ServiceClient, endpointID string) string {
	return client.ServiceURL("endpoints", endpointID)
}
