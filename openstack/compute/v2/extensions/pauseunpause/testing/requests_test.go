package testing

import (
	"testing"

	"github.com/chnsz/golangsdk/openstack/compute/v2/extensions/pauseunpause"
	th "github.com/chnsz/golangsdk/testhelper"
	"github.com/chnsz/golangsdk/testhelper/client"
)

const serverID = "{serverId}"

func TestPause(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockPauseServerResponse(t, serverID)

	err := pauseunpause.Pause(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestUnpause(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockUnpauseServerResponse(t, serverID)

	err := pauseunpause.Unpause(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}
