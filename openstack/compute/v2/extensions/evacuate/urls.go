package evacuate

import (
	"github.com/chnsz/golangsdk"
)

func actionURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}
