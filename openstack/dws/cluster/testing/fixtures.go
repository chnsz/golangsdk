package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/chnsz/golangsdk/openstack/dli/v1/queues"
	th "github.com/chnsz/golangsdk/testhelper"
	"github.com/chnsz/golangsdk/testhelper/client"
)

// expected object
var (
	expectedListResponseData = queues.Queue{
		QueueName:           "tf_acc_test_dli_queue_h8yr3",
		Description:         "",
		Owner:               "niuzhenguo",
		CreateTime:          1626771444081,
		QueueType:           "sql",
		CuCount:             16,
		ChargingMode:        1,
		ResourceId:          "b483aa39-ab65-442b-8c20-2f578a9d868b",
		EnterpriseProjectId: "",
		CidrInVpc:           "172.16.0.0/12",
		CidrInMgntsubnet:    "172.24.128.0/18",
		CidrInSubnet:        "172.24.0.0/18",
		ResourceMode:        1,
		Platform:            "x86_64",
		IsRestarting:        false,
		Feature:             "basic",
		QueueResourceType:   "vm",
	}
)

//mock response body string
var (
	mockListResponse = `
{
    "clusters": [
        {
            "id": "7d85f602-a948-4a30-afd4-e84f47471c15",
            "status": "AVAILABLE",
            "sub_status": "READONLY",
            "task_status": "SNAPSHOTTING",
            "action_progress": {
                "SNAPSHOTTING": "20%"
            },
            "node_type": "dws.d1.xlarge.ultrahigh",
            "subnet_id": "374eca02-cfc4-4de7-8ab5-dbebf7d9a720",
            "security_group_id": "dc3ec145-9029-4b39-b5a3-ace5a01f772b",
            "number_of_node": 3,
            "availability_zone": "cn-north-4b",
            "port": 8000,
            "name": "dws-1",
            "version": "1.2.0",
            "vpc_id": "85b20d7e-9eb7-4b2a-98f3-3c8843ea3574",
            "user_name": "dbadmin",
            "public_ip": {
                "public_bind_type": "auto_assign",
                "eip_id": "85b20d7e-9eb7-4b2a-98f3-3c8843ea3574"
            },
            "public_endpoints": [
                {
                    "public_connect_info": "dws-1.cn-north-4.myhuaweicloud.com",
                    "jdbc_url": "jdbc:postgresql://dws-1.cn-north-4.myhuaweicloud.com/<YOUR_DATABASE_name>"
                }
            ],
            "endpoints": [
                {
                    "connect_info": "dws-1.cn-north-4.myhuaweicloud.com",
                    "jdbc_url": "jdbc:postgresql://dws-1.cn-north-4.myhuaweicloud.com/<YOUR_DATABASE_name>"
                }
            ],
            "updated": "2016-02-10T14:28:14Z",
            "created": "2016-02-10T14:26:14Z",
            "enterprise_project_id": "aca4e50a-266f-4786-827c-f8d6cc3fbada",
            "recent_event": 6
        }
    ]
}
`

	mockGetResponse = `
{
    "cluster": {
        "id": "7d85f602-a948-4a30-afd4-e84f47471c15",
        "status": "AVAILABLE",
        "name": "dws-1",
        "updated": "2018-02-10T14:28:14Z",
        "created": "2018-02-10T14:28:14Z",
        "user_name": "dbadmin",
        "sub_status": "READONLY",
        "task_status": "SNAPSHOTTING",
        "action_progress": {
            "SNAPSHOTTING": "20%"
        },
        "node_type": "dws.m1.xlarge.ultrahigh",
        "subnet_id": "374eca02-cfc4-4de7-8ab5-dbebf7d9a720",
        "security_group_id": "dc3ec145-9029-4b39-b5a3-ace5a01f772b",
        "number_of_node": 3,
        "availability_zone": "cn-north-4b",
        "port": 8000,
        "vpc_id": "85b20d7e-9eb7-4b2a-98f3-3c8843ea3574",
        "public_ip": {
            "public_bind_type": "auto_assign",
            "eip_id": "85b20d7e-9etypeb2a-98f3-3c8843ea3574",
            "eip_address": "100.95.157.20"
        },
        "private_ip": [
            "192.168.0.12",
            "192.168.0.66"
        ],
        "public_endpoints": [
            {
                "public_connect_info": "dws-1.cn-north-4.myhuaweicloud.com",
                "jdbc_url": "jdbc:postgresql://dws-1.cn-north-4.myhuaweicloud.com/<YOUR_DATABASE_name>"
            }
        ],
        "endpoints": [
            {
                "connect_info": "dws-1.cn-north-4.myhuaweicloud.com",
                "jdbc_url": "jdbc:postgresql://dws-1.cn-north-4.myhuaweicloud.com/<YOUR_DATABASE_name>"
            }
        ],
        "version": "1.2.0",
        "maintain_window": {
            "day": "Wed",
            "start_time": "22:00",
            "end_time": "02:00"
        },
        "resize_info": {
            "target_node_num": "6",
            "origin_node_num": "3",
            "status": "GROWING",
            "start_time": "2018-02-14T14:28:14Z",
            "origin_node_type": "dws.m1.xlarge.ultrahigh",
            "target_node_type": "dws.d2.xlarge"
        },
        "enterprise_project_id": "6a6a18fe-417a-4188-9214-75fd08c22065",
        "recent_event": 6,
        "tags": [
            {
                "key": "key1",
                "value": "value1"
            },
            {
                "key": "key2",
                "value": "value2"
            }
        ],
        "parameter_group": {
            "id": "157e9cc4-64a8-11e8-adc0-fa7ae01bbebc",
            "name": "Default-Parameter-Group-dws ",
            "status": "In-Sync"
        }
    }
}
`
	mockNodeTypeResponse = `
{
    "node_types": [
        {
            "spec_name": "dws.d2.xlarge",
            "id": "ebe532d6-665f-40e6-a4d4-3c51545b6a67",
            "detail": [
                {
                    "type": "vCPU",
                    "value": "4"
                },
                {
                    "value": "1675",
                    "type": "LOCAL_DISK",
                    "unit": "GB"
                },
                {
                    "type": "mem",
                    "value": "32",
                    "unit": "GB"
                }
            ]
        },
        {
            "spec_name": "dws.m1.xlarge.ultrahigh",
            "id": "ebe532d6-665f-40e6-a4d4-3c51545b4f71",
            "detail": [
                {
                    "type": "vCPU",
                    "value": "4"
                },
                {
                    "value": "512",
                    "type": "SSD",
                    "unit": "GB"
                },
                {
                    "type": "mem",
                    "value": "32",
                    "unit": "GB"
                }
            ]
        }
    ]
}
`
)

func handleList(t *testing.T) {
	th.Mux.HandleFunc("/clusters", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, mockListResponse)
	})
}

func handleGet(t *testing.T) {
	th.Mux.HandleFunc("/clusters/7d85f602-a948-4a30-afd4-e84f47471c15", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, mockGetResponse)
	})
}

func handleNodeType(t *testing.T) {
	th.Mux.HandleFunc("/node-types", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, mockNodeTypeResponse)
	})
}
