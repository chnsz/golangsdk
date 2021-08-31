package testing

import (
	"testing"

	"github.com/chnsz/golangsdk/openstack/mrs/v2/clusters"
	th "github.com/chnsz/golangsdk/testhelper"
	"github.com/chnsz/golangsdk/testhelper/client"
)

func TestCreateV2Cluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2ClusterCreate(t)

	actual, err := clusters.Create(client.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}
