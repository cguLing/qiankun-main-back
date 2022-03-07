package controller

import (
	"bus-backend-go/model"
	"bus-backend-go/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ServiceController struct {
	log           *logrus.Logger
	stService service.StService
}

func NewServiceController(log *logrus.Logger) *ServiceController {
	return &ServiceController{log: log, stService: service.NewStService(log)}
}


func (controller ServiceController) GetServiceType(c *gin.Context) {
	// 进行查询
	res, err := controller.stService.FindServiceType()
	if err != nil {
		controller.log.Errorf("get service type err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "get service type error",
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

func (controller ServiceController) PostServiceType(c *gin.Context) {
	var st = model.ServiceType{}
	if err := c.BindJSON(&st); err != nil || st.ClassName == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "get service type post data error",
			"data": err,
		})
		return
	}
	res, err := controller.stService.AddServiceType(st)
	if err != nil {
		controller.log.Errorf("post service type err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "post service type error",
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

func (controller ServiceController) PutServiceType(c *gin.Context) {
	var st = model.ServiceType{}
	if err := c.BindJSON(&st); err != nil || st.ClassName == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "get service type post data error",
			"data": err,
		})
		return
	}
	res, err := controller.stService.UpdateServiceType(st)
	if err != nil {
		controller.log.Errorf("put service type err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "put service type error",
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

func (controller ServiceController) DeleteServiceType(c *gin.Context) {
	var st model.ServiceType
	if err := c.BindJSON(&st); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "get service type del data error",
			"data": err,
		})
		return
	}
	res, err := controller.stService.DeleteServiceType(st)
	if err != nil {
		controller.log.Errorf("del service type err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "del service type error",
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