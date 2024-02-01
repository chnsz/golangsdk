package eips

import "github.com/chnsz/golangsdk"

const resourcePath = "publicips"

func getURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL(client.ProjectID, "eip", resourcePath, id)
}
