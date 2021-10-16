package testing

import (
	"testing"

	"github.com/chnsz/golangsdk/openstack/dli/v2/batches"
	th "github.com/chnsz/golangsdk/testhelper"
	"github.com/chnsz/golangsdk/testhelper/client"
)

func TestCreateV2SparkJob(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2SparkJobCreate(t)

	actual, err := batches.Create(client.ServiceClient(), createOpts)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestGetV2SparkJob(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2SparkJobGet(t)

	actual, err := batches.Get(client.ServiceClient(), "6145a791-81a4-4edb-b2d2-ea599caf6550")
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestGetV2SparkJobState(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2SparkJobGetState(t)

	actual, err := batches.GetState(client.ServiceClient(), "6145a791-81a4-4edb-b2d2-ea599caf6550")
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedGetStateResponseData, actual)
}

func TestDeleteV2SparkJob(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2SparkJobDelete(t)

	err := batches.Delete(client.ServiceClient(), "6145a791-81a4-4edb-b2d2-ea599caf6550").ExtractErr()
	th.AssertNoErr(t, err)
}
