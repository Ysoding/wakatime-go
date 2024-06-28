package v1

import (
	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type MachineNamesRequest struct {
	*request.BaseRequest
	User    *string `json:"user,omitempty"`
	Current *bool   `json:"-"`
}

func NewMachineNamesRequest() *MachineNamesRequest {
	return &MachineNamesRequest{
		BaseRequest: &request.BaseRequest{
			Version: "v1",
			Method:  sdk.MethodGet,
		},
		Current: sdk.Bool(true),
	}
}

func (c *Client) MachineNames(req *MachineNamesRequest) (*MachineNamesResponse, error) {
	if req == nil {
		req = NewMachineNamesRequest()
	}

	if req.Current != nil && *req.Current {
		req.BaseRequest.URL = "users/current/machine_names"
	}

	if req.User != nil {
		req.BaseRequest.URL = "users/{user}/machine_names"
	}

	resp := NewMachineNamesResponse()

	err := c.Send(req, resp)
	return resp, err
}

func NewMachineNamesResponse() *MachineNamesResponse {
	return &MachineNamesResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

type MachineNamesResponse struct {
	*response.BaseResponse
	Data []MachineData `json:"data"`
}
