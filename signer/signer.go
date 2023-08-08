package signer

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

const (
	DateFormat           = "20060102T150405Z"
	HeaderHost           = "host"
	HeaderXDateTime      = "X-Sdk-Date"
	HeaderXAuthorization = "Authorization"
	HeaderXContentSha256 = "X-Sdk-Content-Sha256"
	SignAlgorithm        = "SDK-HMAC-SHA256"
)

type Signer struct {
	AccessKey string
	SecretKey string
}

// Sign manipulates the http.Request instance with some required authentication headers for SK/SK auth.
func (s *Signer) Sign(request *http.Request) error {
	t := time.Now()
	request.Header.Set(HeaderXDateTime, t.UTC().Format(DateFormat))

	signedHeaders := SignedHeaders(request)
	canonicalRequest, err := CanonicalRequest(request, signedHeaders)
	if err != nil {
		return err
	}

	stringToSignStr, err := StringToSign(canonicalRequest, t)
	if err != nil {
		return err
	}
	signature, err := SignStringToSign(stringToSignStr, []byte(s.SecretKey))
	if err != nil {
		return err
	}
	authorization := authHeaderValue(signature, s.AccessKey, signedHeaders)
	request.Header.Set("Authorization", authorization)
	return nil
}

// SignedHeaders is a method used to generate signature headers.
func SignedHeaders(r *http.Request) []string {
	var signedHeaders []string
	for key := range r.Header {
		signedHeaders = append(signedHeaders, strings.ToLower(key))
	}
	sort.Strings(signedHeaders)
	return signedHeaders
}

func hmacsha256(keyByte []byte, dataStr string) ([]byte, error) {
	hm := hmac.New(sha256.New, []byte(keyByte))
	if _, err := hm.Write([]byte(dataStr)); err != nil {
		return nil, err
	}
	return hm.Sum(nil), nil
}

// CanonicalRequest builds canonical string depends the official document for signing.
func CanonicalRequest(request *http.Request, signedHeaders []string) (string, error) {
	var hexEncode string

	if hex := request.Header.Get(HeaderXContentSha256); hex != "" {
		hexEncode = hex
	} else {
		bodyData, err := requestPayload(request)
		if err != nil {
			return "", err
		}
		hexEncode, err = HexEncodeSHA256Hash(bodyData)
		if err != nil {
			return "", err
		}
	}

	return strings.Join([]string{
		request.Method,
		CanonicalURI(request),
		CanonicalQueryString(request),
		CanonicalHeaders(request, signedHeaders),
		strings.Join(signedHeaders, ";"),
		hexEncode,
	}, "\n"), nil
}

func requestPayload(request *http.Request) ([]byte, error) {
	if request.Body == nil {
		return []byte(""), nil
	}
	bodyByte, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return []byte(""), err
	}
	request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByte))
	return bodyByte, err
}

// HexEncodeSHA256Hash returns hexcode of sha256.
func HexEncodeSHA256Hash(body []byte) (string, error) {
	hashStruct := sha256.New()
	if len(body) == 0 {
		body = []byte("")
	}
	_, err := hashStruct.Write(body)
	return fmt.Sprintf("%x", hashStruct.Sum(nil)), err
}

// CanonicalURI builds the valid URL path for signing.
func CanonicalURI(req *http.Request) string {
	pattens := strings.Split(req.URL.Path, "/")

	var uri []string
	for _, v := range pattens {
		uri = append(uri, escape(v))
	}

	urlPath := strings.Join(uri, "/")
	if len(urlPath) == 0 || urlPath[len(urlPath)-1] != '/' {
		urlPath += "/"
	}
	return urlPath
}

// CanonicalQueryString is a method used to generate RFC 3986 canonical URI path.
func CanonicalQueryString(request *http.Request) string {
	var keys []string
	queryMap := request.URL.Query()
	for key := range queryMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var query []string
	for _, key := range keys {
		k := escape(key)
		sort.Strings(queryMap[key])
		for _, v := range queryMap[key] {
			kv := fmt.Sprintf("%s=%s", k, escape(v))
			query = append(query, kv)
		}
	}
	queryStr := strings.Join(query, "&")
	request.URL.RawQuery = queryStr
	return queryStr
}

// CanonicalHeaders is a method used to generate  canonical message header (a list of request message headers).
func CanonicalHeaders(request *http.Request, signerHeaders []string) string {
	var canonicalHeaders []string
	header := make(map[string][]string)
	for k, v := range request.Header {
		header[strings.ToLower(k)] = v
	}
	for _, key := range signerHeaders {
		value := header[key]
		if strings.EqualFold(key, HeaderHost) {
			value = []string{request.Host}
		}
		sort.Strings(value)
		for _, v := range value {
			canonicalHeaders = append(canonicalHeaders, key+":"+strings.TrimSpace(v))
		}
	}
	return fmt.Sprintf("%s\n", strings.Join(canonicalHeaders, "\n"))
}

func StringToSign(canonicalRequest string, t time.Time) (string, error) {
	hashStruct := sha256.New()
	_, err := hashStruct.Write([]byte(canonicalRequest))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s\n%s\n%x",
		SignAlgorithm, t.UTC().Format(DateFormat), hashStruct.Sum(nil)), nil
}

// SignStringToSign is a method used to generate signature using given signing string and signing Key.
func SignStringToSign(stringToSign string, signingKey []byte) (string, error) {
	hmsha, err := hmacsha256(signingKey, stringToSign)
	return fmt.Sprintf("%x", hmsha), err
}

// Get the finalized value for the "Authorization" header. The signature parameter is the output from SignStringToSign
func authHeaderValue(signatureStr, accessKeyStr string, signedHeaders []string) string {
	return fmt.Sprintf("%s Access=%s, SignedHeaders=%s, Signature=%s", SignAlgorithm, accessKeyStr, strings.Join(signedHeaders, ";"), signatureStr)
}
