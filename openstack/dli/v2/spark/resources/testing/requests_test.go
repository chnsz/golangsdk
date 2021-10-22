package testing

import (
	"testing"

	"github.com/chnsz/golangsdk/openstack/dli/v2/spark/resources"
	th "github.com/chnsz/golangsdk/testhelper"
	"github.com/chnsz/golangsdk/testhelper/client"
)

func TestV2CreateGroupAndUploadPackage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2CreateGroupAndUploadPackage(t)

	actual, err := resources.CreateGroupAndUpload(client.ServiceClient(), createGroupAndUploadOpts)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedCreateGroupAndUploadResponseData, actual)
}

func TestV2UploadPackage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleV2UploadPackage(t)

	actual, err := resources.Upload(client.ServiceClient(), "pyfiles", uploadOpts)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedUploadResponseData, actual)
}
