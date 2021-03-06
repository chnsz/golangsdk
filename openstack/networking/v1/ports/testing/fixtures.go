package testing

import (
	"fmt"
	"net/http"
	"testing"

	client "github.com/chnsz/golangsdk/openstack/networking/v1/common"
	"github.com/chnsz/golangsdk/openstack/networking/v1/ports"
	th "github.com/chnsz/golangsdk/testhelper"
)

var (
	expectedCreateResponse = fmt.Sprintf(`
{
	"port": %s
}`, buildPortItemResponse("192.168.0.162"))

	expectedUpdateResponse = fmt.Sprintf(`
{
	"port": %s
}`, buildPortItemResponse("192.168.0.25"))

	expectedListResponse = fmt.Sprintf(`
{
	"ports": [
		%s
	]
}`, buildPortItemResponse("192.168.0.162"))

	emptyListResponse = `
{
	"ports": []
}`

	expectedCreateResponseData = buildPortResponseData("192.168.0.162")
	expectedUpdateResponseData = buildPortResponseData("192.168.0.25")
	expectedListResponseData   = []ports.Port{buildPortResponseData("192.168.0.162")}
)

func buildPortItemResponse(address string) string {
	return fmt.Sprintf(`
	{
		"admin_state_up": false,
		"allowed_address_pairs": [
		  {
			"ip_address": "192.168.0.25",
			"mac_address": "fa:16:3e:71:db:e5"
		  }
		],
		"binding:vnic_type": "normal",
		"created_at": "2022-03-14T06:18:48",
		"device_owner": "neutron:VIP_PORT",
		"fixed_ips": [
			{
				"ip_address": "%s",
				"subnet_id": "885cb8c3-0cbe-406d-83d6-fc98856fcf26"
			}
		],
		"id": "05547c10-e318-4067-9db2-01f5dc30be38",
		"mac_address": "fa:16:3e:71:db:e5",
		"network_id": "e4cb3b49-78a0-479b-b37d-bd99b3ec0d8a",
		"status": "DOWN"
	}`, address)
}

func buildPortResponseData(address string) ports.Port {
	return ports.Port{
		AdminStateUp: false,
		AllowedAddressPairs: []ports.AddressPair{
			{
				IpAddress:  "192.168.0.25",
				MacAddress: "fa:16:3e:71:db:e5",
			},
		},
		VnicType:    "normal",
		CreatedAt:   "2022-03-14T06:18:48",
		DeviceOwner: "neutron:VIP_PORT",
		FixedIps: []ports.FixedIp{
			{
				IpAddress: address,
				SubnetId:  "885cb8c3-0cbe-406d-83d6-fc98856fcf26",
			},
		},
		ID:         "05547c10-e318-4067-9db2-01f5dc30be38",
		NetworkId:  "e4cb3b49-78a0-479b-b37d-bd99b3ec0d8a",
		MacAddress: "fa:16:3e:71:db:e5",
		Status:     "DOWN",
	}
}

func handleV1NetworkPortCreate(t *testing.T) {
	th.Mux.HandleFunc("/v1/85636478b0bd8e67e89469c7749d4127/ports", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprint(w, expectedCreateResponse)
	})
}

func handleV1NetworkPortGet(t *testing.T) {
	th.Mux.HandleFunc("/v1/85636478b0bd8e67e89469c7749d4127/ports/05547c10-e318-4067-9db2-01f5dc30be38", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, expectedCreateResponse)
	})
}

func handleV1NetworkPortUpdate(t *testing.T) {
	th.Mux.HandleFunc("/v1/85636478b0bd8e67e89469c7749d4127/ports/05547c10-e318-4067-9db2-01f5dc30be38", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, expectedUpdateResponse)
	})
}

func handleV1NetworkPortDelete(t *testing.T) {
	th.Mux.HandleFunc("/v1/85636478b0bd8e67e89469c7749d4127/ports/05547c10-e318-4067-9db2-01f5dc30be38", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})
}

func handleV1NetworkPortList(t *testing.T) {
	th.Mux.HandleFunc("/v1/85636478b0bd8e67e89469c7749d4127/ports", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		r.ParseForm()
		marker := r.Form.Get("marker")
		if marker == "" {
			fmt.Fprint(w, expectedListResponse)
		} else {
			fmt.Fprint(w, emptyListResponse)
		}
	})
}
