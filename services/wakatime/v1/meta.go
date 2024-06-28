package v1

import (
	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type MetaRequest struct {
	*request.BaseRequest
}

func NewMetaRequest() *MetaRequest {
	return &MetaRequest{
		BaseRequest: &request.BaseRequest{
			Version: "v1",
			Method:  sdk.MethodGet,
			URL:     "meta",
		},
	}
}

func (c *Client) Meta(req *MetaRequest) (*MetaResponse, error) {
	if req == nil {
		req = NewMetaRequest()
	}

	resp := NewMetaResponse()

	err := c.Send(req, resp)
	return resp, err
}

func NewMetaResponse() *MetaResponse {
	return &MetaResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

type MetaResponse struct {
	*response.BaseResponse
	Data IPData `json:"data"`
}
