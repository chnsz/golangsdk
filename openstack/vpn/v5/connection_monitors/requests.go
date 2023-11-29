package connection_monitors

import "github.com/chnsz/golangsdk"

type ListOpts struct {
	ID              string `json:"id"`
	Status          string `json:"status"`
	VpnConnectionId string `json:"vpn_connection_id"`
	Type            string `json:"type"`
	SourceIp        string `json:"source_ip"`
	DestinationIp   string `json:"destination_ip"`
	ProtoType       string `json:"proto_type"`
}

var requestOpts = golangsdk.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
}

func List(c *golangsdk.ServiceClient, opts ListOpts) (*ListResp, error) {
	url := listURL(c)
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}
	url += query.String()

	var r ListResp
	_, err = c.Get(url, &r, &golangsdk.RequestOpts{
		MoreHeaders: requestOpts.MoreHeaders,
	})
	return &r, err
}
