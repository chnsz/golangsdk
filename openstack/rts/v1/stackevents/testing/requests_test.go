package testing

import (
	"testing"

	"github.com/chnsz/golangsdk/openstack/rts/v1/stackevents"
	"github.com/chnsz/golangsdk/pagination"
	th "github.com/chnsz/golangsdk/testhelper"
	fake "github.com/chnsz/golangsdk/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t, ListOutput)

	count := 0
	err := stackevents.List(fake.ServiceClient(), "hello_world", "49181cd6-169a-4130-9455-31185bbfc5bf", nil).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := stackevents.ExtractEvents(page)
		th.AssertNoErr(t, err)

		th.CheckDeepEquals(t, ListExpected, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, count, 1)
}
