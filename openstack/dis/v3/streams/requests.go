package streams

import (
	"github.com/chnsz/golangsdk"
)

type UpdateOpts struct {
	StreamName                 string `json:"stream_name" required:"true"`
	DataDuration               *int   `json:"data_duration,omitempty"`
	DataType                   string `json:"data_type,omitempty"`
	DataSchema                 string `json:"data_schema,omitempty"`
	AutoScaleEnabled           *bool  `json:"auto_scale_enabled,omitempty"`
	AutoScaleMinPartitionCount *int   `json:"auto_scale_min_partition_count,omitempty"`
	AutoScaleMaxPartitionCount *int   `json:"auto_scale_max_partition_count,omitempty"`
}

func Update(c *golangsdk.ServiceClient, name string, opts UpdateOpts) (*golangsdk.Result, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	var r golangsdk.Result
	_, err = c.Put(UpdateURL(c, name), b, &r.Body, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
			"region":       c.AKSKAuthOptions.Region,
		},
	})
	return &r, err
}
