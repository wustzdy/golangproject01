package workspace

import (
	"encoding/json"
	"fmt"
	"golangproject01/record/errorc"
	"gopkg.in/resty.v1"
)

var (
	NotFoundErr = fmt.Errorf("resource not found")
)

type WSClient interface {
	ListProjects(input *ListProjectsInput) ([]ProjectInfo, error)
}

type wSClient struct {
	Host            string
	AdditionHeaders map[string]string
	client          *resty.Client
}

func (c *wSClient) ListProjects(input *ListProjectsInput) ([]ProjectInfo, error) {
	req := c.client.R().
		SetPathParams(map[string]string{"workspaceID": input.WorkspaceID})
	if input.AccountID != nil {
		req.SetQueryParam("accountId", *input.AccountID)
	}
	if len(input.ProjectIDs) > 0 {
		for _, id := range input.ProjectIDs {
			req.QueryParam.Add("projectIds", id)
		}
	}

	rsp, err := req.Get(fmt.Sprintf("%s/%s", c.Host, "workspace/workspaces/id:{workspaceID}/projects"))
	if err != nil {
		return nil, errorc.AddField(err, "id", input.WorkspaceID)
	}

	projects := []ProjectInfo{}
	if err = json.Unmarshal(rsp.Body(), &projects); err != nil {
		return nil, errorc.AddField(err, "body", string(rsp.Body())).AddField("id", input.WorkspaceID)
	}
	return projects, nil
}
