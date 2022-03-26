package testing

import (
	"testing"

	client "github.com/chnsz/golangsdk/openstack/networking/v1/common"
	"github.com/chnsz/golangsdk/openstack/networking/v1/ports"
	th "github.com/chnsz/golangsdk/testhelper"
)

func TestCreateV1NetworkVIP(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV1NetworkVIPCreate(t)

	actual, err := ports.Create(client.ServiceClient(), createOpts)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestGetV1NetworkVIP(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV1NetworkVIPGet(t)

	actual, err := ports.Get(client.ServiceClient(), "05547c10-e318-4067-9db2-01f5dc30be38")
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestUpdateV1NetworkVIP(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV1NetworkVIPUpdate(t)

	actual, err := ports.Update(client.ServiceClient(), "05547c10-e318-4067-9db2-01f5dc30be38", updateOpts)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestDeleteV1NetworkVIP(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV1NetworkVIPDelete(t)

	err := ports.Delete(client.ServiceClient(), "05547c10-e318-4067-9db2-01f5dc30be38").ExtractErr()
	th.AssertNoErr(t, err)
}
