package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/chnsz/golangsdk"
	"github.com/chnsz/golangsdk/openstack/apigw/v2/apis"
	th "github.com/chnsz/golangsdk/testhelper"
	"github.com/chnsz/golangsdk/testhelper/client"
)

const (
	expectedCreateResponse = `
{
	"arrange_necessary": 2,
	"auth_opt": {
		"app_code_auth_type": "DISABLE"
	},
	"auth_type": "AUTHORIZER",
	"authorizer_id": "8b9c7d67ca144a5da7bc9cbccedfe753",
	"backend_api": {
		"enable_client_ssl": false,
		"req_method": "GET",
		"req_protocol": "HTTP",
		"req_uri": "/backend/users",
		"timeout": 6000,
		"url_domain": "69bb5628fce741a1b901b08cde7b814d",
		"vpc_channel_info": {
			"vpc_channel_id": "69bb5628fce741a1b901b08cde7b814d"
		},
		"vpc_channel_status": 1
	},
	"backend_params": [
		{
			"location": "HEADER",
			"name": "X-User-Name",
			"origin": "SYSTEM",
			"value": "$context.authorizer.frontend.user_name"
		}
	],
	"backend_type": "HTTP",
	"cors": false,
	"group_id": "c6e46c49a6734c918262a16c3c1a3a13",
	"group_name": "terraform_test",
	"group_version": "V1",
	"id": "cded6d80fc9f442c9842eaf854f10525",
	"match_mode": "NORMAL",
	"name": "terraform_test",
	"register_time": "2021-08-05T03:33:35.360020923Z",
	"req_method": "POST",
	"req_protocol": "HTTP",
	"req_uri": "/terraform/users",
	"status": 1,
	"type": 2,
	"update_time": "2021-08-05T03:33:35.360021226Z"
}`

	expectedGetResponse = `
{
	"arrange_necessary": 2,
	"auth_opt": {
		"app_code_auth_type": "DISABLE"
	},
	"auth_type": "AUTHORIZER",
	"authorizer_id": "8b9c7d67ca144a5da7bc9cbccedfe753",
	"backend_api": {
		"enable_client_ssl": false,
		"req_method": "GET",
		"req_protocol": "HTTP",
		"req_uri": "/backend/users",
		"timeout": 6000,
		"url_domain": "69bb5628fce741a1b901b08cde7b814d",
		"vpc_channel_info": {
			"vpc_channel_id": "69bb5628fce741a1b901b08cde7b814d"
		},
		"vpc_channel_status": 1
	},
	"backend_params": [
		{
		  "location": "HEADER",
		  "name": "X-User-Name",
		  "origin": "SYSTEM",
		  "value": "$context.authorizer.frontend.user_name"
		}
	],
	"backend_type": "HTTP",
	"cors": false,
	"group_id": "c6e46c49a6734c918262a16c3c1a3a13",
	"group_name": "terraform_test",
	"group_version": "V1",
	"id": "cded6d80fc9f442c9842eaf854f10525",
	"match_mode": "NORMAL",
	"name": "terraform_test",
	"register_time": "2021-08-05T03:33:35.360020923Z",
	"req_method": "POST",
	"req_protocol": "HTTP",
	"req_uri": "/terraform/users",
	"status": 1,
	"type": 2,
	"update_time": "2021-08-05T03:33:35Z"
}
`

	expectedListResponse = `
{
	"total": 1,
	"size": 1,
	"apis": [
		{
			"arrange_necessary": 2,
			"auth_opt": {
				"app_code_auth_type": "DISABLE"
			},
			"auth_type": "AUTHORIZER",
			"authorizer_id": "8b9c7d67ca144a5da7bc9cbccedfe753",
			"backend_api": {
				"enable_client_ssl": false,
				"req_method": "GET",
				"req_protocol": "HTTP",
				"req_uri": "/backend/users",
				"timeout": 6000,
				"url_domain": "69bb5628fce741a1b901b08cde7b814d",
				"vpc_channel_info": {
					"vpc_channel_id": "69bb5628fce741a1b901b08cde7b814d"
				},
				"vpc_channel_status": 1
			},
			"backend_params": [
				{
				  "location": "HEADER",
				  "name": "X-User-Name",
				  "origin": "SYSTEM",
				  "value": "$context.authorizer.frontend.user_name"
				}
			],
			"backend_type": "HTTP",
			"cors": false,
			"group_id": "c6e46c49a6734c918262a16c3c1a3a13",
			"group_name": "terraform_test",
			"group_version": "V1",
			"id": "cded6d80fc9f442c9842eaf854f10525",
			"match_mode": "NORMAL",
			"name": "terraform_test",
			"register_time": "2021-08-05T03:33:35.360020923Z",
			"req_method": "POST",
			"req_protocol": "HTTP",
			"req_uri": "/terraform/users",
			"status": 1,
			"type": 2,
			"update_time": "2021-08-05T03:33:35Z"
		}
	]
}`

	expectedPublishAPIResponse = `
{
	"api_id": "2b0253cba7d348f698de45abacd3ae29",
	"env_id": "c5b32727186c4fe6b60408a8a297be09",
	"publish_id": "8ecdb96299a64e10ad1c7f37a5d24bdb",
	"publish_time": "2021-10-12T07:02:12.72743121Z",
	"remark": "Version 1",
	"version_id": "eaf45032b6a649349cfbfe736d41a711"
}`

	expectedListPublishHistoriesResponse = `
{
	"api_versions": [
		{
			"api_id": "2b0253cba7d348f698de45abacd3ae29",
			"env_id": "c5b32727186c4fe6b60408a8a297be09",
			"env_name": "tf_lance_apig_environment_demo",
			"publish_time": "2021-10-12T07:02:12Z",
			"remark": "Version 1",
			"status": 1,
			"version_id": "eaf45032b6a649349cfbfe736d41a711",
			"version_no": "20211012150212"
		}
	]
}`

	expectedVersionSwitchResponse = `
{
	"api_id": "2b0253cba7d348f698de45abacd3ae29",
	"env_id": "c5b32727186c4fe6b60408a8a297be09",
	"publish_id": "b8a167517eca451f9f0a68d1e7ee96b8",
	"publish_time": "2021-10-12T07:37:00.212793535Z",
	"remark": "Version 1",
	"version_id": "eaf45032b6a649349cfbfe736d41a711"
}`

	expectedOfflineAPIResponse = `
{
	"api_id": "2b0253cba7d348f698de45abacd3ae29",
	"env_id": "c5b32727186c4fe6b60408a8a297be09",
	"api_name": "tf_lance_apig_api_demo",
	"publish_time": "0001-01-01T00:00:00Z"
}`
)

