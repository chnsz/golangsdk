package testing

import (
	"github.com/chnsz/golangsdk"
	"github.com/chnsz/golangsdk/testhelper"
)

func createClient() *golangsdk.ServiceClient {
	return &golangsdk.ServiceClient{
		ProviderClient: &golangsdk.ProviderClient{TokenID: "abc123"},
		Endpoint:       testhelper.Endpoint(),
	}
}
