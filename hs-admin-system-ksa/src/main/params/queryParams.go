package params

// DataPermQuery 数据权限查询参数
type DataPermQuery struct {
	BizType    int    `form:"bizType"`
	BizId      string `form:"bizId"`
	BizIdList  []int  `form:"bizIdList"`
	UserId     int    `form:"userId"`
	UserIdList []int  `form:"userIdList"`
}
