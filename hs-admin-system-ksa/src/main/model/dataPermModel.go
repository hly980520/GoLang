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
