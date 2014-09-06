package github

type Client struct {
	BaseURL string `json:"baseUrl"`
	Token   string `json:"token"`
}

type User struct {
	Login             string `json:"login"`
	ID                uint32 `json:"id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HtmlURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type Label struct {
	URL   string `json:"url"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Milestone struct {
	URL          string `json:"url"`
	Number       uint32 `json:"number"`
	State        string `json:"state"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Creator      User   `json:"creator"`
	OpenIssues   uint32 `json:"open_issues"`
	ClosedIssues uint32 `json:"closed_issues"`
	Created      string `json:"created_at"`
	Updated      string `json:"updated_at"`
	DueOn        string `json:"due_on"`
}

type PullRequest struct {
	URL      string `json:"url"`
	HtmlURL  string `json:"html_url"`
	DiffURL  string `json:"diff_url"`
	PatchURL string `json:"patch_url"`
}

type Issue struct {
	URL         string      `json:"url"`
	HtmlURL     string      `json:"html_url"`
	Number      uint32      `json:"number"`
	State       string      `json:"state"`
	Title       string      `json:"title"`
	Body        string      `json:"body"`
	User        User        `json:"user"`
	Labels      []Label     `json:"labels"`
	Assignee    User        `json:"assignee"`
	Milestone   Milestone   `json:"milestone"`
	Comments    uint32      `json:"comments"`
	PullRequest PullRequest `json:"pull_request"`
	Closed      string      `json:"closed_at"`
	Created     string      `json:"created_at"`
	Updated     string      `json:"updated_at"`
}

type TokenRequest struct {
	Scopes       []string `json:"scopes,omitempty"`
	Note         string   `json:"note,omitempty"`
	NoteURL      string   `json:"note_url,omitempty"`
	ClientID     string   `json:"client_id,omitempty"`
	ClientSecret string   `json:"client_secret,omitempty"`
}

type TokenResponse struct {
	Id     uint32   `json:"id"`
	URL    string   `json:"url"`
	Scopes []string `json:"scopes"`
	Token  string   `json:"token"`
	App    struct {
		URL      string `json:"url"`
		Name     string `json:"name"`
		ClientID string `json:"name"`
	} `json:"app"`
	Note    string `json:"note"`
	NoteURL string `json:"note_url"`
	Updated string `json:"updated_at"`
	Created string `json:"created_at"`
}

type ErrorResponse struct {
	Message          string `json:"message"`
	DocumentationURL string `json:"documentation_url"`
}

func (err *ErrorResponse) Error() string {
	return err.Message
}
