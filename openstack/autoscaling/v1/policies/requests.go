package policies

import (
	"github.com/chnsz/golangsdk"
)

// CreateOptsBuilder is an interface by which can serialize the create parameters
type CreateOptsBuilder interface {
	ToPolicyCreateMap() (map[string]interface{}, error)
}

// CreateOpts is a struct which will be used to create a policy
type CreateOpts struct {
	Name           string             `json:"scaling_policy_name" required:"true"`
	ID             string             `json:"scaling_group_id" required:"true"`
	Type           string             `json:"scaling_policy_type" required:"true"`
	AlarmID        string             `json:"alarm_id,omitempty"`
	SchedulePolicy SchedulePolicyOpts `json:"scheduled_policy,omitempty"`
	Action         ActionOpts         `json:"scaling_policy_action,omitempty"`
	CoolDownTime   int                `json:"cool_down_time,omitempty"`
}

type SchedulePolicyOpts struct {
	LaunchTime      string `json:"launch_time" required:"true"`
	RecurrenceType  string `json:"recurrence_type,omitempty"`
	RecurrenceValue string `json:"recurrence_value,omitempty"`
	StartTime       string `json:"start_time,omitempty"`
	EndTime         string `json:"end_time,omitempty"`
}

type ActionOpts struct {
	Operation          string `json:"operation,omitempty"`
	InstanceNum        int    `json:"instance_number,omitempty"`
	InstancePercentage int    `json:"instance_percentage,omitempty"`
}

func (opts CreateOpts) ToPolicyCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create is a method which can be able to access to create the policy of autoscaling
// service.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToPolicyCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(createURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// UpdateOptsBuilder is an interface which can build the map paramter of update function
type UpdateOptsBuilder interface {
	ToPolicyUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is a struct which represents the parameters of update function
type UpdateOpts struct {
	Name           string             `json:"scaling_policy_name,omitempty"`
	Type           string             `json:"scaling_policy_type,omitempty"`
	AlarmID        string             `json:"alarm_id,omitempty"`
	SchedulePolicy SchedulePolicyOpts `json:"scheduled_policy,omitempty"`
	Action         ActionOpts         `json:"scaling_policy_action,omitempty"`
	CoolDownTime   int                `json:"cool_down_time,omitempty"`
}

func (opts UpdateOpts) ToPolicyUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Update is a method which can be able to update the policy via accessing to the
// autoscaling service with Put method and parameters
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	body, err := opts.ToPolicyUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(updateURL(client, id), body, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete is a method which can be able to access to delete a policy of autoscaling
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, id), nil)
	return
}

// Get is a method which can be able to access to get a policy detailed information
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

// ListOpts is the structure that used to query the policies of scaling group.
type ListOpts struct {
	// The scaling group ID.
	GroupID string `json:"scaling_group_id" required:"true"`
	// The scaling policy name.
	Name string `q:"scaling_policy_name"`
	// The scaling policy type.
	Type string `q:"scaling_policy_type"`
	// The scaling policy ID.
	PolicyID string `q:"scaling_policy_id"`
	// Start number value. The value must be a positive integer.
	StartNumber int `q:"start_number"`
	// Number of records displayed per page.
	// The value must be a positive integer.
	Limit int `q:"limit"`
}

// List is a method used to query the policies of scaling group with given parameters.
// Due to API limitations, there can be a maximum of 10 policies, so pagination is not considered here,
// simply call the API instead.
func List(client *golangsdk.ServiceClient, opts ListOpts) (r ListResult) {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		r.Err = err
		return
	}

	url := listURL(client, opts.GroupID) + query.String()
	_, r.Err = client.Get(url, &r.Body, nil)
	return
}
