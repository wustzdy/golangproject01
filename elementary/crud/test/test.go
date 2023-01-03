package main

import (
	"encoding/json"
	"fmt"
	"golangproject01/elementary/crud/core"
	"golangproject01/elementary/crud/models"
)

func main() {
	Init()
	//CreateTest()

	list, _ := SelectOne("张三")
	if result, err := json.Marshal(list); err == nil {
		fmt.Println("result:", string(result))
	}
}

func Init() {
	core.Init()
}

func CreateTest() {
	project := models.Project{
		Name:        "项目1",
		Description: "这是第一个项目描述",
	}
	core.DB.Create(&project)

	meta := models.Meta{
		UserName:  "张三",
		Email:     "714403855@qq.com",
		ProjectId: project.ID,
	}
	core.DB.Create(&meta)

}
func SelectOne(userName string) (list []models.ProjectMetas, err error) {

	var sql = "select p.id as ProjectId, p.name, p.description,m.user_name,m.email " +
		"from projects p, meta m " +
		"where m.user_name = ? " +
		"and p.id = m.project_id "
	err = core.DB.Raw(sql, userName).Scan(&list).Error
	return

}
