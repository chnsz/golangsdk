package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/chnsz/golangsdk"
	"github.com/chnsz/golangsdk/openstack/compute/v2/extensions/quotasets"
	th "github.com/chnsz/golangsdk/testhelper"
	"github.com/chnsz/golangsdk/testhelper/client"
)

// GetOutput is a sample response to a Get call.
const GetOutput = `
{
   "quota_set" : {
      "instances" : 25,
      "security_groups" : 10,
      "security_group_rules" : 20,
      "cores" : 200,
      "injected_file_content_bytes" : 10240,
      "injected_files" : 5,
      "metadata_items" : 128,
      "ram" : 200000,
      "key_pairs" : 10,
      "injected_file_path_bytes" : 255,
	  "server_groups" : 2,
	  "server_group_members" : 3
   }
}
`

// GetDetailsOutput is a sample response to a Get call with the detailed option.
const GetDetailsOutput = `
{
   "quota_set" : {
	  "id": "555544443333222211110000ffffeeee",
      "instances" : {
          "in_use": 0,
          "limit": 25,
          "reserved": 0
      },
      "security_groups" : {
          "in_use": 0,
          "limit": 10,
          "reserved": 0
      },
      "security_group_rules" : {
          "in_use": 0,
          "limit": 20,
          "reserved": 0
      },
      "cores" : {
          "in_use": 0,
          "limit": 200,
          "reserved": 0
      },
      "injected_file_content_bytes" : {
          "in_use": 0,
          "limit": 10240,
          "reserved": 0
      },
      "injected_files" : {
          "in_use": 0,
          "limit": 5,
          "reserved": 0
      },
      "metadata_items" : {
          "in_use": 0,
          "limit": 128,
          "reserved": 0
      },
      "ram" : {
          "in_use": 0,
          "limit": 200000,
          "reserved": 0
      },
      "key_pairs" : {
          "in_use": 0,
          "limit": 10,
          "reserved": 0
      },
      "injected_file_path_bytes" : {
          "in_use": 0,
          "limit": 255,
          "reserved": 0
      },
      "server_groups" : {
          "in_use": 0,
          "limit": 2,
          "reserved": 0
      },
      "server_group_members" : {
          "in_use": 0,
          "limit": 3,
          "reserved": 0
      }
   }
}
`
const FirstTenantID = "555544443333222211110000ffffeeee"

// FirstQuotaSet is the first result in ListOutput.
var FirstQuotaSet = quotasets.QuotaSet{
	FixedIPs:                 0,
	FloatingIPs:              0,
	InjectedFileContentBytes: 10240,
	InjectedFilePathBytes:    255,
	InjectedFiles:            5,
	KeyPairs:                 10,
	MetadataItems:            128,
	RAM:                      200000,
	SecurityGroupRules:       20,
	SecurityGroups:           10,
	Cores:                    200,
	Instances:                25,
	ServerGroups:             2,
	ServerGroupMembers:       3,
}

// FirstQuotaDetailsSet is the first result in ListOutput.
var FirstQuotaDetailsSet = quotasets.QuotaDetailSet{
	ID:                       FirstTenantID,
	InjectedFileContentBytes: quotasets.QuotaDetail{InUse: 0, Reserved: 0, Limit: 10240},
	InjectedFilePathBytes:    quotasets.QuotaDetail{InUse: 0, Reserved: 0, Limit: 255},
	InjectedFiles:            quotasets.QuotaDetail{InUse: 0, Reserved: 0, Limit: 5},
	KeyPairs:                 quotasets.QuotaDetail{InUse: 0, Reserved: 0, Limit: 10},
	MetadataItems:            quotasets.QuotaDetail{InUse: 0, Reserved: 0, Limit: 128},
	RAM:                      quotasets.QuotaDetail{InUse: 0, Reserved: 0, Limit: 200000},
	SecurityGroupRules:       quotasets.QuotaDetail{InUse: 0, Reserved: 0, Limit: 20},
	SecurityGroups:           quotasets.QuotaDetail{InUse: 0, Reserved: 0, Limit: 10},
	Cores:                    quotasets.QuotaDetail{InUse: 0, Reserved: 0, Limit: 200},
	Instances:                quotasets.QuotaDetail{InUse: 0, Reserved: 0, Limit: 25},
	ServerGroups:             quotasets.QuotaDetail{InUse: 0, Reserved: 0, Limit: 2},
	ServerGroupMembers:       quotasets.QuotaDetail{InUse: 0, Reserved: 0, Limit: 3},
}

//The expected update Body. Is also returned by PUT request
const UpdateOutput = `{"quota_set":{"cores":200,"fixed_ips":0,"floating_ips":0,"injected_file_content_bytes":10240,"injected_file_path_bytes":255,"injected_files":5,"instances":25,"key_pairs":10,"metadata_items":128,"ram":200000,"security_group_rules":20,"security_groups":10,"server_groups":2,"server_group_members":3}}`

//The expected partialupdate Body. Is also returned by PUT request
const PartialUpdateBody = `{"quota_set":{"cores":200, "force":true}}`

//Result of Quota-update
var UpdatedQuotaSet = quotasets.UpdateOpts{
	FixedIPs:                 golangsdk.IntToPointer(0),
	FloatingIPs:              golangsdk.IntToPointer(0),
	InjectedFileContentBytes: golangsdk.IntToPointer(10240),
	InjectedFilePathBytes:    golangsdk.IntToPointer(255),
	InjectedFiles:            golangsdk.IntToPointer(5),
	KeyPairs:                 golangsdk.IntToPointer(10),
	MetadataItems:            golangsdk.IntToPointer(128),
	RAM:                      golangsdk.IntToPointer(200000),
	SecurityGroupRules:       golangsdk.IntToPointer(20),
	SecurityGroups:           golangsdk.IntToPointer(10),
	Cores:                    golangsdk.IntToPointer(200),
	Instances:                golangsdk.IntToPointer(25),
	ServerGroups:             golangsdk.IntToPointer(2),
	ServerGroupMembers:       golangsdk.IntToPointer(3),
}

// HandleGetSuccessfully configures the test server to respond to a Get request for sample tenant
func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-quota-sets/"+FirstTenantID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

// HandleGetDetailSuccessfully configures the test server to respond to a Get Details request for sample tenant
func HandleGetDetailSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-quota-sets/"+FirstTenantID+"/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetDetailsOutput)
	})
}

// HandlePutSuccessfully configures the test server to respond to a Put request for sample tenant
func HandlePutSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-quota-sets/"+FirstTenantID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, UpdateOutput)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateOutput)
	})
}

// HandlePartialPutSuccessfully configures the test server to respond to a Put request for sample tenant that only containes specific values
func HandlePartialPutSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-quota-sets/"+FirstTenantID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, PartialUpdateBody)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateOutput)
	})
}

// HandleDeleteSuccessfully configures the test server to respond to a Delete request for sample tenant
func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-quota-sets/"+FirstTenantID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestBody(t, r, "")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(202)
	})
}
