package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/chnsz/golangsdk/openstack/dli/v1/databases"
	th "github.com/chnsz/golangsdk/testhelper"
	"github.com/chnsz/golangsdk/testhelper/client"
)

const (
	expectedCreateResponse = `
	{
		"is_success": true
	}`

	expectedListResponse = `
{
	"database_count": 1,
	"databases": [
		{
			"database_name": "terraform_test",
			"description": "Created by terraform",
			"enterprise_project_id": "e9ee3f48-f097-406a-aa74-cfece0af3e31",
			"is_shared": false,
			"owner": "terraform",
			"resource_id": "e766d389-d73f-45e7-af79-29445c36339b",
			"table_number": 0
		}
	],
	"is_success": true
}`
)

var (
	createOpts = databases.CreateOpts{
		Name:                "terraform_test",
		EnterpriseProjectId: "e9ee3f48-f097-406a-aa74-cfece0af3e31",
	}

	expectedCreateResponseData = &databases.RequestResp{
		IsSuccess: true,
	}

	expectedListResponseData = &databases.ListResp{
		DatabaseCount: 1,
		Databases: []databases.Database{
			{
				Name:                "terraform_test",
				Description:         "Created by terraform",
				EnterpriseProjectId: "e9ee3f48-f097-406a-aa74-cfece0af3e31",
				IsShared:            false,
				Owner:               "terraform",
				ResourceId:          "e766d389-d73f-45e7-af79-29445c36339b",
				TableNumber:         0,
			},
		},
		IsSuccess: true,
	}

	updateDBOwnerOpts = databases.UpdateDBOwnerOpts{
		NewOwner: "newOwner",
	}
)

func handleV1DatabaseCreate(t *testing.T) {
	th.Mux.HandleFunc("/databases", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedCreateResponse)
	})
}

func handleV1DatabaseList(t *testing.T) {
	th.Mux.HandleFunc("/databases", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedListResponse)
	})
}

func handleV1DatabaseOwnerUpdate(t *testing.T) {
	th.Mux.HandleFunc("/databases/terraform_test/owner", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedCreateResponse)
	})
}

func handleV1DatabaseDelete(t *testing.T) {
	th.Mux.HandleFunc("/databases/terraform_test", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
}
