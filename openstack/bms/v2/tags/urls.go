package tags

import (
	"github.com/chnsz/golangsdk"
)

func resourceURL(client *golangsdk.ServiceClient, serverId string) string {
	return client.ServiceURL("servers", serverId, "tags")
}
