package v1

import "time"

type InsightsData struct {
	InsightType        string      `json:"insight_type"`
	InsightData        interface{} `json:"data"` // This can be either an object or array, so using interface{}
	Range              string      `json:"range"`
	HumanReadableRange string      `json:"human_readable_range"`
	Status             string      `json:"status"`
	IsIncludingToday   bool        `json:"is_including_today"`
	IsUpToDate         bool        `json:"is_up_to_date"`
	PercentCalculated  int         `json:"percent_calculated"`
	Start              string      `json:"start"`
	End                string      `json:"end"`
	Timezone           string      `json:"timezone"`
	Timeout            int         `json:"timeout"`
	WritesOnly         bool        `json:"writes_only"`
	UserID             string      `json:"user_id"`
	CreatedAt          string      `json:"created_at"`
	ModifiedAt         string      `json:"modified_at"`
}

type Leader struct {
	Rank         int `json:"rank"`
	RunningTotal struct {
		TotalSeconds              float64 `json:"total_seconds"`
		HumanReadableTotal        string  `json:"human_readable_total"`
		DailyAverage              float64 `json:"daily_average"`
		HumanReadableDailyAverage string  `json:"human_readable_daily_average"`
		Languages                 []struct {
			Name         string  `json:"name"`
			TotalSeconds float64 `json:"total_seconds"`
		} `json:"languages"`
	} `json:"running_total"`
	User User `json:"user"`
}

type MachineData struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Value      string    `json:"value"`
	IP         string    `json:"ip"`
	LastSeenAt time.Time `json:"last_seen_at"`
	Timezone   string    `json:"timezone"`
	CreatedAt  time.Time `json:"created_at"`
}

type IPData struct {
	IPs            IPAddresses       `json:"ips"`
	IPDescriptions map[string]string `json:"ip_descriptions"`
	LastModifiedAt string            `json:"last_modified_at"`
}

type IPAddresses struct {
	API     IPVersion `json:"api"`
	Website IPVersion `json:"website"`
	Worker  IPVersion `json:"worker"`
}

type IPVersion struct {
	V4 []string `json:"v4"`
	V6 []string `json:"v6"`
}

type OrgDashboardMemberDurationsData struct {
	Project  string  `json:"project"`
	Time     float64 `json:"time"`
	Duration float64 `json:"duration"`
}

type OrgDashboardMemberSummariesData struct {
	GrandTotal       GrandTotal        `json:"grand_total"`
	Projects         []Project         `json:"projects"`
	Languages        []Language        `json:"languages"`
	Editors          []Editor          `json:"editors"`
	OperatingSystems []OperatingSystem `json:"operating_systems"`
	Branches         []Branch          `json:"branches,omitempty"` // included only when project url parameter used
	Entities         []Entity          `json:"entities,omitempty"` // included only when project url parameter used
	Range            DayRange          `json:"range"`
}

