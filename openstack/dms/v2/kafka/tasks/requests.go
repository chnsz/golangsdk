package tasks

import (
	"github.com/chnsz/golangsdk"
)

var requestOpts = golangsdk.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
}

func Get(client *golangsdk.ServiceClient, instanceID, taskID string) (*Task, error) {
	var rst golangsdk.Result
	_, rst.Err = client.Get(getURL(client, instanceID, taskID), &rst.Body, &golangsdk.RequestOpts{
		MoreHeaders: requestOpts.MoreHeaders,
	})

	var t Task
	if err := rst.ExtractInto(&t); err != nil {
		return nil, err
	}
	return &t, nil
}
