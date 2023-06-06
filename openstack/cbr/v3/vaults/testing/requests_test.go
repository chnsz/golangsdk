package testing

import (
	"testing"

	"github.com/chnsz/golangsdk/openstack/cbr/v3/vaults"
	th "github.com/chnsz/golangsdk/testhelper"
	"github.com/chnsz/golangsdk/testhelper/client"
)

func TestCreateV3VaultsMarshall(t *testing.T) {
	res, err := createOpts.ToVaultCreateMap()
	th.AssertNoErr(t, err)
	th.AssertJSONEquals(t, expectedCreateRequest, res)
}

func TestCreateV3Vault(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleVaultCreate(t)

	actual, err := vaults.Create(client.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestGetV3Vault(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleVaultGet(t)

	actual, err := vaults.Get(client.ServiceClient(), "ad7627ae-5b0b-492e-b6bd-cd809b745197").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestUpdateV3Vault(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleVaultUpdate(t)

	actual, err := vaults.Update(client.ServiceClient(), "ad7627ae-5b0b-492e-b6bd-cd809b745197", updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateResponseData, actual)
}

func TestDeleteV3Vault(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleVaultDelete(t)

	err := vaults.Delete(client.ServiceClient(), "ad7627ae-5b0b-492e-b6bd-cd809b745197").ExtractErr()
	th.AssertNoErr(t, err)
}

func TestListV3Vault(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleVaultList(t)

	allPages, err := vaults.List(client.ServiceClient(), vaults.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := vaults.ExtractVaults(allPages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedListResponseData, actual)
}

func TestBackupPolicyBindingV3Vault(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleVaultBindBackupPolicy(t)

	actual, err := vaults.BindPolicy(client.ServiceClient(), "ad7627ae-5b0b-492e-b6bd-cd809b745197",
		bindBackupPolicyOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedBackupPolicyBindResponseData, actual)
}

func TestReplicationPolicyBindingV3Vault(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleVaultBindReplicationPolicy(t)

	actual, err := vaults.BindPolicy(client.ServiceClient(), "ad7627ae-5b0b-492e-b6bd-cd809b745197",
		bindReplicationPolicyOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedReplicationPolicyBindResponseData, actual)
}

func TestBackupPolicyUnbindingV3Vault(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleVaultUnbindBackupPolicy(t)

	actual, err := vaults.UnbindPolicy(client.ServiceClient(), "ad7627ae-5b0b-492e-b6bd-cd809b745197",
		unbindPolicyOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedBackupPolicyBindResponseData, actual)
}

func TestReplicationPolicyUnbindingV3Vault(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleVaultUnbindReplicationPolicy(t)

	actual, err := vaults.UnbindPolicy(client.ServiceClient(), "ad7627ae-5b0b-492e-b6bd-cd809b745197",
		unbindPolicyOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedReplicationPolicyBindResponseData, actual)
}

func TestAssociateResourcesV3Vault(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleVaultAssociateResources(t)

	actual, err := vaults.AssociateResources(client.ServiceClient(), "ad7627ae-5b0b-492e-b6bd-cd809b745197",
		associateResourcesOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedAssociateResourcesResponseData, actual)
}

func TestDissociateResourcesV3Vault(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleVaultDissociateResources(t)

	_, err := vaults.DissociateResources(client.ServiceClient(), "ad7627ae-5b0b-492e-b6bd-cd809b745197",
		dissociateResourcesOpts).Extract()
	th.AssertNoErr(t, err)
}
