package model

import "time"

type AdminDataPerm struct {
	Id          int       `xorm:"not null pk autoincr comment('pk 记录id') INT"`
	BizType     int       `xorm:"not null comment('业务类型[1:MSG-消息模板 2:PDP-位置信息 3:CPS-商户列表 4:appconfig-投教]') index(idx_bt_ui_deleted) TINYINT"`
	BizId       string    `xorm:"not null comment('业务id') VARCHAR(100)"`
	UserId      int       `xorm:"not null comment('admin用户id') index(idx_bt_ui_deleted) INT"`
	Remark      string    `xorm:"comment('备注') VARCHAR(500)"`
	Deleted     int       `xorm:"not null default 0 comment('是否删除[否:0 是:记录id]') index(idx_bt_ui_deleted) INT"`
	CreatedBy   string    `xorm:"not null comment('创建人') VARCHAR(50)"`
	CreatedTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	UpdatedBy   string    `xorm:"not null comment('更新人') VARCHAR(50)"`
	UpdatedTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
}

type AdminDataPermDto struct {
	Id          int       `json:"id"`
	BizType     int       `json:"bizType"`
	BizId       string    `json:"bizId"`
	UserId      int       `json:"userId"`
	Remark      string    `json:"remark"`
	CreatedBy   string    `json:"createdBy"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedBy   string    `json:"updatedBy"`
	UpdatedTime time.Time `json:"updatedTime"`
}

// DataPermCreate 数据权限新增参数
type DataPermCreate struct {
	BizType   int    `form:"bizType"`
	BizId     string `form:"bizId"`
	UserId    int    `form:"userId"`
	Remark    string `form:"remark"`
	CreatedBy string `form:"createdBy"`
}

// DataPermQuery 数据权限查询参数
type DataPermQuery struct {
	BizType    int    `form:"bizType"`
	BizId      string `form:"bizId"`
	BizIdList  []int  `form:"bizIdList"`
	UserId     int    `form:"userId"`
	UserIdList []int  `form:"userIdList"`
}

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

// UserDataPermCheck 用户数据权限校验参数
type UserDataPermCheck struct {
	BizType int    `form:"bizType"`
	BizId   string `form:"bizId"`
	UserId  int    `form:"userId"`
}

func (entity *AdminDataPerm) ToDto() (dto *AdminDataPermDto) {
	if entity == nil {
		return nil
	}
	return &AdminDataPermDto{
		Id:          entity.Id,
		BizType:     entity.BizType,
		BizId:       entity.BizId,
		UserId:      entity.UserId,
		Remark:      entity.Remark,
		CreatedBy:   entity.CreatedBy,
		CreatedTime: entity.CreatedTime,
		UpdatedBy:   entity.UpdatedBy,
		UpdatedTime: entity.UpdatedTime,
	}
}
