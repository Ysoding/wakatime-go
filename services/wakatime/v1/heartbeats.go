package v1

import (
	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type HeartbeatsRequest struct {
	*request.BaseRequest
	User    *string `json:"user,omitempty"`
	Current *bool   `json:"-"`

	Date *string   `json:"date,omitempty"`
	IDs  *[]string `json:"writes_only,omitempty"`
}

func NewHeartbeatsRequest() *HeartbeatsRequest {
	return &HeartbeatsRequest{
		BaseRequest: &request.BaseRequest{
			Version: "v1",
			Method:  sdk.MethodDelete,
		},
		Current: sdk.Bool(true),
	}
}

func (c *Client) Heartbeats(req *HeartbeatsRequest) (*HeartbeatsResponse, error) {
	if req == nil {
		req = NewHeartbeatsRequest()
	}

	if req.Current != nil && *req.Current {
		req.BaseRequest.URL = "users/current"
	}

	if req.User != nil {
		req.BaseRequest.URL = "users/{user}"
	}

	req.BaseRequest.URL += "/heartbeats.bulk"

	resp := NewHeartbeatsResponse()

	err := c.Send(req, resp)
	return resp, err
}

func NewHeartbeatsResponse() *HeartbeatsResponse {
	return &HeartbeatsResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

type HeartbeatsResponse struct {
	*response.BaseResponse
}
