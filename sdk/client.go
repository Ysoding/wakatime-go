package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
	"github.com/corpix/uarand"
)

var defaultUserAgent = uarand.GetRandom()

type SignFunc func(*http.Request) error

type Client struct {
	signMethod string
	Config     *Config
	Credential *Credential
}

func (c *Client) Init() *Client {
	c.signMethod = SignBasic
	return c
}

func (c *Client) WithConfig(config *Config) *Client {
	c.Config = config
	return c
}

func (c *Client) WithCredential(cred *Credential) *Client {
	c.Credential = cred
	return c
}

func (c *Client) WithSecret(secretID, secretKey string) *Client {
	c.Credential = NewCredentials(secretID, secretKey)
	return c
}

func (c *Client) Send(req request.Request, resp response.Response) error {
	method := req.GetMethod()

	builder := GetParameterBuilder(method)
	jsonReq, _ := json.Marshal(req)

	encodedURL, err := builder.BuildURL(req.GetURL(), jsonReq)
	if err != nil {
		return err
	}

	endPoint := c.Config.Endpoint
	if endPoint == "" {
		endPoint = defaultEndpoint
	}

	reqURL := fmt.Sprintf("%s://%s/%s/%s", c.Config.Scheme, endPoint, req.GetVersion(), encodedURL)
	fmt.Println(reqURL)

	body, err := builder.BuildBody(jsonReq)
	if err != nil {
		return err
	}

	sign := func(r *http.Request) error {
		signer := NewSigner(c.signMethod, c.Credential)
		_ = signer.Sign(r)
		return nil
	}

	rawResp, err := c.doSend(method, reqURL, body, req.GetHeaders(), sign)
	if err != nil {
		return err
	}

	return response.ParseFromHttpResponse(rawResp, resp)
}

func (c *Client) doSend(method, url, data string, headers map[string]string, sign SignFunc) (*http.Response, error) {
	client := &http.Client{Timeout: c.Config.Timeout}

	req, err := http.NewRequest(method, url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	c.setHeaders(req, headers)

	err = sign(req)
	if err != nil {
		return nil, err
	}

	return client.Do(req)
}

func (c *Client) setHeaders(req *http.Request, headers map[string]string) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", defaultUserAgent)

	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

func Bool(v bool) *bool       { return &v }
func Int(v int) *int          { return &v }
func Int64(v int64) *int64    { return &v }
func String(v string) *string { return &v }
