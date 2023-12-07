package certificate

import "github.com/chnsz/golangsdk"

func baseURL(c *golangsdk.ServiceClient, vpnGatewayID string) string {
	return c.ServiceURL("vpn-gateways", vpnGatewayID, "certificate")
}

func resourceURL(c *golangsdk.ServiceClient, vpnGatewayID string, certificateID string) string {
	return c.ServiceURL("vpn-gateways", vpnGatewayID, "certificate", certificateID)
}
