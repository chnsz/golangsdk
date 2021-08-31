package alarmreminding

import (
	"github.com/chnsz/golangsdk"
)

func WarnAlert(client *golangsdk.ServiceClient) (r WarnAlertResult) {
	url := WarnAlertURL(client)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
