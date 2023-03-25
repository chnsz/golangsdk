package acls

import "github.com/chnsz/golangsdk"

// CreateOpts is the structure that used to create a new ACL policy.
type CreateOpts struct {
	// The ACL name.
	Name string `json:"acl_name" required:"true"`
	// The ACL type. The valid values are as follows:
	// + PERMIT
	// + DENY
	Type string `json:"acl_type" required:"true"`
	// The value of the ACL policy.
	// One or more values are supported, separated by commas (,).
	Value string `json:"acl_value" required:"true"`
	// The entity type. The valid values are as follows:
	// + IP
	// + DOMAIN
	// + DOMAIN_ID
	EntityType string `json:"entity_type" required:"true"`
}

var requestOpts = golangsdk.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
}

// Create is a method used to create a private DNAT rule using given parameters.
func Create(c *golangsdk.ServiceClient, instanceId string, opts CreateOpts) (*Policy, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	var r Policy
	_, err = c.Post(rootURL(c, instanceId), b, &r, &golangsdk.RequestOpts{
		MoreHeaders: requestOpts.MoreHeaders,
	})
	return &r, err
}

// Get is a method used to obtain the ACL policy detail by its ID.
func Get(c *golangsdk.ServiceClient, instanceId, policyId string) (*Policy, error) {
	var r Policy
	_, err := c.Get(resourceURL(c, instanceId, policyId), &r, &golangsdk.RequestOpts{
		MoreHeaders: requestOpts.MoreHeaders,
	})
	return &r, err
}

// UpdateOpts is the structure that used to modify an existing ACL policy configuration.
type UpdateOpts struct {
	// The ACL name.
	Name string `json:"acl_name" required:"true"`
	// The ACL type. The valid values are as follows:
	// + PERMIT
	// + DENY
	Type string `json:"acl_type" required:"true"`
	// The value of the ACL policy.
	// One or more values are supported, separated by commas (,).
	Value string `json:"acl_value" required:"true"`
	// The entity type. The valid values are as follows:
	// + IP
	// + DOMAIN
	// + DOMAIN_ID
	// The entity type does not support update.
	EntityType string `json:"entity_type" required:"true"`
}

// Update is a method used to modify the ACL policy configuration using given parameters.
func Update(c *golangsdk.ServiceClient, instanceId, policyId string, opts UpdateOpts) (*Policy, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	var r Policy
	_, err = c.Put(resourceURL(c, instanceId, policyId), b, &r, &golangsdk.RequestOpts{
		MoreHeaders: requestOpts.MoreHeaders,
	})
	return &r, err
}

// Delete is a method to remove the specified ACL policy using its ID and related dedicated instance ID.
func Delete(c *golangsdk.ServiceClient, instanceId, policyId string) error {
	_, err := c.Delete(resourceURL(c, instanceId, policyId), &golangsdk.RequestOpts{
		MoreHeaders: requestOpts.MoreHeaders,
	})
	return err
}
