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
	slService service.SlService
	mlService service.MlService
	elService service.ElService
}

func NewServiceController(log *logrus.Logger) *ServiceController {
	return &ServiceController{
		log: log,
		stService: service.NewStService(log),
		slService: service.NewSlService(log),
		mlService: service.NewMlService(log),
		elService: service.NewElService(log),
	}
}

// 服务类型
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

// 服务列表
func (controller ServiceController) GetServiceList(c *gin.Context) {
	// 进行查询
	res, err := controller.slService.FindServiceList()
	if err != nil {
		controller.log.Errorf("get service list err: %v", err)
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

func (controller ServiceController) PostServiceList(c *gin.Context) {
	var sl = model.ServiceList{}
	if err := c.BindJSON(&sl); err != nil || sl.Name == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "get service list post data error",
			"data": err,
		})
		return
	}
	res, err := controller.slService.AddServiceList(sl)
	if err != nil {
		controller.log.Errorf("post service list err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "post service list error",
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

func (controller ServiceController) PutServiceList(c *gin.Context) {
	var sl = model.ServiceList{}
	if err := c.BindJSON(&sl); err != nil || sl.Name == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "get service list post data error",
			"data": err,
		})
		return
	}
	res, err := controller.slService.UpdateServiceList(sl)
	if err != nil {
		controller.log.Errorf("put service list err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "put service list error",
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

func (controller ServiceController) DeleteServiceList(c *gin.Context) {
	var sl model.ServiceList
	if err := c.BindJSON(&sl); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "get service list del data error",
			"data": err,
		})
		return
	}
	res, err := controller.slService.DeleteServiceList(sl)
	if err != nil {
		controller.log.Errorf("del service type err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "del service list error",
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

// 微服务列表
func (controller ServiceController) GetMicroList(c *gin.Context) {
	// 进行查询
	res, err := controller.mlService.FindMicroList()
	if err != nil {
		controller.log.Errorf("get micro list err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "get micro list error",
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

func (controller ServiceController) PostMicroList(c *gin.Context) {
	var sl = model.MicroList{}
	if err := c.BindJSON(&sl); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "get micro list post data error",
			"data": err,
		})
		return
	}
	res, err := controller.mlService.AddMicroList(sl)
	if err != nil {
		controller.log.Errorf("post micro list err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "post micro list error",
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

func (controller ServiceController) PutMicroList(c *gin.Context) {
	var sl = model.MicroList{}
	if err := c.BindJSON(&sl); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "get micro list post data error",
			"data": err,
		})
		return
	}
	res, err := controller.mlService.UpdateMicroList(sl)
	if err != nil {
		controller.log.Errorf("put micro list err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "put micro list error",
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

func (controller ServiceController) DeleteMicroList(c *gin.Context) {
	var sl model.MicroList
	if err := c.BindJSON(&sl); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "get micro list del data error",
			"data": err,
		})
		return
	}
	res, err := controller.mlService.DeleteMicroList(sl)
	if err != nil {
		controller.log.Errorf("del micro list err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "del micro list error",
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

// 收藏服务列表
func (controller ServiceController) GetEnshrineList(c *gin.Context) {
	userName := c.Query("userName")
	if userName == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "get userName error",
		})
		return
	}
	// 进行查询
	res, err := controller.elService.FindEnshrineList(string(userName))
	if err != nil {
		controller.log.Errorf("get enshrine list err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "get enshrine list error",
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

func (controller ServiceController) PostEnshrineList(c *gin.Context) {
	var sl = model.EnshrineList{}
	if err := c.BindJSON(&sl); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "get enshrine list post data error",
			"data": err,
		})
		return
	}
	res, err := controller.elService.AddEnshrineList(sl)
	if err != nil {
		controller.log.Errorf("post enshrine list err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "post enshrine list error",
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

func (controller ServiceController) DeleteEnshrineList(c *gin.Context) {
	var sl model.EnshrineList
	if err := c.BindJSON(&sl); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "get enshrine list del data error",
			"data": err,
		})
		return
	}
	res, err := controller.elService.DeleteEnshrineList(sl)
	if err != nil {
		controller.log.Errorf("del enshrine list err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "del enshrine list error",
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