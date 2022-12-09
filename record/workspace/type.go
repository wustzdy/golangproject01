package workspace

type ListProjectsInput struct {
	WorkspaceID string   `json:"-"`
	AccountID   *string  `json:"-"`
	ProjectIDs  []string `json:"-"`
}

type ProjectInfo struct {
	ID              string   `json:"id"`
	AccountID       string   `json:"accountId"`
	WorkspaceID     string   `json:"workspaceId"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	ResponsibleUser UserInfo `json:"responsibleUser"`
}

type UserInfo struct {
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
}
