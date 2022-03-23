package flavors

import (
	"fmt"

	"github.com/chnsz/golangsdk"
)

var RequestOpts golangsdk.RequestOpts = golangsdk.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
}

//list the flavors informations about a specified id of database
func List(client *golangsdk.ServiceClient, dataStoreID string, region string) (r ListResult) {
	url := listURL(client)
	url += fmt.Sprintf("?dbId=%s&region=%s", dataStoreID, region)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}
