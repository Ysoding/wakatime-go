package v1

import (
	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type ProgramLanguagesRequest struct {
	*request.BaseRequest
}

func NewProgramLanguagesRequest() *ProgramLanguagesRequest {
	return &ProgramLanguagesRequest{
		BaseRequest: &request.BaseRequest{
			Version: "v1",
			Method:  sdk.MethodGet,
			URL:     "program_languages",
		},
	}
}

func (c *Client) ProgramLanguages(req *ProgramLanguagesRequest) (*ProgramLanguagesResponse, error) {
	if req == nil {
		req = NewProgramLanguagesRequest()
	}

	resp := NewProgramLanguagesResponse()

	err := c.Send(req, resp)
	return resp, err
}

func NewProgramLanguagesResponse() *ProgramLanguagesResponse {
	return &ProgramLanguagesResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

type ProgramLanguagesResponse struct {
	*response.BaseResponse
	Data       []ProgramLanguageDetail `json:"data"`
	Total      int                     `json:"total"`
	TotalPages int                     `json:"total_pages"`
}
