package v1

import (
	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type InsightsRequest struct {
	*request.BaseRequest
	User        *string `json:"user,omitempty"`
	Current     *bool   `json:"-"`
	InsightType string  `json:"insight_type"`
	Range       string  `json:"range"`

	Timeout    *int  `json:"timeout,omitempty"`
	WritesOnly *bool `json:"writes_only,omitempty"`
	Weekday    *int  `json:"weekday,omitempty"`
}

func NewInsightsRequest(insightType, rng string) *InsightsRequest {
	return &InsightsRequest{
		BaseRequest: &request.BaseRequest{
			Version: "v1",
			Method:  sdk.MethodGet,
		},
		Current:     sdk.Bool(true),
		Range:       rng,
		InsightType: insightType,
	}
}

func (c *Client) Insights(req *InsightsRequest) (*InsightsResponse, error) {
	if req == nil {
		req = NewInsightsRequest("projects", "last_7_days")
	}

	if req.Current != nil && *req.Current {
		req.BaseRequest.URL = "users/current"
	}

	if req.User != nil {
		req.BaseRequest.URL = "users/{user}"
	}

	req.BaseRequest.URL += "/insights/{insight_type}/{range}"

	resp := NewInsightsResponse()

	err := c.Send(req, resp)
	return resp, err
}

func NewInsightsResponse() *InsightsResponse {
	return &InsightsResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

type InsightsResponse struct {
	*response.BaseResponse
	Data TimeData `json:"data"`
}
