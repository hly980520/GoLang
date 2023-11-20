package dataPermManager

import (
	"fmt"
	"github.com/xyctruth/stream"
	"hs-admin-system-ksa/main/config"
	"hs-admin-system-ksa/main/model"
	"hs-admin-system-ksa/main/params"
	"strconv"
	"time"
)

func SelectPage(dataPage model.DataPage[model.AdminDataPermDto], query params.DataPermQuery) (result *model.DataPage[model.AdminDataPermDto], err error) {
	sql := buildSelectSql(query)

	var pageNo int
	if dataPage.PageNo <= 1 {
		pageNo = 1
	} else {
		pageNo = dataPage.PageNo
	}

	var dataList []model.AdminDataPerm

	orderBy := fmt.Sprintf("%s %s", dataPage.OrderBy, dataPage.Order)
	sql += " ORDER BY " + orderBy
	sql += " LIMIT " + strconv.Itoa((pageNo-1)*dataPage.PageSize) + "," + strconv.Itoa(dataPage.PageSize)
	fmt.Println("Sql: ", sql)
	count, err := config.SqlServer.SQL(sql).FindAndCount(&dataList)
	//count, err := config.SqlServer.FindAndCount(&dataList)
	if err != nil {
		fmt.Println("分页条件查询失败 ", err.Error())
		return nil, err
	}
	dataPage.TotalCount = count

	list := stream.NewSliceByMapping[model.AdminDataPerm, model.AdminDataPermDto, model.AdminDataPermDto](dataList).Map(func(perm model.AdminDataPerm) model.AdminDataPermDto {
		return *perm.ToDto()
	}).ToSlice()

	dataPage.DataList = list
	return &dataPage, nil
}

// SelectById 根据id查询
func SelectById(id int) (result *model.AdminDataPerm, err error) {
	data := new(model.AdminDataPerm)
	_, e := config.SqlServer.ID(id).Get(data)
	if e != nil {
		fmt.Println("查询数据库异常 id: ", id, err.Error())
		return nil, e
	}
	return data, nil
}

func UpdateById(updateParams *params.DataPermUpdate) (count int64, err error) {
	dataPerm := new(model.AdminDataPerm)
	dataPerm.BizType = updateParams.BizType
	dataPerm.BizId = updateParams.BizId
	dataPerm.UserId = updateParams.UserId
	dataPerm.Remark = updateParams.Remark
	dataPerm.UpdatedBy = updateParams.UpdatedBy
	dataPerm.UpdatedTime = time.Now()

	return config.SqlServer.ID(updateParams.Id).Update(dataPerm)
}

func DeleteById(id int, userAccount string) (count int64, err error) {
	return config.SqlServer.ID(id).Update(&model.AdminDataPerm{Deleted: id, UpdatedBy: userAccount, UpdatedTime: time.Now()})
}

func Create(createParams *params.DataPermCreate) (count int64, err error) {
	dataPerm := new(model.AdminDataPerm)
	dataPerm.BizType = createParams.BizType
	dataPerm.BizId = createParams.BizId
	dataPerm.UserId = createParams.UserId
	dataPerm.Remark = createParams.Remark
	dataPerm.CreatedBy = createParams.CreatedBy
	dataPerm.UpdatedBy = createParams.CreatedBy
	dataPerm.CreatedTime = time.Now()
	dataPerm.UpdatedTime = time.Now()
	return config.SqlServer.InsertOne(dataPerm)
}

func buildSelectSql(query params.DataPermQuery) string {
	var sql = "SELECT `id`, `biz_type`, `biz_id`, `user_id`, `remark`, `deleted`, `created_by`, `created_time`, `updated_by`, `updated_time` FROM `admin_data_perm` WHERE deleted = 0"
	if query.BizType != 0 {
		sql += " AND biz_type = " + strconv.Itoa(query.BizType)
	}
	if query.BizId != "" {
		sql += " AND biz_id = " + query.BizId
	}
	if query.UserId != 0 {
		sql += " AND user_id = " + strconv.Itoa(query.UserId)
	}
	return sql
}
