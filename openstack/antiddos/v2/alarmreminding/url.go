package alarmreminding

import (
	"github.com/chnsz/golangsdk"
)

func WarnAlertURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("warnalert", "alertconfig", "query")
}
