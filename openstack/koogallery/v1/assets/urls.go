package assets

import "github.com/chnsz/golangsdk"

func listURL(sc *golangsdk.ServiceClient) string {
	return sc.ServiceURL("assset", "deployed-object")
}
