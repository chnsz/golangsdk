package testing

import (
	"testing"

	"github.com/chnsz/golangsdk/openstack/dis/v2/streams"
	th "github.com/chnsz/golangsdk/testhelper"
	"github.com/chnsz/golangsdk/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleList(t)
	listResult, err := streams.List(client.ServiceClient(), streams.ListStreamsOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, listResult.TotalNumber, 2)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGet(t)
	name := "dis-ML2k"

	resp, err := streams.Get(client.ServiceClient(), name, streams.GetOpts{})

	th.AssertNoErr(t, err)
	th.AssertEquals(t, resp.StreamName, name)
}
