package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hs-admin-system-ksa/main/enum"
	"hs-admin-system-ksa/main/model"
	"hs-admin-system-ksa/main/result"
	"hs-admin-system-ksa/main/service"
	"net/http"
	"strconv"
)

type DataPermApi struct {
}

func NewDataPermApi() DataPermApi {
	return DataPermApi{}
}

var dataPermService = service.NewDataPermService()

func (d DataPermApi) QueryPage(context *gin.Context) {
	var queryParams model.DataPermQuery
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
	updateParams := new(model.DataPermUpdate)
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
	createParams := new(model.DataPermCreate)
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

func (d DataPermApi) Check(context *gin.Context) {
	userDataPermCheck := new(model.UserDataPermCheck)
	err := context.ShouldBind(userDataPermCheck)
	if err != nil || userDataPermCheck.UserId == 0 || userDataPermCheck.BizType == 0 || userDataPermCheck.BizId == "" {
		fmt.Println("请求参数异常  缺少必要参数 ")
		context.JSON(http.StatusOK, result.Error("10000001", "缺少必要参数"))
		return
	}
	check, err := dataPermService.Check(userDataPermCheck)
	if err != nil {
		fmt.Println("服务异常, ", err.Error())
		context.JSON(http.StatusOK, result.Error("10000003", err.Error()))
		return
	}
	context.JSON(http.StatusOK, result.Ok(check))
}

func (d DataPermApi) QueryBizTypeList(context *gin.Context) {
	context.JSON(http.StatusOK, result.Ok(enum.GetAllBizType()))
}

func (d DataPermApi) UpdateUsers(context *gin.Context) {
	userDataPermUpdate := new(model.UserDataPermUpdate)
	err := context.ShouldBind(userDataPermUpdate)
	if err != nil || userDataPermUpdate.BizType == 0 || userDataPermUpdate.BizId == "" || len(userDataPermUpdate.UserIdList) == 0 {
		fmt.Println("请求参数异常  缺少必要参数 ")
		context.JSON(http.StatusOK, result.Error("10000001", "缺少必要参数"))
		return
	}
	userDataPermUpdate.EditedBy = "peng.huang"
	context.JSON(http.StatusOK, result.Ok(nil))
}
