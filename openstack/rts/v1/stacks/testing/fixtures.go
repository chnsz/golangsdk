package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/chnsz/golangsdk"
	"github.com/chnsz/golangsdk/openstack/rts/v1/stacks"
	th "github.com/chnsz/golangsdk/testhelper"
	fake "github.com/chnsz/golangsdk/testhelper/client"
)

// CreateExpected represents the expected object from a Create request.
var CreateExpected = &stacks.CreatedStack{
	ID: "16ef0584-4458-41eb-87c8-0dc8d5f66c87",
	Links: []golangsdk.Link{
		{
			Href: "http://168.28.170.117:8004/v1/98606384f58drad0bhdb7d02779549ac/stacks/stackcreated/16ef0584-4458-41eb-87c8-0dc8d5f66c87",
			Rel:  "self",
		},
	},
}

// CreateOutput represents the response body from a Create request.
const CreateOutput = `
{
  "stack": {
    "id": "16ef0584-4458-41eb-87c8-0dc8d5f66c87",
    "links": [
    {
      "href": "http://168.28.170.117:8004/v1/98606384f58drad0bhdb7d02779549ac/stacks/stackcreated/16ef0584-4458-41eb-87c8-0dc8d5f66c87",
      "rel": "self"
    }
    ]
  }
}`

// HandleCreateSuccessfully creates an HTTP handler at `/stacks` on the test handler mux
// that responds with a `Create` response.
func HandleCreateSuccessfully(t *testing.T, output string) {
	th.Mux.HandleFunc("/stacks", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, output)
	})
}

// ListExpected represents the expected object from a List request.
var ListExpected = []stacks.ListedStack{
	{
		Description: "Simple template to test heat commands",
		Links: []golangsdk.Link{
			{
				Href: "http://166.76.160.117:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/postman_stack/16ef0584-4458-41eb-87c8-0dc8d5f66c87",
				Rel:  "self",
			},
		},
		StatusReason: "Stack CREATE completed successfully",
		Name:         "postman_stack",
		CreationTime: time.Date(2015, 2, 3, 20, 7, 39, 0, time.UTC),
		Status:       "CREATE_COMPLETE",
		ID:           "16ef0584-4458-41eb-87c8-0dc8d5f66c87",
	},
	{
		Description: "Simple template to test heat commands",
		Links: []golangsdk.Link{
			{
				Href: "http://166.76.160.117:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/golangsdk-test-stack-2/db6977b2-27aa-4775-9ae7-6213212d4ada",
				Rel:  "self",
			},
		},
		StatusReason: "Stack successfully updated",
		Name:         "golangsdk-test-stack-2",
		CreationTime: time.Date(2014, 12, 11, 17, 39, 16, 0, time.UTC),
		UpdatedTime:  time.Date(2014, 12, 11, 17, 40, 37, 0, time.UTC),
		Status:       "UPDATE_COMPLETE",
		ID:           "db6977b2-27aa-4775-9ae7-6213212d4ada",
	},
}

// FullListOutput represents the response body from a List request without a marker.
const FullListOutput = `
{
  "stacks": [
  {
    "description": "Simple template to test heat commands",
    "links": [
    {
      "href": "http://166.76.160.117:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/postman_stack/16ef0584-4458-41eb-87c8-0dc8d5f66c87",
      "rel": "self"
    }
    ],
    "stack_status_reason": "Stack CREATE completed successfully",
    "stack_name": "postman_stack",
    "creation_time": "2015-02-03T20:07:39",
    "updated_time": null,
    "stack_status": "CREATE_COMPLETE",
    "id": "16ef0584-4458-41eb-87c8-0dc8d5f66c87",
	"tags": ["rackspace", "atx"]
  },
  {
    "description": "Simple template to test heat commands",
    "links": [
    {
      "href": "http://166.76.160.117:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/golangsdk-test-stack-2/db6977b2-27aa-4775-9ae7-6213212d4ada",
      "rel": "self"
    }
    ],
    "stack_status_reason": "Stack successfully updated",
    "stack_name": "golangsdk-test-stack-2",
    "creation_time": "2014-12-11T17:39:16",
    "updated_time": "2014-12-11T17:40:37",
    "stack_status": "UPDATE_COMPLETE",
    "id": "db6977b2-27aa-4775-9ae7-6213212d4ada",
	"tags": ["sfo", "satx"]
  }
  ]
}
`

