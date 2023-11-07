package endpoints

import "github.com/chnsz/golangsdk"

type CreateResult struct {
	commonResult
}

type GetResult struct {
	commonResult
}

type UpdateResult struct {
	commonResult
}

type DeleteResult struct {
	golangsdk.ErrResult
}

type commonResult struct {
	golangsdk.Result
}

type CreateOpt struct {
	Name        string        `json:"name"`
	Direction   string        `json:"direction"`
	Region      string        `json:"region"`
	IPAddresses []IPAddresses `json:"ipaddresses"`
}

type IPAddresses struct {
	SubnetID string `json:"subnet_id",required:"true"`
	IP       string `json:"ip,omitempty"`
}

type Endpoint struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Direction         string `json:"direction"`
	Status            string `json:"status"`
	VpcID             string `json:"vpc_id"`
	IPAddressCount    int    `json:"ipaddress_count"`
	ResolverRuleCount int    `json:"resolver_rule_count"`
	CreateTime        string `json:"create_time"`
	UpdateTime        string `json:"update_time"`
}

type UpdateOpts struct {
	Name string `json:"name"`
}

func (r commonResult) Extract() (*Endpoint, error) {
	type response struct {
		Endpoint `json:"endpoint"`
	}
	var res response
	err := r.ExtractInto(&res)
	return &res.Endpoint, err
}
