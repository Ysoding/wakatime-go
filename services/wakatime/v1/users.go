package v1

import (
	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type UserRequest struct {
	*request.BaseRequest
	User    *string `json:"user_id,omitempty"`
	Current *bool   `json:"-"`
}

func NewUserRequest() *UserRequest {
	return &UserRequest{
		BaseRequest: &request.BaseRequest{
			Version: "v1",
			Method:  sdk.MethodGet,
		},
		Current: sdk.Bool(true),
	}
}

func (c *Client) User(req *UserRequest) (*UserResponse, error) {
	if req == nil {
		req = NewUserRequest()
	}

	if req.Current != nil && *req.Current {
		req.BaseRequest.URL = "users/current"
	}

	if req.User != nil {
		req.BaseRequest.URL = "users/{user_id}"
	}

	resp := NewUserResponse()

	err := c.Send(req, resp)
	return resp, err
}

func NewUserResponse() *UserResponse {
	return &UserResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

type UserResponse struct {
	*response.BaseResponse
	Data User `json:"data"`
}
