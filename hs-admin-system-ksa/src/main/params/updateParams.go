package params

// DataPermUpdate 数据权限更新参数
type DataPermUpdate struct {
	Id        int    `form:"id"`
	BizType   int    `form:"bizType"`
	BizId     string `form:"bizId"`
	UserId    int    `form:"userId"`
	Remark    string `form:"remark"`
	UpdatedBy string `form:"updatedBy"`
}

// UserDataPermUpdate 用户数据权限修改参数
type UserDataPermUpdate struct {
	BizType    int    `form:"bizType"`
	BizId      string `form:"bizId"`
	UserIdList []int  `form:"userIdList"`
	EditedBy   string `form:"editedBy"`
}
