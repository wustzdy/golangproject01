package main

import (
	"golangproject01/record/config"
	"golangproject01/record/errorc"
	"golangproject01/record/workspace"
)

func main() {

}

/*
@GetMapping(BASE_PATH + "/id:{workspaceId}/projects")
List<Project> getWorkspaceProjectByAccountId(@PathVariable("workspaceId") String workspaceId,
@RequestParam(value = "accountId", required = false) String accountId,
@RequestParam(value = "projectIds", required = false) List<String> projectIds);
*/
func GetUserProjectIDMap(spaceID, userID string) (map[uint]bool, error) {
	prjs, err := config.WSClient.ListProjects(&workspace.ListProjectsInput{
		WorkspaceID: spaceID,
		AccountID:   &userID,
	})
	if err != nil {
		return nil, err
	}

	ret := map[uint]bool{}
	for _, prj := range prjs {
		id, err := hashid(prj)
		if err != nil {
			return nil, errorc.Wrap(err)
		}
		ret[uint(id)] = true
	}
	return ret, nil
}

func hashid(prj workspace.ProjectInfo) (int64, error) {
	return 0, nil
}
