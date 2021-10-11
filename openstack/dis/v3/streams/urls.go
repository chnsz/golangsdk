package streams

import "github.com/chnsz/golangsdk"

const (
	resourcePath = "streams"
)

// updateURL PUT /v3/{project_id}/streams/{stream_name}
func UpdateURL(c *golangsdk.ServiceClient, streamName string) string {
	return c.ServiceURL(resourcePath, streamName)
}
