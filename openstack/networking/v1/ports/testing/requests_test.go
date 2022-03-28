package testing

import (
	"testing"

	client "github.com/chnsz/golangsdk/openstack/networking/v1/common"
	"github.com/chnsz/golangsdk/openstack/networking/v1/ports"
	th "github.com/chnsz/golangsdk/testhelper"
)

func TestCreateV1NetworkPort(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV1NetworkPortCreate(t)

	createOpts := ports.CreateOpts{
		DeviceOwner: "neutron:VIP_PORT",
		FixedIps: []ports.FixedIp{
			{
				IpAddress: "192.168.0.162",
				SubnetId:  "885cb8c3-0cbe-406d-83d6-fc98856fcf26",
			},
		},
		NetworkId: "e4cb3b49-78a0-479b-b37d-bd99b3ec0d8a",
	}

	actual, err := ports.Create(client.ServiceClient(), createOpts)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &expectedCreateResponseData, actual)
}

func TestGetV1NetworkPort(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV1NetworkPortGet(t)

	actual, err := ports.Get(client.ServiceClient(), "05547c10-e318-4067-9db2-01f5dc30be38")
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &expectedCreateResponseData, actual)
}

func TestUpdateV1NetworkPort(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV1NetworkPortUpdate(t)

	updateOpts := ports.UpdateOpts{
		AllowedAddressPairs: []ports.AddressPair{
			{
				IpAddress: "192.168.0.25",
			},
		},
	}

	actual, err := ports.Update(client.ServiceClient(), "05547c10-e318-4067-9db2-01f5dc30be38", updateOpts)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &expectedUpdateResponseData, actual)
}

func TestDeleteV1NetworkPort(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV1NetworkPortDelete(t)

	err := ports.Delete(client.ServiceClient(), "05547c10-e318-4067-9db2-01f5dc30be38").ExtractErr()
	th.AssertNoErr(t, err)
}

func TestListV1NetworkPort(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV1NetworkPortList(t)

	listOpts := ports.ListOpts{
		ID: "05547c10-e318-4067-9db2-01f5dc30be38",
	}
	allPages, err := ports.List(client.ServiceClient(), listOpts).AllPages()
	th.AssertNoErr(t, err)

	actual, err := ports.ExtractPorts(allPages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedListResponseData, actual)
}
