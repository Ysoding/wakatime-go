package v1

import (
	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type AllTimeRequest struct {
	*request.BaseRequest
	Project *string `json:"project"`
	User    *string `json:"user"`
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
	} else {
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

type TimeData struct {
	DailyAverage      float64   `json:"daily_average"`
	Decimal           string    `json:"decimal"`
	Digital           string    `json:"digital"`
	IsUpToDate        bool      `json:"is_up_to_date"`
	PercentCalculated int       `json:"percent_calculated"`
	Range             TimeRange `json:"range"`
	Text              string    `json:"text"`
	Timeout           int       `json:"timeout"`
	TotalSeconds      float64   `json:"total_seconds"`
}

type TimeRange struct {
	End       string `json:"end"`
	EndDate   string `json:"end_date"`
	EndText   string `json:"end_text"`
	Start     string `json:"start"`
	StartDate string `json:"start_date"`
	StartText string `json:"start_text"`
	Timezone  string `json:"timezone"`
}
