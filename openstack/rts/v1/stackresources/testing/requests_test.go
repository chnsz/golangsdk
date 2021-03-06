package testing

import (
	"testing"

	"github.com/chnsz/golangsdk/openstack/rts/v1/stackresources"
	th "github.com/chnsz/golangsdk/testhelper"
	fake "github.com/chnsz/golangsdk/testhelper/client"
)

func TestListResources(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t, ListOutput)

	//count := 0
	actual, err := stackresources.List(fake.ServiceClient(), "hello_world", stackresources.ListOpts{})
	if err != nil {
		t.Errorf("Failed to extract resources: %v", err)
	}
	th.AssertDeepEquals(t, ListExpected, actual)
}