var (
	createOpts = apis.APIOpts{
		GroupId:      "c6e46c49a6734c918262a16c3c1a3a13",
		Type:         2,
		AuthType:     "AUTHORIZER",
		AuthorizerId: "8b9c7d67ca144a5da7bc9cbccedfe753",
		Cors:         golangsdk.Disabled,
		MatchMode:    "NORMAL",
		Name:         "terraform_test",
		ReqURI:       "/terraform/users",
		ReqMethod:    "POST",
		ReqProtocol:  "HTTP",
		BackendType:  "HTTP",
		WebInfo: &apis.Web{
			ClientSslEnable: golangsdk.Disabled,
			ReqMethod:       "POST",
			ReqProtocol:     "HTTP",
			ReqURI:          "/backend/users",
			Timeout:         6000,
			VpcChannelInfo: &apis.VpcChannel{
				VpcChannelId: "69bb5628fce741a1b901b08cde7b814d",
			},
			VpcChannelStatus: 1,
		},
		BackendParams: []apis.BackendParamBase{
			{
				Location: "HEADER",
				Name:     "X-User-Name",
				Origin:   "SYSTEM",
				Value:    "$context.authorizer.frontend.user_name",
			},
		},
	}

	expectedCreateResponseData = &apis.APIResp{
		ArrangeNecessary: 2,
		AuthType:         "AUTHORIZER",
		AuthorizerId:     "8b9c7d67ca144a5da7bc9cbccedfe753",
		AuthOpt: apis.AuthOpt{
			AppCodeAuthType: "DISABLE",
		},
		WebInfo: apis.Web{
			ClientSslEnable: golangsdk.Disabled,
			ReqMethod:       "GET",
			ReqProtocol:     "HTTP",
			ReqURI:          "/backend/users",
			Timeout:         6000,
			DomainURL:       "69bb5628fce741a1b901b08cde7b814d",
			VpcChannelInfo: &apis.VpcChannel{
				VpcChannelId: "69bb5628fce741a1b901b08cde7b814d",
			},
			VpcChannelStatus: 1,
		},
		BackendParams: []apis.BackendParamResp{
			{
				Location: "HEADER",
				Name:     "X-User-Name",
				Origin:   "SYSTEM",
				Value:    "$context.authorizer.frontend.user_name",
			},
		},
		BackendType:  "HTTP",
		Cors:         false,
		GroupId:      "c6e46c49a6734c918262a16c3c1a3a13",
		GroupName:    "terraform_test",
		GroupVersion: "V1",
		ID:           "cded6d80fc9f442c9842eaf854f10525",
		MatchMode:    "NORMAL",
		Name:         "terraform_test",
		RegisterTime: "2021-08-05T03:33:35.360020923Z",
		ReqMethod:    "POST",
		ReqProtocol:  "HTTP",
		ReqURI:       "/terraform/users",
		Status:       1,
		Type:         2,
		UpdateTime:   "2021-08-05T03:33:35.360021226Z",
	}

	expectedGetResponseData = &apis.APIResp{
		ArrangeNecessary: 2,
		AuthType:         "AUTHORIZER",
		AuthorizerId:     "8b9c7d67ca144a5da7bc9cbccedfe753",
		AuthOpt: apis.AuthOpt{
			AppCodeAuthType: "DISABLE",
		},
		WebInfo: apis.Web{
			ClientSslEnable: golangsdk.Disabled,
			ReqMethod:       "GET",
			ReqProtocol:     "HTTP",
			ReqURI:          "/backend/users",
			Timeout:         6000,
			DomainURL:       "69bb5628fce741a1b901b08cde7b814d",
			VpcChannelInfo: &apis.VpcChannel{
				VpcChannelId: "69bb5628fce741a1b901b08cde7b814d",
			},
			VpcChannelStatus: 1,
		},
		BackendParams: []apis.BackendParamResp{
			{
				Location: "HEADER",
				Name:     "X-User-Name",
				Origin:   "SYSTEM",
				Value:    "$context.authorizer.frontend.user_name",
			},
		},
		BackendType:  "HTTP",
		Cors:         false,
		GroupId:      "c6e46c49a6734c918262a16c3c1a3a13",
		GroupName:    "terraform_test",
		GroupVersion: "V1",
		ID:           "cded6d80fc9f442c9842eaf854f10525",
		MatchMode:    "NORMAL",
		Name:         "terraform_test",
		RegisterTime: "2021-08-05T03:33:35.360020923Z",
		ReqMethod:    "POST",
		ReqProtocol:  "HTTP",
		ReqURI:       "/terraform/users",
		Status:       1,
		Type:         2,
		UpdateTime:   "2021-08-05T03:33:35Z",
	}

	listOpts = &apis.ListOpts{
		Name: "terraform_test",
	}

	expectedListResponseData = []apis.APIResp{
		{
			ArrangeNecessary: 2,
			AuthType:         "AUTHORIZER",
			AuthorizerId:     "8b9c7d67ca144a5da7bc9cbccedfe753",
			AuthOpt: apis.AuthOpt{
				AppCodeAuthType: "DISABLE",
			},
			WebInfo: apis.Web{
				ClientSslEnable: golangsdk.Disabled,
				ReqMethod:       "GET",
				ReqProtocol:     "HTTP",
				ReqURI:          "/backend/users",
				Timeout:         6000,
				DomainURL:       "69bb5628fce741a1b901b08cde7b814d",
				VpcChannelInfo: &apis.VpcChannel{
					VpcChannelId: "69bb5628fce741a1b901b08cde7b814d",
				},
				VpcChannelStatus: 1,
			},
			BackendParams: []apis.BackendParamResp{
				{
					Location: "HEADER",
					Name:     "X-User-Name",
					Origin:   "SYSTEM",
					Value:    "$context.authorizer.frontend.user_name",
				},
			},
			BackendType:  "HTTP",
			Cors:         false,
			GroupId:      "c6e46c49a6734c918262a16c3c1a3a13",
			GroupName:    "terraform_test",
			GroupVersion: "V1",
			ID:           "cded6d80fc9f442c9842eaf854f10525",
			MatchMode:    "NORMAL",
			Name:         "terraform_test",
			RegisterTime: "2021-08-05T03:33:35.360020923Z",
			ReqMethod:    "POST",
			ReqProtocol:  "HTTP",
			ReqURI:       "/terraform/users",
			Status:       1,
			Type:         2,
			UpdateTime:   "2021-08-05T03:33:35Z",
		},
	}

	publishOpts = apis.PublishOpts{
		Action:      "online",
		ApiId:       "2b0253cba7d348f698de45abacd3ae29",
		EnvId:       "c5b32727186c4fe6b60408a8a297be09",
		Description: "Version 1",
	}

	expectedPublishResponseData = &apis.PublishResp{
		ApiId:       "2b0253cba7d348f698de45abacd3ae29",
		EnvId:       "c5b32727186c4fe6b60408a8a297be09",
		PublishId:   "8ecdb96299a64e10ad1c7f37a5d24bdb",
		PublishTime: "2021-10-12T07:02:12.72743121Z",
		Description: "Version 1",
		VersionId:   "eaf45032b6a649349cfbfe736d41a711",
	}

	listPublishHistoriesOpts = apis.ListPublishHistoriesOpts{}

	expectedListPublishHistoriesResponseData = []apis.ApiVersionInfo{
		{
			ApiId:       "2b0253cba7d348f698de45abacd3ae29",
			EnvId:       "c5b32727186c4fe6b60408a8a297be09",
			EnvName:     "tf_lance_apig_environment_demo",
			PublishTime: "2021-10-12T07:02:12Z",
			Description: "Version 1",
			Status:      1,
			VersionId:   "eaf45032b6a649349cfbfe736d41a711",
			Version:     "20211012150212",
		},
	}

	expectedVersionSwitchResponseData = &apis.PublishResp{
		ApiId:       "2b0253cba7d348f698de45abacd3ae29",
		EnvId:       "c5b32727186c4fe6b60408a8a297be09",
		PublishId:   "b8a167517eca451f9f0a68d1e7ee96b8",
		PublishTime: "2021-10-12T07:37:00.212793535Z",
		Description: "Version 1",
		VersionId:   "eaf45032b6a649349cfbfe736d41a711",
	}

	offlineOpts = apis.PublishOpts{
		Action: "offline",
		ApiId:  "2b0253cba7d348f698de45abacd3ae29",
		EnvId:  "c5b32727186c4fe6b60408a8a297be09",
	}

	expectedOfflineAPIResponseData = &apis.PublishResp{
		ApiId:       "2b0253cba7d348f698de45abacd3ae29",
		EnvId:       "c5b32727186c4fe6b60408a8a297be09",
		ApiName:     "tf_lance_apig_api_demo",
		PublishTime: "0001-01-01T00:00:00Z",
	}
)

