package certificate

import "github.com/chnsz/golangsdk"

type Opts struct {
	Name             string `json:"name,omitempty"`
	Certificate      string `json:"certificate,omitempty"`
	PrivateKey       string `json:"private_key,omitempty"`
	CertificateChain string `json:"certificate_chain,omitempty"`
	EncCertificate   string `json:"enc_certificate,omitempty"`
	EncPrivateKey    string `json:"enc_private_key,omitempty"`
}

func Create(client *golangsdk.ServiceClient, vpnGatewayID string, opts Opts) (r CreateResult) {
	b, err := golangsdk.BuildRequestBody(opts, "certificate")
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(baseURL(client, vpnGatewayID), b, &r.Body, nil)

	return
}

func Update(client *golangsdk.ServiceClient, vpnGatewayID string, certificateID string, opts Opts) (r UpdateResult) {
	b, err := golangsdk.BuildRequestBody(opts, "certificate")
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(resourceURL(client, vpnGatewayID, certificateID), b, &r.Body, nil)

	return
}

func Get(client *golangsdk.ServiceClient, vpnGatewayID string) (r GetResult) {
	_, r.Err = client.Get(baseURL(client, vpnGatewayID), &r.Body, nil)
	return
}
