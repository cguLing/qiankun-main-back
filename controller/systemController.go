package controller

import (
	"bus-backend-go/model"
	"bus-backend-go/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

type SystemController struct {
	log           *logrus.Logger
	systemService service.SystemService
}

func NewSystemController(log *logrus.Logger) *SystemController {
	return &SystemController{log: log, systemService: service.NewSystemService(log)}
}

/**
 * @api GET /api/v1/system/userAdmin 查询用户是否为超级管理员
 * @apiGroup system

 * @apiRequest json
 * @apiHeader Authorization Bearer <token>
 * @apiParam ldap string 查询是否为超管的用户ldap
 * @apiExample json
 * {"ldap": "chenjh03"}

 * @apiSuccess 200 OK
 * @apiParam code int 成功代码
 * @apiParam msg string 成功描述
 * @apiParam data string 成功返回：1是0不是
 * @apiExample json
 * {"code":200, "msg":"success", "data": 1}
 * {"code":200, "msg":"success", "data": 0}

 * @apiError 500
 * @apiParam code int 错误代码
 * @apiParam msg string 错误描述
 * @apiParam data string 错误信息
 * @apiExample json
 * {"code":500, "msg":"error desc", "data": "error msg"}
 */
func (controller SystemController) GetUserAdmin(c *gin.Context) {
	// TODO 如果没有传入ldap，则默认查询发起请求的人自己，需要获取token
	userAdmin := model.UserAdmin{}
	err := c.BindQuery(&userAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "get user admin error: missing ldap",
			"data": err,
		})
		return
	}
	// 进行查询
	res, err := controller.systemService.FindUserAdminByLdap(userAdmin.Ldap)
	if err != nil {
		// 如果是因为没查到结果返回的错误
		if err == gorm.ErrRecordNotFound{
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "success",
				"data": 0,
			})
			return
		}
		// 如果是其他错误
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "get user admin error",
			"data": err,
		})
		return
	}
	// 正常查询到结果
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": res,
	})
	return
}


/**
 * @api POST /api/v1/system/userAdmin 新增超级管理员
 * @apiGroup system

 * @apiRequest json
 * @apiHeader Authorization Bearer <token>
 * @apiParam ldap string 查询是否为超管的用户ldap
 * @apiExample json
 * {"ldap": "chenjh03"}

 * @apiSuccess 200 OK
 * @apiParam code int 成功代码
 * @apiParam msg string 成功描述
 * @apiParam data string 成功信息
 * @apiExample json
 * {"code":200, "msg":"success", "data": "success"}

 * @apiError 500
 * @apiParam code int 错误代码
 * @apiParam msg string 错误描述
 * @apiParam data string 错误信息
 * @apiExample json
 * {"code":500, "msg":"error desc", "data": "error msg"}
 */
func (controller SystemController) PostUserAdmin(c *gin.Context) {
	var admin = model.UserAdmin{}
	if err := c.BindJSON(&admin); err != nil || admin.Ldap == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "get user admin error",
			"data": err,
		})
		return
	}
	// TODO 从headers获取操作者
	admin.Creator = "root"
	// 去添加
	_, err := controller.systemService.AddUserAdmin(admin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "add user admin error",
			"data": err,
		})
		return
	}
	// 正常返回
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": "success",
	})
	return
}


func (controller SystemController) PutUserAdmin(c *gin.Context) {
	var ua = model.UserAdmin{}
	if err := c.BindJSON(&ua); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "get user list post data error",
			"data": err,
		})
		return
	}
	res, err := controller.systemService.UpdateUserAdmin(ua)
	if err != nil {
		controller.log.Errorf("put user list err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "put user list error",
			"data": err.Error(),
		})
		return
	}
	// 正常查询到结果
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": res,
	})
	return
}