package testing

import (
	"testing"

	"github.com/chnsz/golangsdk/openstack/fgs/v2/dependencies"
	"github.com/chnsz/golangsdk/pagination"
	th "github.com/chnsz/golangsdk/testhelper"
	"github.com/chnsz/golangsdk/testhelper/client"
)

func TestCreateV2Dependency(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2DependencyCreate(t)

	actual, err := dependencies.Create(client.ServiceClient(), createOpts)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &expectedGetResponseData, actual)
}

func TestGetV2Dependency(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2DependencyGet(t)

	actual, err := dependencies.Get(client.ServiceClient(), "e6cc2ebe-0bae-4b69-a1d6-8198bc356ff8")
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &expectedGetResponseData, actual)
}

func TestListV2Dependencies(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2DependenciesList(t)

	actual := make([]dependencies.Dependency, 0)
	err := dependencies.List(client.ServiceClient(), listOpts).EachPage(
		func(page pagination.Page) (bool, error) {
			resp, err := dependencies.ExtractDependencies(page)
			th.AssertNoErr(t, err)
			actual = append(actual, resp.Dependencies...)
			return true, nil
		})
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, expectedListResponseData, actual)
}

func TestUpdateV2Dependency(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2DependencyUpdate(t)

	actual, err := dependencies.Update(client.ServiceClient(), "e6cc2ebe-0bae-4b69-a1d6-8198bc356ff8", createOpts)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &expectedGetResponseData, actual)
}

func TestDeleteV2Dependency(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2DependencyDelete(t)

	err := dependencies.Delete(client.ServiceClient(), "e6cc2ebe-0bae-4b69-a1d6-8198bc356ff8").ExtractErr()
	th.AssertNoErr(t, err)
}
