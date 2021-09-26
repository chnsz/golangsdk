package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/chnsz/golangsdk"
	"github.com/chnsz/golangsdk/openstack/bms/v1/flavors"
	th "github.com/chnsz/golangsdk/testhelper"
	fake "github.com/chnsz/golangsdk/testhelper/client"
)

func TestListFlavor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/baremetalservers/flavors", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "flavors": [
        {
			"id": "physical.d2.large",
            "name": "physical.d2.large",
            "vcpus": "48",
            "ram": 196608,
            "disk": "120600",
            "swap": "",
            "links": [
                {
                    "rel": "self",
                    "href": "https://compute-ext.region.cn-north-1.hwclouds.com/v1.0/0970dd7a1300f5672ff2c003c60ae115/flavors/physical.d2.large",
                    "type": null
                },
                {
                    "rel": "bookmark",
                    "href": "https://compute-ext.region.cn-north-1.hwclouds.com/0970dd7a1300f5672ff2c003c60ae115/flavors/physical.d2.large",
                    "type": null
                }
            ],
            "OS-FLV-EXT-DATA:ephemeral": 0,
            "rxtx_factor": 1.0,
            "OS-FLV-DISABLED:disabled": false,
            "rxtx_quota": null,
            "rxtx_cap": null,
            "os-flavor-access:is_public": true,
            "os_extra_specs": {
                "capabilities:cpu_arch": "x86_64",
                "cond:operation:az": "cn-north-1a(sellout),cn-north-1b(sellout)",
                "capabilities:hypervisor_type": "ironic",
                "baremetal:__support_evs": "false",
                "baremetal:extBootType": "LocalDisk",
                "capabilities:board_type": "d2l",
                "cond:operation:status": "abandon",
                "baremetal:net_num": "2",
                "baremetal:netcard_detail": "2 x 2*10GE",
                "baremetal:disk_detail": "2*600G SAS System Disk RAID 1+ 12*10T SATA",
                "baremetal:cpu_detail": "Intel Xeon Gold 5118 V5 (2*12Core 2.30GHz)",
                "resource_type": "ironic",
                "baremetal:memory_detail": "192GB DDR4"
            }
        },
        {
            "id": "physical.io2.xlarge",
            "name": "physical.io2.xlarge",
            "vcpus": "88",
            "ram": 393216,
            "disk": "8800",
            "swap": "",
            "links": [
                {
                    "rel": "self",
                    "href": "https://compute-ext.region.cn-north-1.hwclouds.com/v1.0/0970dd7a1300f5672ff2c003c60ae115/flavors/physical.io2.xlarge",
                    "type": null
                },
                {
                    "rel": "bookmark",
                    "href": "https://compute-ext.region.cn-north-1.hwclouds.com/0970dd7a1300f5672ff2c003c60ae115/flavors/physical.io2.xlarge",
                    "type": null
                }
            ],
            "OS-FLV-EXT-DATA:ephemeral": 0,
            "rxtx_factor": 1.0,
            "OS-FLV-DISABLED:disabled": false,
            "rxtx_quota": null,
            "rxtx_cap": null,
            "os-flavor-access:is_public": true,
            "os_extra_specs": {
                "capabilities:cpu_arch": "x86_64",
                "cond:operation:az": "cn-north-1a(normal),cn-north-1b(normal)",
                "capabilities:hypervisor_type": "ironic",
                "baremetal:__support_evs": "false",
                "baremetal:extBootType": "LocalDisk",
                "capabilities:board_type": "io2xl",
                "cond:operation:status": "normal",
                "baremetal:net_num": "2",
                "baremetal:netcard_detail": "2 x 2*10GE",
                "baremetal:disk_detail": "2*800G SSD Raid 1 + 10* 800G SSD",
                "baremetal:cpu_detail": "Intel Xeon Gold 6161 V5(2*22core*2.2 GHz)",
                "resource_type": "ironic",
                "baremetal:memory_detail": "12*32GB DDR4"
            }
        }
    ]
}
			`)
	})

	actual, err := flavors.List(fake.ServiceClient(), flavors.ListOpts{}).Extract()
	if err != nil {
		t.Errorf("Failed to extract flavors: %v", err)
	}

	expected := []flavors.Flavor{
		{
			ID:         "physical.d2.large",
			Name:       "physical.d2.large",
			VCPUs:      "48",
			RAM:        196608,
			Disk:       "120600",
			IsPublic:   true,
			Disabled:   false,
			RxTxFactor: 1.0,
			Ephemeral:  0,
			Links: []golangsdk.Link{
				{
					Href: "https://compute-ext.region.cn-north-1.hwclouds.com/v1.0/0970dd7a1300f5672ff2c003c60ae115/flavors/physical.d2.large",
					Rel:  "self",
				},
				{
					Href: "https://compute-ext.region.cn-north-1.hwclouds.com/0970dd7a1300f5672ff2c003c60ae115/flavors/physical.d2.large",
					Rel:  "bookmark",
				},
			},
			OsExtraSpecs: flavors.OsExtraSpecs{
				Type:            "ironic",
				CPUArch:         "x86_64",
				FlavorType:      "d2l",
				HypervisorType:  "ironic",
				SupportEvs:      "false",
				BootFrom:        "LocalDisk",
				NetNum:          "2",
				CPUDetail:       "Intel Xeon Gold 5118 V5 (2*12Core 2.30GHz)",
				MemoryDetail:    "192GB DDR4",
				DiskDetail:      "2*600G SAS System Disk RAID 1+ 12*10T SATA",
				NetcardDetail:   "2 x 2*10GE",
				OperationStatus: "abandon",
				OperationAZ:     "cn-north-1a(sellout),cn-north-1b(sellout)",
			},
		},
		{
			ID:         "physical.io2.xlarge",
			Name:       "physical.io2.xlarge",
			VCPUs:      "88",
			RAM:        393216,
			Disk:       "8800",
			IsPublic:   true,
			Disabled:   false,
			RxTxFactor: 1.0,
			Ephemeral:  0,
			Links: []golangsdk.Link{
				{
					Href: "https://compute-ext.region.cn-north-1.hwclouds.com/v1.0/0970dd7a1300f5672ff2c003c60ae115/flavors/physical.io2.xlarge",
					Rel:  "self",
				},
				{
					Href: "https://compute-ext.region.cn-north-1.hwclouds.com/0970dd7a1300f5672ff2c003c60ae115/flavors/physical.io2.xlarge",
					Rel:  "bookmark",
				},
			},
			OsExtraSpecs: flavors.OsExtraSpecs{
				Type:            "ironic",
				CPUArch:         "x86_64",
				FlavorType:      "io2xl",
				HypervisorType:  "ironic",
				SupportEvs:      "false",
				BootFrom:        "LocalDisk",
				NetNum:          "2",
				CPUDetail:       "Intel Xeon Gold 6161 V5(2*22core*2.2 GHz)",
				MemoryDetail:    "12*32GB DDR4",
				DiskDetail:      "2*800G SSD Raid 1 + 10* 800G SSD",
				NetcardDetail:   "2 x 2*10GE",
				OperationStatus: "normal",
				OperationAZ:     "cn-north-1a(normal),cn-north-1b(normal)",
			},
		},
	}

	th.AssertDeepEquals(t, expected, actual)
}
