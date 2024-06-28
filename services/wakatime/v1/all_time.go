package v1

import (
	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type AllTimeRequest struct {
	*request.BaseRequest
	Project *string `json:"project,omitempty"`
	User    *string `json:"user,omitempty"`
	Current *bool   `json:"-"`
}

func NewAllTimeRequest() *AllTimeRequest {
	return &AllTimeRequest{
		BaseRequest: &request.BaseRequest{
			Version: "v1",
			Method:  sdk.MethodGet,
		},
		Current: sdk.Bool(true),
	}
}

func (c *Client) AllTime(req *AllTimeRequest) (*AllTimeResponse, error) {
	if req == nil {
		req = NewAllTimeRequest()
	}

	if req.Current != nil && *req.Current {
		req.BaseRequest.URL = "users/current/all_time_since_today"
	}

	if req.User != nil {
		req.BaseRequest.URL = "users/{user}/all_time_since_today"
	}

	resp := NewAllTimeResponse()

	err := c.Send(req, resp)
	return resp, err
}

func NewAllTimeResponse() *AllTimeResponse {
	return &AllTimeResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

type AllTimeResponse struct {
	*response.BaseResponse
	Data TimeData `json:"data"`
}
