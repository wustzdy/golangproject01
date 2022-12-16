package main

import (
	"encoding/json"
	"fmt"
	"golangproject01/gormRelations/hasTo/B/core"
	"golangproject01/gormRelations/hasTo/B/models"
	"gorm.io/gorm"
)

func main() {
	Init()
	createTest()
	selectTest()
	selectSql()
}

func Init() {
	core.Init()
}
func createTest() {

	u := models.User{
		CreditCard: models.CreditCard{
			Number: "123",
		},
		Model: gorm.Model{
			ID: 9999,
		},
	}
	core.DB.Create(&u)

}
func selectTest() {
	var user []models.User
	core.DB.Model(models.User{}).Preload("CreditCard").Find(&user)
	json, _ := json.Marshal(user)
	fmt.Println("user:", string(json))

}

// 原生sql
func selectSql() (list []models.User, err error) {
	var sql = "select u.* " +
		"from users u "
	err = core.DB.Raw(sql).Scan(&list).Error
	result, _ := json.Marshal(list)
	fmt.Println("result:", string(result))
	return

}

/*func FindDeviceTemplateByDeviceId(ctx context.Context, db *gorm.DB, deviceId int64) (list []DeviceTemplate, err error) {
	var sql = "select a.device_id, a.template_id, b.name template_name " +
		"from device_template a, template b " +
		"where a.device_id = ? " +
		"and a.template_id = b.id "
	err = db.Raw(sql, deviceId).Scan(&list).Error
	return
}*/
