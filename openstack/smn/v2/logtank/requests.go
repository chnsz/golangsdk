package logtank

import (
	"github.com/chnsz/golangsdk"
)

var RequestOpts = golangsdk.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json;charset=UTF-8"},
}

// Ops is a struct that contains all the parameters.
type Ops struct {
	LogGroupID string `json:"log_group_id" required:"true"`

	LogStreamID string `json:"log_stream_id" required:"true"`
}

// Create a logtank with given parameters.
func Create(client *golangsdk.ServiceClient, topicUrn string, ops Ops) (r CreateResult) {
	b, err := golangsdk.BuildRequestBody(ops, "")
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(createURL(client, topicUrn), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{201},
		MoreHeaders: RequestOpts.MoreHeaders,
	})

	return
}

// List all the logtanks under the topicUrn
func List(client *golangsdk.ServiceClient, topicUrn string) (r ListResult) {
	_, r.Err = client.Get(getURL(client, topicUrn), &r.Body, &golangsdk.RequestOpts{
		MoreHeaders: RequestOpts.MoreHeaders,
	})
	return
}

// Update a logtank with given parameters.
func Update(client *golangsdk.ServiceClient, topicUrn string, logTankID string, ops Ops) (r UpdateResult) {
	b, err := golangsdk.BuildRequestBody(ops, "")
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(updateURL(client, topicUrn, logTankID), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders,
	})
	return
}

// Delete a logtank by topicUrn and LogTankID.
func Delete(client *golangsdk.ServiceClient, topicUrn string, LogTankID string) (r DeleteResult) {
	_, r.Err = client.Delete(DeleteURL(client, topicUrn, LogTankID), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders,
	})
	return
}
