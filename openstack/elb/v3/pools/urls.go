package pools

import "github.com/chnsz/golangsdk"

const (
	rootPath                = "elb"
	resourcePath            = "pools"
	masterSlaveResourcePath = "master-slave-pools"
	memberPath              = "members"
)

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func masterSlaveRootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(rootPath, masterSlaveResourcePath)
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}

func masterSlaveResourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, masterSlaveResourcePath, id)
}

func memberRootURL(c *golangsdk.ServiceClient, poolId string) string {
	return c.ServiceURL(rootPath, resourcePath, poolId, memberPath)
}

func memberResourceURL(c *golangsdk.ServiceClient, poolID string, memeberID string) string {
	return c.ServiceURL(rootPath, resourcePath, poolID, memberPath, memeberID)
}
