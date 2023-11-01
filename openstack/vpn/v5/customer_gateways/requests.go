package customer_gateways

import (
	"github.com/chnsz/golangsdk"
	"github.com/chnsz/golangsdk/pagination"
)

type ListOpts struct {
	Limit  int    `q:"limit"`
	Marker string `q:"marker"`
}

func List(c *golangsdk.ServiceClient, opts ListOpts) ([]CustomerGateway, error) {
	url := listURL(c)
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}
	url += query.String()

	pages, err := pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		p := CustomerGatewaysPage{pagination.MarkerPageBase{PageResult: r}}
		p.MarkerPageBase.Owner = p
		return p
	}).AllPages()

	if err != nil {
		return nil, err
	}
	return extractCustomerGateways(pages)
}
