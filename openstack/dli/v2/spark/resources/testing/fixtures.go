package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/chnsz/golangsdk/openstack/dli/v2/spark/resources"
	th "github.com/chnsz/golangsdk/testhelper"
	"github.com/chnsz/golangsdk/testhelper/client"
)

const (
	expectedCreateGroupAndUploadResponse = `
{
	"create_time": 1634873850329,
	"details": [
	 	{
			"create_time": 1634873850329,
			"resource_name": "simple_pyspark_test_DLF_refresh.py",
			"resource_type": "pyFile",
			"status": "READY",
			"underlying_name": "simple_pyspark_test_DLF_refresh.py",
			"update_time": 1634873850329
		}
	],
	"group_name": "terraform-test",
	"is_async": false,
	"owner": "terraform",
	"resources": [
		"simple_pyspark_test_DLF_refresh.py"
	],
	"status": "READY",
	"update_time": 1634873850329
}`

	expectedUploadResponse = `
{
	"create_time": 1634873850329,
	"details": [
	 	{
			"create_time": 1634804544123,
			"resource_name": "simple_pyspark_test_DLF_refresh.py",
			"resource_type": "pyFile",
			"status": "READY",
			"underlying_name": "simple_pyspark_test_DLF_refresh.py",
			"update_time": 1634804544123
		}
	],
	"group_name": "terraform-test",
	"is_async": false,
	"owner": "terraform",
	"resources": [
		"simple_pyspark_test_DLF_refresh.py"
	],
	"status": "READY",
	"update_time": 1634804544123
}`
)

var (
	createGroupAndUploadOpts = resources.CreateGroupAndUploadOpts{
		Group: "terraform-test",
		Kind:  "pyFile",
		Paths: []string{
			"https://terraform-test/dli/packages/simple_pyspark_test_DLF_refresh.py",
		},
	}

	uploadOpts = resources.UploadOpts{
		Group: "terraform-test",
		Paths: []string{
			"https://terraform-test/dli/packages/simple_pyspark_test_DLF_refresh.py",
		},
	}

	expectedCreateGroupAndUploadResponseData = &resources.Group{
		CreateTime: 1634873850329,
		UpdateTime: 1634873850329,
		Details: []resources.Detail{
			{
				CreateTime:     1634873850329,
				UpdateTime:     1634873850329,
				ResourceName:   "simple_pyspark_test_DLF_refresh.py",
				ResourceType:   "pyFile",
				Status:         "READY",
				UnderlyingName: "simple_pyspark_test_DLF_refresh.py",
			},
		},
		GroupName: "terraform-test",
		IsAsync:   false,
		Owner:     "terraform",
		Resources: []string{
			"simple_pyspark_test_DLF_refresh.py",
		},
		Status: "READY",
	}

	expectedUploadResponseData = &resources.Group{
		CreateTime: 1634873850329,
		UpdateTime: 1634804544123,
		Details: []resources.Detail{
			{
				CreateTime:     1634804544123,
				UpdateTime:     1634804544123,
				ResourceName:   "simple_pyspark_test_DLF_refresh.py",
				ResourceType:   "pyFile",
				Status:         "READY",
				UnderlyingName: "simple_pyspark_test_DLF_refresh.py",
			},
		},
		GroupName: "terraform-test",
		IsAsync:   false,
		Owner:     "terraform",
		Resources: []string{
			"simple_pyspark_test_DLF_refresh.py",
		},
		Status: "READY",
	}
)

func handleV2CreateGroupAndUploadPackage(t *testing.T) {
	th.Mux.HandleFunc("/resources", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprint(w, expectedCreateGroupAndUploadResponse)
	})
}

func handleV2UploadPackage(t *testing.T) {
	th.Mux.HandleFunc("/resources/pyfiles", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprint(w, expectedUploadResponse)
	})
}
