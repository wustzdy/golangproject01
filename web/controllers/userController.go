package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	models "golangproject01/web/models"
	"io/ioutil"
	"net/http"
)

type UserController struct {
	BaseController
}

type Article struct {
	Name       string `json:"name"` // 注意最终json的变量名是后面的！
	University string `json:"university"`
	College    string `json:"college"`
	Major      string `json:"major"`
	Class      string `json:"class"`
}

func (user UserController) Index(c *gin.Context) {
	name := c.DefaultQuery("name", "无参数")
	c.String(200, name)
}
func (user UserController) UserList(c *gin.Context) {
	//userList := make([]models.User, 2)
	var userList models.User
	fmt.Println("userList:", userList)
	models.DB.Find(&userList)
	c.JSON(200, gin.H{
		"result": userList,
	})
}

func (user UserController) Map(context *gin.Context) {
	context.JSON(200, gin.H{"user": "zhangSan", "secret": 12345})
}

func (user UserController) Json(context *gin.Context) {
	a := Article{"zhudayang", "武汉科技大学", "计算机科学与技术学院", "计算机科学与技术", "1205"}
	context.JSON(200, a)
}

func (user UserController) Test(c *gin.Context) {
	user.success(c)
}

func (user UserController) Add(c *gin.Context) {
	m := models.User{
		UserName: "张三李四",
		Age:      12,
		Email:    "wust@qq.com",
	}
	models.DB.Create(&m)
	c.String(200, "增加用户")
}
func (user UserController) query(c *gin.Context) {
	m := models.User{}
	models.DB.Find(&m)
	c.String(200, "查询用户")
}

func (user UserController) Edit(c *gin.Context) {
	user1 := models.User{Id: 2}
	models.DB.Create(&user1)
	user1.UserName = "武汉科技大学"
	models.DB.Save(&user1)
	user.success(c)
}

func (user UserController) Delete(c *gin.Context) {
	user1 := models.User{Id: 2}
	models.DB.Delete(&user1)
	user.success(c)
}
func (user UserController) Post(c *gin.Context) {
	bodyByts, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// 返回错误信息
		c.String(http.StatusBadRequest, err.Error())
		// 执行退出
		c.Abort()
	}
	// 返回的 code 和 对应的参数星系
	c.String(http.StatusOK, "%s \n", string(bodyByts))
}

type CreateAutoTaskReq struct {
	DataType            int         `json:"dataType" binding:"min=0,max=1"`
	SetVerID            string      `json:"setVerID"`
	SceneType           SceneType   `json:"sceneType" binding:"min=0,max=9"`
	SceneSource         SceneSource `json:"sceneSource"`
	TaskType            uint        `json:"taskType" binding:"min=1,max=1"`
	CustomDataConfigID  *string     `json:"customDataConfigID"`
	CustomModuleID      *int64      `json:"customModuleID"`      // 自定训练配置id
	CustomModuleConfigs *string     `json:"customModuleConfigs"` // 自定义训练配置超参yml
}
type SceneType int
type SceneSource int

const (
	// TaskTypeObjectDetection 物体检测
	TaskTypeObjectDetection SceneType = 0
	// TaskTypeImageClassification 图像分类
	TaskTypeImageClassification SceneType = 1
	// TaskTypeSingleDetection 单目标检测
	TaskTypeSingleDetection SceneType = 2
	// TaskTypeMultiDetection 多目标检测
	TaskTypeMultiDetection SceneType = 3
	// TaskTypeBinaryClassification 二分类
	TaskTypeBinaryClassification SceneType = 4
	// TaskTypeMultiClassification 多分类
	TaskTypeMultiClassification SceneType = 5
	// TaskTypeInstanceSegmentation 实例分割
	TaskTypeInstanceSegmentation SceneType = 6
	// TaskTypeSemanticSegmentation 语义分割
	TaskTypeSemanticSegmentation SceneType = 7
	// TaskTypeEnglishOCR 英文OCR
	TaskTypeEnglishOCR SceneType = 8
	// TaskTypeDetectionClassification
	TaskTypeDetectionClassification SceneType = 9
	// TaskTypeVideoClassification 视频分类
	TaskTypeVideoClassification SceneType = 10
	// TaskTypePointCloud 点云
	TaskTypePointCloud SceneType = 11
)

func (user UserController) TestJson(c *gin.Context) {
	var req CreateAutoTaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		return
	}
	resJson, _ := json.Marshal(req)
	fmt.Println("--CreateAutoTaskReq-req:", string(resJson))

	c.JSON(200, resJson)
}