func handleV2APICreate(t *testing.T) {
	th.Mux.HandleFunc("/instances/33fc92ffb7e749df952ecc7729d972bc/apis",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, _ = fmt.Fprint(w, expectedCreateResponse)
		})
}

func handleV2APIGet(t *testing.T) {
	th.Mux.HandleFunc("/instances/33fc92ffb7e749df952ecc7729d972bc/apis/cded6d80fc9f442c9842eaf854f10525",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedGetResponse)
		})
}

func handleV2APIList(t *testing.T) {
	th.Mux.HandleFunc("/instances/33fc92ffb7e749df952ecc7729d972bc/apis",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedListResponse)
		})
}

func handleV2APIUpdate(t *testing.T) {
	th.Mux.HandleFunc("/instances/33fc92ffb7e749df952ecc7729d972bc/apis/cded6d80fc9f442c9842eaf854f10525",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PUT")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedGetResponse)
		})
}

func handleV2APIDelete(t *testing.T) {
	th.Mux.HandleFunc("/instances/33fc92ffb7e749df952ecc7729d972bc/apis/cded6d80fc9f442c9842eaf854f10525",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		})
}

func handleV2APIPublish(t *testing.T) {
	th.Mux.HandleFunc("/instances/33fc92ffb7e749df952ecc7729d972bc/apis/action",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, _ = fmt.Fprint(w, expectedPublishAPIResponse)
		})
}

func handleV2APIVersionSwitch(t *testing.T) {
	th.Mux.HandleFunc("/instances/33fc92ffb7e749df952ecc7729d972bc/apis/publish/2b0253cba7d348f698de45abacd3ae29",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PUT")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedVersionSwitchResponse)
		})
}

func handleV2APIListPublishHistories(t *testing.T) {
	th.Mux.HandleFunc("/instances/33fc92ffb7e749df952ecc7729d972bc/apis/publish/2b0253cba7d348f698de45abacd3ae29",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedListPublishHistoriesResponse)
		})
}

func handleV2APIOffline(t *testing.T) {
	th.Mux.HandleFunc("/instances/33fc92ffb7e749df952ecc7729d972bc/apis/action",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, _ = fmt.Fprint(w, expectedOfflineAPIResponse)
		})
}
