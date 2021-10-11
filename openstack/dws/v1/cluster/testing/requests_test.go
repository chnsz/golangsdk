package testing

import (
	"testing"

	"github.com/chnsz/golangsdk/openstack/dws/v1/cluster"
	th "github.com/chnsz/golangsdk/testhelper"
	"github.com/chnsz/golangsdk/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleList(t)
	id := "7d85f602-a948-4a30-afd4-e84f47471c15"
	listResult, err := cluster.List(client.ServiceClient(), cluster.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, listResult.Clusters[0].ID, id)
	th.AssertEquals(t, listResult.Clusters[0].PublicIp.PublicBindType, cluster.PublicBindTypeAuto)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGet(t)
	id := "7d85f602-a948-4a30-afd4-e84f47471c15"

	clusterDetail, err := cluster.Get(client.ServiceClient(), id)

	th.AssertNoErr(t, err)
	th.AssertEquals(t, clusterDetail.ID, id)
	th.AssertEquals(t, clusterDetail.ResizeInfo.OriginNodeNum, 3)
}