// HandleListSuccessfully creates an HTTP handler at `/stacks` on the test handler mux
// that responds with a `List` response.
func HandleListSuccessfully(t *testing.T, output string) {
	th.Mux.HandleFunc("/stacks", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Set("Content-Type", "application/json")
		r.ParseForm()
		marker := r.Form.Get("marker")
		switch marker {
		case "":
			fmt.Fprintf(w, output)
		case "db6977b2-27aa-4775-9ae7-6213212d4ada":
			fmt.Fprintf(w, `[]`)
		default:
			t.Fatalf("Unexpected marker: [%s]", marker)
		}
	})
}

// GetExpected represents the expected object from a Get request.
var GetExpected = &stacks.RetrievedStack{
	DisableRollback: true,
	Description:     "Simple template to test heat commands",
	Parameters: map[string]string{
		"flavor":         "m1.tiny",
		"OS::stack_name": "postman_stack",
		"OS::stack_id":   "16ef0584-4458-41eb-87c8-0dc8d5f66c87",
	},
	StatusReason: "Stack CREATE completed successfully",
	Name:         "postman_stack",
	Outputs:      []*stacks.Output{},
	CreationTime: time.Date(2015, 2, 3, 20, 7, 39, 0, time.UTC),
	Links: []golangsdk.Link{
		{
			Href: "http://166.76.160.117:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/postman_stack/16ef0584-4458-41eb-87c8-0dc8d5f66c87",
			Rel:  "self",
		},
	},
	Capabilities:        []interface{}{},
	NotificationTopics:  []interface{}{},
	Status:              "CREATE_COMPLETE",
	ID:                  "16ef0584-4458-41eb-87c8-0dc8d5f66c87",
	TemplateDescription: "Simple template to test heat commands",
	Tags:                []string{"rackspace", "atx"},
}

// GetOutput represents the response body from a Get request.
const GetOutput = `
{
  "stack": {
    "disable_rollback": true,
    "description": "Simple template to test heat commands",
    "parameters": {
      "flavor": "m1.tiny",
      "OS::stack_name": "postman_stack",
      "OS::stack_id": "16ef0584-4458-41eb-87c8-0dc8d5f66c87"
    },
    "stack_status_reason": "Stack CREATE completed successfully",
    "stack_name": "postman_stack",
    "outputs": [],
    "creation_time": "2015-02-03T20:07:39",
    "links": [
    {
      "href": "http://166.76.160.117:8004/v1/98606384f58d4ad0b3db7d0d779549ac/stacks/postman_stack/16ef0584-4458-41eb-87c8-0dc8d5f66c87",
      "rel": "self"
    }
    ],
    "capabilities": [],
    "notification_topics": [],
    "timeout_mins": null,
    "stack_status": "CREATE_COMPLETE",
    "updated_time": null,
    "id": "16ef0584-4458-41eb-87c8-0dc8d5f66c87",
    "template_description": "Simple template to test heat commands",
	"tags": ["rackspace", "atx"]
  }
}
`

// HandleGetSuccessfully creates an HTTP handler at `/stacks/postman_stack/16ef0584-4458-41eb-87c8-0dc8d5f66c87`
// on the test handler mux that responds with a `Get` response.
func HandleGetSuccessfully(t *testing.T, output string) {
	th.Mux.HandleFunc("/stacks/postman_stack", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, output)
	})
}

// HandleUpdateSuccessfully creates an HTTP handler at `/stacks/postman_stack/16ef0584-4458-41eb-87c8-0dc8d5f66c87`
// on the test handler mux that responds with an `Update` response.
func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/stacks/golangsdk-test-stack-2/db6977b2-27aa-4775-9ae7-6213212d4ada", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
	})
}

// HandleDeleteSuccessfully creates an HTTP handler at `/stacks/postman_stack/16ef0584-4458-41eb-87c8-0dc8d5f66c87`
// on the test handler mux that responds with a `Delete` response.
func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/stacks/golangsdk-test-stack-2/db6977b2-27aa-4775-9ae7-6213212d4ada", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})
}
