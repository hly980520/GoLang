package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hs-admin-system-ksa/main/api"
)

// InitRouters 初始化路由信息
func InitRouters() (e error) {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	systemRouter := router.Group("/system/rest")
	//初始化数据权限api
	initDataPermApi(systemRouter)
	//启动服务
	err := router.Run(":8081")
	if err != nil {
		fmt.Println("Routers Init Failed, ", err.Error())
		return err
	}
	fmt.Println("Routers Init Success ")
	return nil
}

// InitDataPermApi 初始化数据权限api
func initDataPermApi(rg *gin.RouterGroup) {
	dataPermApi := api.NewDataPermApi()
	dataPermRouter := rg.Group("/data-perm")
	dataPermRouter.POST("/page", dataPermApi.QueryPage)
	dataPermRouter.POST("/details", dataPermApi.Details)
	dataPermRouter.POST("/delete", dataPermApi.Delete)
	dataPermRouter.POST("/update", dataPermApi.Update)
	dataPermRouter.POST("/add", dataPermApi.Add)
	dataPermRouter.POST("/update-users", dataPermApi.UpdateUsers)
	dataPermRouter.POST("/check", dataPermApi.Check)
	dataPermRouter.POST("/bizTypeList", dataPermApi.QueryBizTypeList)
}
