package github

type PullRequest struct {
	URL      string `json:"url"`
	HtmlURL  string `json:"html_url"`
	DiffURL  string `json:"diff_url"`
	PatchURL string `json:"patch_url"`
}
