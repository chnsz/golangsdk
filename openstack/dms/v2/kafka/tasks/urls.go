package tasks

import "github.com/chnsz/golangsdk"

func getURL(client *golangsdk.ServiceClient, instanceID, taskID string) string {
	return client.ServiceURL(client.ProjectID, "instances", instanceID, "tasks", taskID)
}
