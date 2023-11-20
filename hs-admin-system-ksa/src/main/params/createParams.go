package params

// DataPermCreate 数据权限新增参数
type DataPermCreate struct {
	BizType   int    `form:"bizType"`
	BizId     string `form:"bizId"`
	UserId    int    `form:"userId"`
	Remark    string `form:"remark"`
	CreatedBy string `form:"createdBy"`
}
