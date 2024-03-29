package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/chnsz/golangsdk/openstack/identity/v3/roles"
	th "github.com/chnsz/golangsdk/testhelper"
	fake "github.com/chnsz/golangsdk/testhelper/client"
)

// ListOutput provides a single page of Role results.
const ListOutput = `
{
    "links": {
        "next": null,
        "previous": null,
        "self": "http://example.com/identity/v3/roles"
    },
    "roles": [
        {
            "domain_id": "default",
            "id": "2844b2a08be147a08ef58317d6471f1f",
            "links": {
                "self": "http://example.com/identity/v3/roles/2844b2a08be147a08ef58317d6471f1f"
            },
            "name": "admin-read-only"
        },
        {
            "domain_id": "1789d1",
            "id": "9fe1d3",
            "links": {
                "self": "https://example.com/identity/v3/roles/9fe1d3"
            },
            "name": "support",
            "extra": {
                "description": "read-only support role"
            }
        }
    ]
}
`

const EmptyRespOutput = `
{
	"links": {
		"next": null,
		"previous": null,
		"self": "http://example.com/identity/v3/roles?page=2\u0026per_page=300"
	},
	"roles": [],
	"total_number": 2
}`

// GetOutput provides a Get result.
const GetOutput = `
{
    "role": {
        "domain_id": "1789d1",
        "id": "9fe1d3",
        "links": {
            "self": "https://example.com/identity/v3/roles/9fe1d3"
        },
        "name": "support",
        "extra": {
            "description": "read-only support role"
        }
    }
}
`

// CreateRequest provides the input to a Create request.
const CreateRequest = `
{
    "role": {
        "domain_id": "1789d1",
        "name": "support",
        "description": "read-only support role"
    }
}
`

// UpdateRequest provides the input to as Update request.
const UpdateRequest = `
{
    "role": {
        "description": "admin read-only support role"
    }
}
`

// UpdateOutput provides an update result.
const UpdateOutput = `
{
    "role": {
        "domain_id": "1789d1",
        "id": "9fe1d3",
        "links": {
            "self": "https://example.com/identity/v3/roles/9fe1d3"
        },
        "name": "support",
        "extra": {
            "description": "admin read-only support role"
        }
    }
}
`

const ListAssignmentOutput = `
{
    "role_assignments": [
        {
            "links": {
                "assignment": "http://identity:35357/v3/domains/161718/users/313233/roles/123456"
            },
            "role": {
                "id": "123456"
            },
            "scope": {
                "domain": {
                    "id": "161718"
                }
            },
            "user": {
                "id": "313233"
            }
        },
        {
            "links": {
                "assignment": "http://identity:35357/v3/projects/456789/groups/101112/roles/123456",
                "membership": "http://identity:35357/v3/groups/101112/users/313233"
            },
            "role": {
                "id": "123456"
            },
            "scope": {
                "project": {
                    "id": "456789"
                }
            },
            "user": {
                "id": "313233"
            }
        }
    ],
    "links": {
        "self": "http://identity:35357/v3/role_assignments?effective",
        "previous": null,
        "next": null
    }
}
`

// FirstRole is the first role in the List request.
var FirstRole = roles.Role{
	DomainID: "default",
	ID:       "2844b2a08be147a08ef58317d6471f1f",
	Links: map[string]interface{}{
		"self": "http://example.com/identity/v3/roles/2844b2a08be147a08ef58317d6471f1f",
	},
	Name:  "admin-read-only",
	Extra: map[string]interface{}{},
}

// SecondRole is the second role in the List request.
var SecondRole = roles.Role{
	DomainID: "1789d1",
	ID:       "9fe1d3",
	Links: map[string]interface{}{
		"self": "https://example.com/identity/v3/roles/9fe1d3",
	},
	Name: "support",
	Extra: map[string]interface{}{
		"description": "read-only support role",
	},
}

// SecondRoleUpdated is how SecondRole should look after an Update.
var SecondRoleUpdated = roles.Role{
	DomainID: "1789d1",
	ID:       "9fe1d3",
	Links: map[string]interface{}{
		"self": "https://example.com/identity/v3/roles/9fe1d3",
	},
	Name: "support",
	Extra: map[string]interface{}{
		"description": "admin read-only support role",
	},
}

