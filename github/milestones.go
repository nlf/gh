package github

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

type MilestoneSlice []Milestone

func (client Client) GetMilestones(repo string, state string) (MilestoneSlice, error) {
	query := make(map[string]string)
	if state != "" {
		query["state"] = state
	}

	milestones := MilestoneSlice{}
	err := client.Request("GET", "/repos/"+repo+"/milestones", query, nil, &milestones)
	if err != nil {
		return nil, err
	}

	return milestones, nil
}
