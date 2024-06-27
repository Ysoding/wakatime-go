package v1

import (
	"time"

	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type UserRequest struct {
	*request.BaseRequest
	User    *string `json:"user_id"`
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

func (c *Client) User(req *UserRequest) (*UsersResponse, error) {
	if req == nil {
		req = NewUserRequest()
	}

	if req.Current != nil && *req.Current {
		req.BaseRequest.URL = "users/current"
	}

	if req.User != nil {
		req.BaseRequest.URL = "users/{user_id}"
	}

	resp := NewUsersResponse()

	err := c.Send(req, resp)
	return resp, err
}

func NewUsersResponse() *UsersResponse {
	return &UsersResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

type UsersResponse struct {
	*response.BaseResponse
	Data User `json:"data"`
}

type User struct {
	ID                   string    `json:"id"`
	HasPremiumFeatures   bool      `json:"has_premium_features"`
	DisplayName          string    `json:"display_name"`
	FullName             string    `json:"full_name"`
	Email                string    `json:"email"`
	Photo                string    `json:"photo"`
	IsEmailPublic        bool      `json:"is_email_public"`
	IsEmailConfirmed     bool      `json:"is_email_confirmed"`
	PublicEmail          *string   `json:"public_email"` // Nullable
	PhotoPublic          bool      `json:"photo_public"`
	Timezone             string    `json:"timezone"`
	LastHeartbeatAt      time.Time `json:"last_heartbeat_at"`
	LastPlugin           string    `json:"last_plugin"`
	LastPluginName       string    `json:"last_plugin_name"`
	LastProject          string    `json:"last_project"`
	LastBranch           string    `json:"last_branch"`
	Plan                 string    `json:"plan"`
	Username             string    `json:"username"`
	Website              string    `json:"website"`
	HumanReadableWebsite string    `json:"human_readable_website"`
	WonderfulDevUsername string    `json:"wonderfuldev_username"`
	GitHubUsername       string    `json:"github_username"`
	TwitterUsername      string    `json:"twitter_username"`
	LinkedInUsername     string    `json:"linkedin_username"`
	City                 City      `json:"city"`
	LoggedTimePublic     bool      `json:"logged_time_public"`
	LanguagesUsedPublic  bool      `json:"languages_used_public"`
	IsHireable           bool      `json:"is_hireable"`
	CreatedAt            time.Time `json:"created_at"`
	ModifiedAt           time.Time `json:"modified_at"`
}

type City struct {
	CountryCode string `json:"country_code"`
	Name        string `json:"name"`
	State       string `json:"state"`
	Title       string `json:"title"`
}
