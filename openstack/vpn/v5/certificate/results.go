package certificate

import "github.com/chnsz/golangsdk"

type commonResult struct {
	golangsdk.Result
}

type CreateResult struct {
	commonResult
}

type UpdateResult struct {
	commonResult
}

type GetResult struct {
	commonResult
}

type GetCertificate struct {
	ID                           string `json:"id"`
	Name                         string `json:"name"`
	ProjectID                    string `json:"project_id"`
	VPNGatewayID                 string `json:"vgw_id"`
	Status                       string `json:"status"`
	Issuer                       string `json:"issuer"`
	SignatureAlgorithm           string `json:"signature_algorithm"`
	CertificateSerialNumber      string `json:"certificate_serial_number"`
	CertificateSubject           string `json:"certificate_subject"`
	CertificateExpireTime        string `json:"certificate_expire_time"`
	CertificateChainSerialNumber string `json:"certificate_chain_serial_number"`
	CertificateChainSubject      string `json:"certificate_chain_subject"`
	CertificateChainExpireTime   string `json:"certificate_chain_expire_time"`
	EncCertificateSerialNumber   string `json:"enc_certificate_serial_number"`
	EncCertificateSubject        string `json:"enc_certificate_subject"`
	EncCertificateExpireTime     string `json:"enc_certificate_expire_time"`
	CreatedAt                    string `json:"created_at"`
	UpdatedAt                    string `json:"updated_at"`
}

type GetResponseBody struct {
	RequestID   string         `json:"request_id"`
	Certificate GetCertificate `json:"certificate"`
}

type ResourceCertificate struct {
	ID                           string `json:"id"`
	Name                         string `json:"name"`
	ProjectID                    string `json:"project_id"`
	VPNGatewayID                 string `json:"vgw_id"`
	Issuer                       string `json:"issuer"`
	SignatureAlgorithm           string `json:"signature_algorithm"`
	CertificateSerialNumber      string `json:"certificate_serial_number"`
	CertificateSubject           string `json:"certificate_subject"`
	CertificateExpireTime        string `json:"certificate_expire_time"`
	CertificateChainSerialNumber string `json:"certificate_chain_serial_number"`
	CertificateChainSubject      string `json:"certificate_chain_subject"`
	CertificateChainExpireTime   string `json:"certificate_chain_expire_time"`
	EncCertificateSerialNumber   string `json:"enc_certificate_serial_number"`
	EncCertificateSubject        string `json:"enc_certificate_subject"`
	EncCertificateExpireTime     string `json:"enc_certificate_expire_time"`
	CreatedAt                    string `json:"created_at"`
}

type ResourceResponseBody struct {
	RequestID   string              `json:"request_id"`
	Certificate ResourceCertificate `json:"certificate"`
}

func (r commonResult) Extract() (ResourceResponseBody, error) {
	var result ResourceResponseBody
	err := r.Result.ExtractInto(&result)
	return result, err
}

func (r GetResult) Extract() (GetResponseBody, error) {
	var result GetResponseBody
	err := r.Result.ExtractInto(&result)
	return result, err
}
