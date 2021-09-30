package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/chnsz/golangsdk/testhelper"
	"github.com/chnsz/golangsdk/testhelper/client"
)

//mock response body string
var (
	mockListResponse = `
{
    "total_number": 2,
    "stream_names": [
        "dis-ML2k",
        "dis-XDPf"
    ],
    "stream_info_list": [
        {
            "private": false,
            "stream_id": "2kE3HAYBm5erI2sFtLb",
            "stream_name": "dis-ML2k",
            "create_time": 1632639303534,
            "retention_period": 24,
            "status": "RUNNING",
            "stream_type": "ADVANCED",
            "data_type": "CSV",
            "partition_count": 1,
            "tags": [
                {
                    "key": "key",
                    "value": "foo"
                }
            ],
            "sys_tags": [
                {
                    "key": "_sys_enterprise_project_id",
                    "value": "0"
                }
            ],
            "data_schema": "{\"type\":\"record\",\"name\":\"RecordName\",\"fields\":[{\"type\":\"string\",\"name\":\"name\"},{\"type\":\"int\",\"name\":\"age\"}]}",
            "auto_scale_enabled": true,
            "auto_scale_min_partition_count": 1,
            "auto_scale_max_partition_count": 1
        },
        {
            "workSpaceId": "0970dd7a1300f5672ff2c003c60ae115",
            "dayuInstanceId": "7f9529baa12645469ca5f89f4c4395f0",
            "private": false,
            "stream_id": "mz3V7jcVHcNDuaS3VpD",
            "stream_name": "dis-XDPf",
            "create_time": 1631344533496,
            "retention_period": 24,
            "status": "RUNNING",
            "stream_type": "COMMON",
            "data_type": "BLOB",
            "partition_count": 1,
            "tags": [],
            "sys_tags": [
                {
                    "key": "_sys_enterprise_project_id",
                    "value": "0"
                }
            ],
            "auto_scale_enabled": false,
            "auto_scale_min_partition_count": 0,
            "auto_scale_max_partition_count": 0
        }
    ],
    "has_more_streams": false
}
`

	mockGetResponse = `
{
    "stream_id": "2kE3HAYBm5erI2sFtLb",
    "stream_name": "dis-ML2k",
    "create_time": 1632639303534,
    "last_modified_time": 1632642001517,
    "retention_period": 24,
    "status": "RUNNING",
    "stream_type": "ADVANCED",
    "data_type": "CSV",
    "data_schema": "{\"type\":\"record\",\"name\":\"RecordName\",\"fields\":[{\"type\":\"string\",\"name\":\"name\"},{\"type\":\"int\",\"name\":\"age\"}]}",
    "csv_properties": {
        "delimiter": ";"
    },
    "writable_partition_count": 1,
    "readable_partition_count": 2,
    "partitions": [
        {
            "status": "ACTIVE",
            "partition_id": "shardId-0000000000",
            "hash_range": "[0 : 9223372036854775807]",
            "sequence_number_range": "[0 : 0]"
        },
        {
            "status": "DELETED",
            "partition_id": "shardId-0000000001",
            "hash_range": "[-1 : -1]",
            "sequence_number_range": "[0 : 0]"
        }
    ],
    "has_more_partitions": false,
    "tags": [
        {
            "key": "key",
            "value": "foo"
        }
    ],
    "sys_tags": [
        {
            "key": "_sys_enterprise_project_id",
            "value": "0"
        }
    ],
    "auto_scale_enabled": true,
    "auto_scale_min_partition_count": 1,
    "auto_scale_max_partition_count": 1
}
`
)

func handleList(t *testing.T) {
	th.Mux.HandleFunc("/streams", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, mockListResponse)
	})
}

func handleGet(t *testing.T) {
	th.Mux.HandleFunc("/streams/dis-ML2k", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, mockGetResponse)
	})
}
