package logtank

import (
	"github.com/chnsz/golangsdk"
)

func createURL(c *golangsdk.ServiceClient, topicUrn string) string {
	return c.ServiceURL("topics", topicUrn, "logtanks")
}

func getURL(c *golangsdk.ServiceClient, topicUrn string) string {
	return c.ServiceURL("topics", topicUrn, "logtanks")
}

func updateURL(c *golangsdk.ServiceClient, topicUrn string, logTankId string) string {
	return c.ServiceURL("topics", topicUrn, "logtanks", logTankId)
}

func DeleteURL(c *golangsdk.ServiceClient, topicUrn string, logTankId string) string {
	return c.ServiceURL("topics", topicUrn, "logtanks", logTankId)
}
