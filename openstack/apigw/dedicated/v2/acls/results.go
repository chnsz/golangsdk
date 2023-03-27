package acls

// Policy is the structure represents the ACL policy details.
type Policy struct {
	// The ACL name.
	Name string `json:"acl_name"`
	// The ACL type. The valid values are as follows:
	// + PERMIT
	// + DENY
	Type string `json:"acl_type"`
	// The value of the ACL policy.
	Value string `json:"acl_value"`
	// The entity type. The valid values are as follows:
	// + IP
	// + DOMAIN
	// + DOMAIN_ID
	EntityType string `json:"entity_type"`
	// The ID of the ACL policy.
	ID string `json:"id"`
	// The latest update time.
	UpdatedAt string `json:"update_time"`
}
