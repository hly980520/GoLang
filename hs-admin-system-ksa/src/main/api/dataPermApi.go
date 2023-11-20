package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hs-admin-system-ksa/main/model"
	"hs-admin-system-ksa/main/params"
	"hs-admin-system-ksa/main/result"
	dataPermService "hs-admin-system-ksa/main/service"
	"net/http"
	"strconv"
)

type DataPermApi struct {
}

func NewDataPermApi() Api {
	return &DataPermApi{}
}

func (d DataPermApi) QueryPage(context *gin.Context) {
	var queryParams params.DataPermQuery
	var page model.Page[model.AdminDataPermDto]
	err := context.ShouldBind(&page)
	if err != nil {
		fmt.Println("解析查询参数异常, ", err.Error())
		context.JSON(http.StatusOK, result.Error("10000001", err.Error()))
		return
	}
	err = context.ShouldBind(&queryParams)
	if err != nil {
		fmt.Println("解析查询参数异常, ", err.Error())
		context.JSON(http.StatusOK, result.Error("10000001", err.Error()))
		return
	}
	queryPage, err := dataPermService.QueryPage(page.ToDataPage(), queryParams)
	if err != nil {
		fmt.Println("服务异常, ", err.Error())
		context.JSON(http.StatusOK, result.Error("10000003", err.Error()))
		return
	}
	context.JSON(http.StatusOK, result.Ok(queryPage.ToPageVo()))
}

func (d DataPermApi) Details(context *gin.Context) {
	idStr, flag := context.GetPostForm("id")
	if !flag {
		fmt.Println("解析查询参数异常, 请求参数不存在")
		context.JSON(http.StatusOK, result.Error("10000001", "缺少必要参数"))
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("解析查询参数异常, idStr: ", idStr, err.Error())
		context.JSON(http.StatusOK, result.Error("10000002", err.Error()))
		return
	}
	data, err := dataPermService.QueryById(id)
	if err != nil {
		fmt.Println("服务异常, ", err.Error())
		context.JSON(http.StatusOK, result.Error("10000003", err.Error()))
		return
	}
	context.JSON(http.StatusOK, result.Ok(data))

}

func (d DataPermApi) Update(context *gin.Context) {
	updateParams := new(params.DataPermUpdate)
	err := context.ShouldBind(updateParams)
	if err != nil {
		fmt.Println("解析查询参数异常, ", err.Error())
		context.JSON(http.StatusOK, result.Error("10000002", err.Error()))
		return
	}
	if updateParams == nil || updateParams.Id == 0 {
		fmt.Println("请求参数异常  缺少必要参数 ")
		context.JSON(http.StatusOK, result.Error("10000001", "缺少必要参数"))
		return
	}
	updateParams.UpdatedBy = "peng.haung"
	err = dataPermService.UpdateById(updateParams)
	if err != nil {
		fmt.Println("服务异常, ", err.Error())
		context.JSON(http.StatusOK, result.Error("10000003", err.Error()))
		return
	}
	context.JSON(http.StatusOK, result.Ok(nil))
}

func (d DataPermApi) Delete(context *gin.Context) {
	idStr, flag := context.GetPostForm("id")
	if !flag {
		fmt.Println("解析查询参数异常, 请求参数不存在")
		context.JSON(http.StatusOK, result.Error("10000001", "缺少必要参数"))
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("解析查询参数异常, idStr: ", idStr, err.Error())
		context.JSON(http.StatusOK, result.Error("10000002", err.Error()))
		return
	}
	err = dataPermService.DeleteById(id, "peng.huang")
	if err != nil {
		fmt.Println("服务异常, ", err.Error())
		context.JSON(http.StatusOK, result.Error("10000003", err.Error()))
		return
	}
	context.JSON(http.StatusOK, result.Ok(nil))
}

func (d DataPermApi) Add(context *gin.Context) {
	createParams := new(params.DataPermCreate)
	err := context.ShouldBind(createParams)
	if err != nil {
		fmt.Println("请求参数异常  缺少必要参数 ")
		context.JSON(http.StatusOK, result.Error("10000001", "缺少必要参数"))
		return
	}
	createParams.CreatedBy = "peng.huang"
	err = dataPermService.Create(createParams)
	if err != nil {
		fmt.Println("服务异常, ", err.Error())
		context.JSON(http.StatusOK, result.Error("10000003", err.Error()))
		return
	}
	context.JSON(http.StatusOK, result.Ok(nil))
}