type DayRange struct {
	Date     string `json:"date"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Text     string `json:"text"`
	Timezone string `json:"timezone"`
}

type DashboardMember struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	FullName   string `json:"full_name"`
	IsViewOnly bool   `json:"is_view_only"`
	Photo      string `json:"photo"`
	Username   string `json:"username"`
}

type Organization struct {
	ID                                       string    `json:"id"`
	Name                                     string    `json:"name"`
	DefaultProjectPrivacy                    string    `json:"default_project_privacy"`
	InvitedPeopleCount                       int       `json:"invited_people_count"`
	InvitedPeopleCountHumanReadable          string    `json:"invited_people_count_human_readable"`
	IsDurationVisible                        bool      `json:"is_duration_visible"`
	PeopleCount                              int       `json:"people_count"`
	PeopleCountHumanReadable                 string    `json:"people_count_human_readable"`
	Timeout                                  int       `json:"timeout"`
	Timezone                                 int       `json:"timezone"`
	WritesOnly                               int       `json:"writes_only"`
	CanCurrentUserListDashboards             bool      `json:"can_current_user_list_dashboards"`
	CanCurrentUserCreateDashboards           bool      `json:"can_current_user_create_dashboards"`
	CanCurrentUserDisplayCodingOnDashboards  bool      `json:"can_current_user_display_coding_on_dashboards"`
	CanCurrentUserViewAllDashboards          bool      `json:"can_current_user_view_all_dashboards"`
	CanCurrentUserAddPeopleToDashboards      bool      `json:"can_current_user_add_people_to_dashboards"`
	CanCurrentUserRemovePeopleFromDashboards bool      `json:"can_current_user_remove_people_from_dashboards"`
	CanCurrentUserEditAndDeleteDashboards    bool      `json:"can_current_user_edit_and_delete_dashboards"`
	CanCurrentUserAddPeopleToOrg             bool      `json:"can_current_user_add_people_to_org"`
	CanCurrentUserRemovePeopleFromOrg        bool      `json:"can_current_user_remove_people_from_org"`
	CanCurrentUserManageGroups               bool      `json:"can_current_user_manage_groups"`
	CanCurrentUserViewAuditLog               bool      `json:"can_current_user_view_audit_log"`
	CanCurrentUserEditOrg                    bool      `json:"can_current_user_edit_org"`
	CanCurrentUserManageBilling              bool      `json:"can_current_user_manage_billing"`
	CanCurrentUserDeleteOrg                  bool      `json:"can_current_user_delete_org"`
	CreatedAt                                time.Time `json:"created_at"`
	ModifiedAt                               time.Time `json:"modified_at"`
}

type PrivateLeaderboard struct {
	CanDelete                 bool      `json:"can_delete"`
	CanEdit                   bool      `json:"can_edit"`
	CreatedAt                 time.Time `json:"created_at"`
	HasAvailableSeat          bool      `json:"has_available_seat"`
	ID                        string    `json:"id"`
	MembersCount              int       `json:"members_count"`
	MembersWithTimezonesCount int       `json:"members_with_timezones_count"`
	ModifiedAt                time.Time `json:"modified_at"`
	Name                      string    `json:"name"`
	TimeRange                 string    `json:"time_range"`
}

type LeaderboardMember struct {
	Rank int             `json:"rank"`
	Page int             `json:"page"`
	User LeaderboardUser `json:"user"`
}

type LeaderboardUser struct {
	ID                   string `json:"id"`
	Email                string `json:"email"`
	Username             string `json:"username"`
	FullName             string `json:"full_name"`
	DisplayName          string `json:"display_name"`
	Website              string `json:"website"`
	HumanReadableWebsite string `json:"human_readable_website"`
	IsHireable           bool   `json:"is_hireable"`
	City                 City   `json:"city"`
	IsEmailPublic        bool   `json:"is_email_public"`
	PhotoPublic          bool   `json:"photo_public"`
}

type LeaderboardRange struct {
	StartDate string `json:"start_date"`
	StartText string `json:"start_text"`
	EndDate   string `json:"end_date"`
	EndText   string `json:"end_text"`
	Name      string `json:"name"`
	Text      string `json:"text"`
}

type ProgramLanguageDetail struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Color      string    `json:"color"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

type ProjectDetail struct {
	ID                            string    `json:"id"`
	Name                          string    `json:"name"`
	Repository                    string    `json:"repository"`
	Badge                         string    `json:"badge"`
	Color                         string    `json:"color"`
	Clients                       []string  `json:"clients"`
	HasPublicURL                  bool      `json:"has_public_url"`
	HumanReadableLastHeartbeatAt  string    `json:"human_readable_last_heartbeat_at"`
	LastHeartbeatAt               string    `json:"last_heartbeat_at"`
	HumanReadableFirstHeartbeatAt string    `json:"human_readable_first_heartbeat_at"`
	FirstHeartbeatAt              string    `json:"first_heartbeat_at"`
	URL                           string    `json:"url"`
	URLEncodedName                string    `json:"urlencoded_name"`
	CreatedAt                     time.Time `json:"created_at"`
}

type StatsData struct {
	TotalSeconds                       float64           `json:"total_seconds"`
	TotalSecondsIncludingOtherLanguage float64           `json:"total_seconds_including_other_language"`
	HumanReadableTotal                 string            `json:"human_readable_total"`
	HumanReadableTotalIncludingOther   string            `json:"human_readable_total_including_other_language"`
	DailyAverage                       float64           `json:"daily_average"`
	DailyAverageIncludingOtherLanguage float64           `json:"daily_average_including_other_language"`
	HumanReadableDailyAverage          string            `json:"human_readable_daily_average"`
	HumanReadableDailyAverageIncluding string            `json:"human_readable_daily_average_including_other_language"`
	Categories                         []Category        `json:"categories"`
	Projects                           []Project         `json:"projects"`
	Languages                          []Language        `json:"languages"`
	Editors                            []Editor          `json:"editors"`
	OperatingSystems                   []OperatingSystem `json:"operating_systems"`
	Dependencies                       []Dependency      `json:"dependencies"`
	Machines                           []Machine         `json:"machines"`
	BestDay                            BestDay           `json:"best_day"`
	Range                              string            `json:"range"`
	HumanReadableRange                 string            `json:"human_readable_range"`
	Holidays                           int               `json:"holidays"`
	DaysIncludingHolidays              int               `json:"days_including_holidays"`
	DaysMinusHolidays                  int               `json:"days_minus_holidays"`
	Status                             string            `json:"status"`
	PercentCalculated                  int               `json:"percent_calculated"`
	IsAlreadyUpdating                  bool              `json:"is_already_updating"`
	IsCodingActivityVisible            bool              `json:"is_coding_activity_visible"`
	IsOtherUsageVisible                bool              `json:"is_other_usage_visible"`
	IsStuck                            bool              `json:"is_stuck"`
	IsIncludingToday                   bool              `json:"is_including_today"`
	IsUpToDate                         bool              `json:"is_up_to_date"`
	Start                              string            `json:"start"`
	End                                string            `json:"end"`
	Timezone                           string            `json:"timezone"`
	Timeout                            int               `json:"timeout"`
	WritesOnly                         bool              `json:"writes_only"`
	UserID                             string            `json:"user_id"`
	Username                           string            `json:"username"`
	CreatedAt                          time.Time         `json:"created_at"`
	ModifiedAt                         time.Time         `json:"modified_at"`
}

type BestDay struct {
	Date         string  `json:"date"`
	Text         string  `json:"text"`
	TotalSeconds float64 `json:"total_seconds"`
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

type StatsAggregatedData struct {
	Categories       []CategoryStats        `json:"categories"`
	DailyAverage     DailyAverageStats      `json:"daily_average"`
	Editors          []EditorStats          `json:"editors"`
	Languages        []LanguageStats        `json:"languages"`
	OperatingSystems []OperatingSystemStats `json:"operating_systems"`
	Total            TotalStats             `json:"total"`
	Range            SummaryRangeStats      `json:"range"`
	Timeout          int                    `json:"timeout"`
	WritesOnly       bool                   `json:"writes_only"`
}

type AverageData struct {
	Seconds float64 `json:"seconds"`
	Text    string  `json:"text"`
}

type CountData struct {
	Text string `json:"text"`
}

type StatsBase struct {
	Average AverageData `json:"average"`
	Count   CountData   `json:"count"`
	Max     AverageData `json:"max"`
	Median  AverageData `json:"median"`
	Sum     AverageData `json:"sum"`
}

type CategoryStats struct {
	Name       string `json:"name"`
	IsVerified bool   `json:"is_verified"`
	StatsBase
}

type DailyAverageStats struct {
	StatsBase
}

type EditorStats struct {
	Name       string `json:"name"`
	IsVerified bool   `json:"is_verified"`
	StatsBase
}

type LanguageStats struct {
	Name       string `json:"name"`
	IsVerified bool   `json:"is_verified"`
	StatsBase
}

type OperatingSystemStats struct {
	Name       string `json:"name"`
	IsVerified bool   `json:"is_verified"`
	StatsBase
}

type TotalStats struct {
	StatsBase
}

type SummaryRangeStats struct {
	EndDate   string `json:"end_date"`
	EndText   string `json:"end_text"`
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	StartText string `json:"start_text"`
	Text      string `json:"text"`
}

type SummaryData struct {
	GrandTotal       GrandTotal        `json:"grand_total"`
	Categories       []Category        `json:"categories"`
	Projects         []Project         `json:"projects"`
	Languages        []Language        `json:"languages"`
	Editors          []Editor          `json:"editors"`
	OperatingSystems []OperatingSystem `json:"operating_systems"`
	Dependencies     []Dependency      `json:"dependencies"`
	Machines         []Machine         `json:"machines"`
	Branches         []Branch          `json:"branches,omitempty"`
	Entities         []Entity          `json:"entities,omitempty"`
	Range            SummaryRange      `json:"range"`
}

type GrandTotal struct {
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Text         string  `json:"text"`
	TotalSeconds float64 `json:"total_seconds"`
}

type SummaryBase struct {
	Name         string  `json:"name"`
	TotalSeconds float64 `json:"total_seconds"`
	Percent      float64 `json:"percent"`
	Digital      string  `json:"digital"`
	Text         string  `json:"text"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Seconds      int     `json:"seconds"`
}

type Category SummaryBase
type Project SummaryBase
type Language SummaryBase
type Editor SummaryBase
type OperatingSystem SummaryBase
type Dependency SummaryBase
type Machine struct {
	SummaryBase
	MachineNameID string `json:"machine_name_id"`
}

type Branch SummaryBase

type Entity SummaryBase

type SummaryRange struct {
	Date     string `json:"date"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Text     string `json:"text"`
	Timezone string `json:"timezone"`
}

type CumulativeTotal struct {
	Seconds float64 `json:"seconds"`
	Text    string  `json:"text"`
	Decimal string  `json:"decimal"`
	Digital string  `json:"digital"`
}

type DailyAverage struct {
	Holidays                  int     `json:"holidays"`
	DaysIncludingHolidays     int     `json:"days_including_holidays"`
	DaysMinusHolidays         int     `json:"days_minus_holidays"`
	Seconds                   float64 `json:"seconds"`
	Text                      string  `json:"text"`
	SecondsIncludingOtherLang float64 `json:"seconds_including_other_language"`
	TextIncludingOtherLang    string  `json:"text_including_other_language"`
}

type UserAgent struct {
	ID                 string    `json:"id"`
	Value              string    `json:"value"`
	Editor             string    `json:"editor"`
	Version            string    `json:"version"`
	OS                 string    `json:"os"`
	LastSeenAt         time.Time `json:"last_seen_at"`
	IsBrowserExtension bool      `json:"is_browser_extension"`
	IsDesktopApp       bool      `json:"is_desktop_app"`
	CreatedAt          time.Time `json:"created_at"`
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
