package testing

import (
	"testing"

	"github.com/chnsz/golangsdk/openstack/antiddos/v2/alarmreminding"
	th "github.com/chnsz/golangsdk/testhelper"
	"github.com/chnsz/golangsdk/testhelper/client"
)

func TestQueryTraffic(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleWarnAlertSuccessfully(t)

	actual, err := alarmreminding.GetWarnAlert(client.ServiceClient()).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &WarnAlertResponse, actual)
}
