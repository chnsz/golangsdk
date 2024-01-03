package cloudvolumes

import (
	"github.com/chnsz/golangsdk"
)

type QoSUpdateOpts struct {
	Iops       int `json:"iops" required:"true"`
	Throughput int `json:"throughput,omitempty"`
}

func (opts QoSUpdateOpts) ToVolumeUpdateQoSMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

type UpdateQoSOptsBuilder interface {
	ToVolumeUpdateQoSMap() (map[string]interface{}, error)
}

func UpdateQoS(client *golangsdk.ServiceClient, id string, opts UpdateQoSOptsBuilder) (r JobResult) {
	b, err := opts.ToVolumeUpdateQoSMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(qoSURL(client, id), b, &r.Body, nil)
	return
}
