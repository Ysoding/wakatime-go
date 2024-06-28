package v1

type SummariesRangeOption string

const (
	SummariesRangeToday                  SummariesRangeOption = "Today"
	SummariesRangeYesterday              SummariesRangeOption = "Yesterday"
	SummariesRangeLast7Days              SummariesRangeOption = "Last 7 Days"
	SummariesRangeLast7DaysFromYesterday SummariesRangeOption = "Last 7 Days from Yesterday"
	SummariesRangeLast14Days             SummariesRangeOption = "Last 14 Days"
	SummariesRangeLast30Days             SummariesRangeOption = "Last 30 Days"
	SummariesRangeThisWeek               SummariesRangeOption = "This Week"
	SummariesRangeLastWeek               SummariesRangeOption = "Last Week"
	SummariesRangeThisMonth              SummariesRangeOption = "This Month"
	SummariesRangeLastMonth              SummariesRangeOption = "Last Month"
)
