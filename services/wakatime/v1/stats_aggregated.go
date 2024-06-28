package v1

import (
	"fmt"
	"time"

	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type StatsAggregatedRequest struct {
	*request.BaseRequest
	Range *string `json:"range"`
}

func NewStatsAggregatedRequest(rng string) *StatsAggregatedRequest {
	return &StatsAggregatedRequest{
		BaseRequest: &request.BaseRequest{
			Version: "v1",
			Method:  sdk.MethodGet,
			URL:     "stats/{range}",
		},
		Range: sdk.String(rng),
	}
}

func (c *Client) StatsAggregated(req *StatsAggregatedRequest) (*StatsAggregatedResponse, error) {
	if req == nil {
		req = NewStatsAggregatedRequest(fmt.Sprintf("%d", time.Now().AddDate(-1, 0, 0).Year()))
	}

	resp := NewStatsAggregatedResponse()

	err := c.Send(req, resp)
	return resp, err
}

func NewStatsAggregatedResponse() *StatsAggregatedResponse {
	return &StatsAggregatedResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

type StatsAggregatedResponse struct {
	*response.BaseResponse
	Data StatsAggregatedData `json:"data"`
}
