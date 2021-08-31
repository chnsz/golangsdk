package common

import (
	"github.com/chnsz/golangsdk"
	"github.com/chnsz/golangsdk/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient() *golangsdk.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint
	return sc
}