// ExpectedRolesSlice is the slice of roles expected to be returned from ListOutput.
var ExpectedRolesSlice = []roles.Role{FirstRole, SecondRole}

// HandleListRolesSuccessfully creates an HTTP handler at `/roles` on the
// test handler mux that responds with a list of two roles.
func HandleListRolesSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/roles", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		r.ParseForm()
		pageNum := r.Form.Get("page")
		switch pageNum {
		case "":
			fmt.Fprintf(w, ListOutput)
		case "2":
			fmt.Fprintf(w, EmptyRespOutput)
		default:
			t.Fatalf("/roles invoked with unexpected page=[%s]", pageNum)
		}
	})
}

// HandleGetRoleSuccessfully creates an HTTP handler at `/roles` on the
// test handler mux that responds with a single role.
func HandleGetRoleSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/roles/9fe1d3", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, GetOutput)
	})
}

// HandleCreateRoleSuccessfully creates an HTTP handler at `/roles` on the
// test handler mux that tests role creation.
func HandleCreateRoleSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/roles", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestJSONRequest(t, r, CreateRequest)

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, GetOutput)
	})
}

// HandleUpdateRoleSuccessfully creates an HTTP handler at `/roles` on the
// test handler mux that tests role update.
func HandleUpdateRoleSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/roles/9fe1d3", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestJSONRequest(t, r, UpdateRequest)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, UpdateOutput)
	})
}

// HandleDeleteRoleSuccessfully creates an HTTP handler at `/roles` on the
// test handler mux that tests role deletion.
func HandleDeleteRoleSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/roles/9fe1d3", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

func HandleAssignSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/projects/{project_id}/users/{user_id}/roles/{role_id}", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	th.Mux.HandleFunc("/projects/{project_id}/groups/{group_id}/roles/{role_id}", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	th.Mux.HandleFunc("/domains/{domain_id}/users/{user_id}/roles/{role_id}", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	th.Mux.HandleFunc("/domains/{domain_id}/groups/{group_id}/roles/{role_id}", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})
}

func HandleUnassignSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/projects/{project_id}/users/{user_id}/roles/{role_id}", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	th.Mux.HandleFunc("/projects/{project_id}/groups/{group_id}/roles/{role_id}", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	th.Mux.HandleFunc("/domains/{domain_id}/users/{user_id}/roles/{role_id}", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	th.Mux.HandleFunc("/domains/{domain_id}/groups/{group_id}/roles/{role_id}", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})
}

// FirstRoleAssignment is the first role assignment in the List request.
var FirstRoleAssignment = roles.RoleAssignment{
	Catalog:     "BASE",
	Description: "Tenant Administrator",
	DisplayName: "Tenant Administrator",
	ID:          "699bd62cda304d2cad03fd2fb190b8cf",
	Name:        "te_admin",
	Type:        "AA",
	Policy: roles.Policy{
		Statement: []roles.Statement{},
		Version:   "v1",
	},
}

// SecondRoleAssignemnt is the second role assignemnt in the List request.
var SecondRoleAssignment = roles.RoleAssignment{
	Catalog:     "BASE",
	Description: "Security Administrator",
	DisplayName: "Security Administrator",
	ID:          "699bd62cda304d2cad03fd2fb190b8ce",
	Name:        "secu_admin",
	Type:        "AA",
	Policy: roles.Policy{
		Statement: []roles.Statement{},
		Version:   "v1",
	},
}

// ExpectedRoleAssignmentsSlice is the slice of role assignments expected to be
// returned from ListAssignmentOutput.
var ExpectedRoleAssignmentsSlice = []roles.RoleAssignment{FirstRoleAssignment, SecondRoleAssignment}

// HandleListRoleAssignmentsSuccessfully creates an HTTP handler at `/role_assignments` on the
// test handler mux that responds with a list of two role assignments.
func HandleListRoleAssignmentsSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/domains/{domain_id}/groups/{group_id}/roles", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ListAssignmentOutput)
	})

	th.Mux.HandleFunc("/projects/{project_id}/groups/{group_id}/roles", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ListAssignmentOutput)
	})

	th.Mux.HandleFunc("/projects/{project_id}/users/{user_id}/roles", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ListAssignmentOutput)
	})

	th.Mux.HandleFunc("/domains/{domain_id}/users/{user_id}/roles", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ListAssignmentOutput)
	})
}
