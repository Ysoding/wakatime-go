package v1

import (
	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type LeadersRequest struct {
	*request.BaseRequest
	Language    *string `json:"language,omitempty"`
	IsHireable  *string `json:"is_hireable,omitempty"`
	CountryCode *string `json:"country_code,omitempty"`
	Page        *int    `json:"page,omitempty"`
}

func NewLeadersRequest() *LeadersRequest {
	return &LeadersRequest{
		BaseRequest: &request.BaseRequest{
			Version: "v1",
			Method:  sdk.MethodGet,
			URL:     "leaders",
		},
	}
}

func (c *Client) Leaders(req *LeadersRequest) (*LeadersResponse, error) {
	if req == nil {
		req = NewLeadersRequest()
	}

	resp := NewLeadersResponse()

	err := c.Send(req, resp)
	return resp, err
}

func NewLeadersResponse() *LeadersResponse {
	return &LeadersResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

type LeadersResponse struct {
	*response.BaseResponse

	CurrentUser struct {
		Rank int  `json:"rank"`
		Page int  `json:"page"`
		User User `json:"user"`
	} `json:"current_user"`

	Data []Leader `json:"data"`

	Page       int `json:"page"`
	TotalPages int `json:"total_pages"`
	Range      struct {
		StartDate string `json:"start_date"`
		StartText string `json:"start_text"`
		EndDate   string `json:"end_date"`
		EndText   string `json:"end_text"`
		Name      string `json:"name"`
		Text      string `json:"text"`
	} `json:"range"`
	Language    string `json:"language"`
	IsHireable  bool   `json:"is_hireable"`
	CountryCode string `json:"country_code"`
	ModifiedAt  string `json:"modified_at"`
	Timeout     int    `json:"timeout"`
	WritesOnly  bool   `json:"writes_only"`
}
