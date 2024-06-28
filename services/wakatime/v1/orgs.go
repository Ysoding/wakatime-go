package v1

import (
	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type OrgsRequest struct {
	*request.BaseRequest
	User    *string `json:"user,omitempty"`
	Current *bool   `json:"-"`
}

func NewOrgsRequest() *OrgsRequest {
	return &OrgsRequest{
		BaseRequest: &request.BaseRequest{
			Version: "v1",
			Method:  sdk.MethodGet,
		},
		Current: sdk.Bool(true),
	}
}

func (c *Client) Orgs(req *OrgsRequest) (*OrgsResponse, error) {
	if req == nil {
		req = NewOrgsRequest()
	}

	if req.Current != nil && *req.Current {
		req.BaseRequest.URL = "users/current/orgs"
	}

	if req.User != nil {
		req.BaseRequest.URL = "users/{user}/orgs"
	}

	resp := NewOrgsResponse()

	err := c.Send(req, resp)
	return resp, err
}

func NewOrgsResponse() *OrgsResponse {
	return &OrgsResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

type OrgsResponse struct {
	*response.BaseResponse
	Data       []Organization `json:"data"`
	NextPage   int            `json:"next_page"`
	Page       int            `json:"page"`
	PrevPage   int            `json:"prev_page"`
	Total      int            `json:"total"`
	TotalPages int            `json:"total_pages"`
}
