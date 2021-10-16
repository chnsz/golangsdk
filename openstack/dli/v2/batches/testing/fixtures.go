package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/chnsz/golangsdk/openstack/dli/v2/batches"
	th "github.com/chnsz/golangsdk/testhelper"
	"github.com/chnsz/golangsdk/testhelper/client"
)

const (
	expectedCreateResponse = `
{
	"cluster_name": "driver_behavior",
	"create_time": 1634285547245,
	"id": "6145a791-81a4-4edb-b2d2-ea599caf6550",
	"name": "terraform_spark_job",
	"owner": "terraform",
	"queue": "terraform_test",
	"sc_type": "CUSTOMIZED",
	"state": "starting",
	"update_time": 1634285547245
}`

	expectedGetStateResponse = `
{
	"id": "6145a791-81a4-4edb-b2d2-ea599caf6550",
	"state": "starting"
}`
)

var (
	createOpts = batches.CreateOpts{
		ClassName: "driver_behavior",
		Queue:     "terraform_test",
		Name:      "terraform_spark_job",
		Configurations: map[string]interface{}{
			"spark.dli.metaAccess.enable": "true",
		},
		File: "driver_package/driver_behavior.jar",
		Jars: []string{"jar_package/jackson-core-2.13.0-javadoc.jar"},
		Groups: []batches.Group{
			{
				Name: "driver_package",
			},
		},
		MaxRetryTimes:  20,
		Specification:  "A",
		DriverMemory:   "5G",
		DriverCores:    2,
		ExecutorMemory: "8G",
		ExecutorCores:  2,
		NumExecutors:   12,
	}

	expectedCreateResponseData = &batches.CreateResp{
		ClusterName: "driver_behavior",
		ID:          "6145a791-81a4-4edb-b2d2-ea599caf6550",
		Name:        "terraform_spark_job",
		Owner:       "terraform",
		Queue:       "terraform_test",
		ScType:      "CUSTOMIZED",
		State:       "starting",
		CreateTime:  1634285547245,
		UpdateTime:  1634285547245,
	}

	expectedGetStateResponseData = &batches.StateResp{
		ID:    "6145a791-81a4-4edb-b2d2-ea599caf6550",
		State: "starting",
	}
)

func handleV2SparkJobCreate(t *testing.T) {
	th.Mux.HandleFunc("/batches", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedCreateResponse)
	})
}

func handleV2SparkJobGet(t *testing.T) {
	th.Mux.HandleFunc("/batches/6145a791-81a4-4edb-b2d2-ea599caf6550", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedCreateResponse)
	})
}

func handleV2SparkJobGetState(t *testing.T) {
	th.Mux.HandleFunc("/batches/6145a791-81a4-4edb-b2d2-ea599caf6550/state",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedGetStateResponse)
		})
}

func handleV2SparkJobDelete(t *testing.T) {
	th.Mux.HandleFunc("/batches/6145a791-81a4-4edb-b2d2-ea599caf6550", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
}
