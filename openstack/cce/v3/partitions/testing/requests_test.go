package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/chnsz/golangsdk/openstack/cce/v3/common"
	"github.com/chnsz/golangsdk/openstack/cce/v3/partitions"
	th "github.com/chnsz/golangsdk/testhelper"
)

func TestListNode(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/c59fd21fd2a94963b822d8985b884673/clusters/cec124c2-58f1-11e8-ad73-0255ac101926/partitions", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "kind":"List",
    "apiVersion":"v3",
    "items":[
        {
            "kind":"Partition",
            "apiVersion":"v3",
            "metadata":{
                "name":"center",
                "creationTimestamp":"2023-03-21 07:58:56.617633 +0000 UTC"
            },
            "spec":{
                "hostNetwork":{
                    "subnetID":"82f94fd8-a10a-4088-a514-3198b44fa1b6"
                },
                "containerNetwork":[
                    {
                        "subnetID":"86e5f9b8-53ec-4a58-a6f3-594b341f30e9"
                    }
                ],
                "publicBorderGroup":"center",
                "category":"Default"
            }
        },
        {
            "kind":"Partition",
            "apiVersion":"v3",
            "metadata":{
                "name":"cn-south-1-ies-fstxz",
                "creationTimestamp":"2023-03-21 12:12:10.269567 +0000 UTC"
            },
            "spec":{
                "hostNetwork":{
                    "subnetID":"d7131ed5-f813-4dbc-86f8-bcbdc07dce6f"
                },
                "containerNetwork":[
                    {
                        "subnetID":"b2f23c46-edaa-4e66-b82f-50edafa638f5"
                    },
                    {
                        "subnetID":"dee746d5-6c78-43fb-bc36-ac26c581a3ec"
                    }
                ],
                "publicBorderGroup":"cn-south-1-ies-fstxz",
                "category":"IES"
            }
        }
    ]
}
		`)
	})

	listPartitions := partitions.ListOpts{Name: "cn-south-1-ies-fstxz"}
	actual, err := partitions.List(fake.ServiceClient(), "cec124c2-58f1-11e8-ad73-0255ac101926", listPartitions)

	if err != nil {
		t.Errorf("Failed to extract partitions: %v", err)
	}

	expected := []partitions.Partitions{
		{
			Kind:       "Partition",
			Apiversion: "v3",
			Metadata: partitions.Metadata{
				Name: "cn-south-1-ies-fstxz",
			},
			Spec: partitions.Spec{
				Category:          "IES",
				PublicBorderGroup: "cn-south-1-ies-fstxz",
				HostNetwork: partitions.HostNetwork{
					SubnetID: "d7131ed5-f813-4dbc-86f8-bcbdc07dce6f",
				},
				ContainerNetwork: []partitions.ContainerNetwork{
					{
						SubnetID: "b2f23c46-edaa-4e66-b82f-50edafa638f5",
					},
					{
						SubnetID: "dee746d5-6c78-43fb-bc36-ac26c581a3ec",
					},
				},
			},
		},
	}

	th.AssertDeepEquals(t, expected, actual)
}

func TestGetV3Partition(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/c59fd21fd2a94963b822d8985b884673/clusters/cec124c2-58f1-11e8-ad73-0255ac101926/partitions/cn-south-1-ies-fstxz", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, Output)
	})

	actual, err := partitions.Get(fake.ServiceClient(), "cec124c2-58f1-11e8-ad73-0255ac101926", "cn-south-1-ies-fstxz").Extract()
	th.AssertNoErr(t, err)
	expected := Expected
	th.AssertDeepEquals(t, expected, actual)
}
