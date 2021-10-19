package testing

import (
	"testing"

	"github.com/chnsz/golangsdk/openstack/dli/v1/databases"
	th "github.com/chnsz/golangsdk/testhelper"
	"github.com/chnsz/golangsdk/testhelper/client"
)

func TestCreateV1Database(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV1DatabaseCreate(t)

	actual, err := databases.Create(client.ServiceClient(), createOpts)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestListV1Database(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV1DatabaseList(t)

	actual, err := databases.List(client.ServiceClient(), databases.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedListResponseData, actual)
}

func TestUpdateV1DatabaseOwner(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV1DatabaseOwnerUpdate(t)

	actual, err := databases.UpdateDBOwner(client.ServiceClient(), "terraform_test", updateDBOwnerOpts)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestDeleteV1Database(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV1DatabaseDelete(t)

	err := databases.Delete(client.ServiceClient(), "terraform_test").ExtractErr()
	th.AssertNoErr(t, err)
}
