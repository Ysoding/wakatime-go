package sdk

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

type Signer interface {
	Sign(r *http.Request) http.Header
}

type BaseSignature struct {
	Credentials *Credential
}

type SignatureV1 struct {
	BaseSignature
}

func NewSigner(signMethod string, cred *Credential) Signer {
	switch signMethod {
	case SignBasic:
		return NewSignatureV1(cred)
	default:
		return NewSignatureV1(cred)
	}
}

func NewSignatureV1(cred *Credential) Signer {
	return &SignatureV1{
		BaseSignature: BaseSignature{
			Credentials: cred,
		},
	}
}

func (s SignatureV1) Sign(r *http.Request) http.Header {
	r.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(s.Credentials.SecretKey))))
	return r.Header
}
