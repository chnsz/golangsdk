package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/chnsz/golangsdk/auth"
	th "github.com/chnsz/golangsdk/testhelper"
)

const (
	ak       = "AccessKey"
	sk       = "SecretKey"
	service  = "demo"
	regionId = "test-region-1"
	host     = "example.huaweicloud.com"
	endpoint = "https://" + host
	path     = "/path"
)

type TestBody struct {
	Name string
	Id   int
}

type TestParam struct {
	Name, Method, Endpoint, Path string
	Body                         interface{}
	Queries                      map[string]interface{}
	Headers                      map[string]string
}

type TestCase struct {
	TestParam
	Expected string
}

var (
	testParam1 = TestParam{
		Name:     "test1",
		Method:   "GET",
		Endpoint: endpoint,
		Path:     path,
		Body:     nil,
		Queries:  map[string]interface{}{"limit": 1},
		Headers:  map[string]string{"X-Sdk-Date": "20060102T150405Z", "TEST_UNDERSCORE": "TEST_VALUE"},
	}
	testParam2 = TestParam{
		Name:     "test2",
		Method:   "POST",
		Endpoint: endpoint,
		Path:     path,
		Body:     &TestBody{Name: "test", Id: 1},
		Queries:  map[string]interface{}{"key": "value"},
		Headers:  map[string]string{"X-Sdk-Date": "20060102T150405Z", "TEST_UNDERSCORE": "TEST_VALUE", "Content-Type": "application/json"},
	}
)

func buildRequestBody(tc TestCase) io.Reader {
	if tc.Body == nil {
		return nil
	}

	jsonData, err := json.Marshal(tc.Body)
	if err != nil {
		panic(err)
	}

	return bytes.NewReader(jsonData)
}

func buildRequestHeaders(tc TestCase) http.Header {
	headers := make(map[string][]string)
	if tc.Headers != nil {
		for k, v := range tc.Headers {
			headers[k] = []string{v}
		}
	}
	return headers
}

func buildRequestPath(tc TestCase) string {
	queryParams := ""
	if tc.Queries != nil {
		for k, v := range tc.Queries {
			queryParams += fmt.Sprintf("&%s=%v", k, v)
		}
	}

	if queryParams != "" {
		queryParams = "?" + queryParams[1:]
	}

	return tc.Path + queryParams
}

func buildReqWithTestcase(tc TestCase) *http.Request {
	request, err := http.NewRequest(tc.Method, buildRequestPath(tc), buildRequestBody(tc))
	if err != nil {
		panic(err)
	}

	request.Header = buildRequestHeaders(tc)
	return request
}

func TestSigner_Sign(t *testing.T) {
	cases := []TestCase{
		{
			TestParam: testParam1,
			Expected: "SDK-HMAC-SHA256 Access=AccessKey, SignedHeaders=test_underscore;x-sdk-date," +
				" Signature=daeadcf4a306599f7422401cc55d3718a23ef159fed370c62c551a2c45ae9655",
		},
		{
			TestParam: testParam2,
			Expected: "SDK-HMAC-SHA256 Access=AccessKey, SignedHeaders=content-type;test_underscore;x-sdk-date," +
				" Signature=0238c3bedddfe8bf113c7530f783322380dedf720d514861623cfe5f73aa8f64",
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			req := buildReqWithTestcase(c)
			err := auth.Sign(req, ak, sk)
			th.CheckNoErr(t, err)
			th.AssertEquals(t, c.Expected, req.Header.Get("Authorization"))
		})
	}
}

func TestDerivedSigner_Sign(t *testing.T) {
	cases := []TestCase{
		{
			TestParam: testParam1,
			Expected: "V11-HMAC-SHA256 Credential=AccessKey/20060102/test-region-1/demo," +
				" SignedHeaders=test_underscore;x-sdk-date," +
				" Signature=00196d6926ef54489f9f1a1bdfa24870cadbf0af55cc1dffe26574341f1da7b0",
		},
		{
			TestParam: testParam2,
			Expected: "V11-HMAC-SHA256 Credential=AccessKey/20060102/test-region-1/demo," +
				" SignedHeaders=content-type;test_underscore;x-sdk-date," +
				" Signature=ec6f411f852fb2badae3d4927dd1a559e3a3f02a7689527a9a116f8edc7ce3fb",
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			req := buildReqWithTestcase(c)
			err := auth.SignDerived(req, ak, sk, service, regionId)
			th.CheckNoErr(t, err)
			th.AssertEquals(t, c.Expected, req.Header.Get("Authorization"))
		})
	}
}
