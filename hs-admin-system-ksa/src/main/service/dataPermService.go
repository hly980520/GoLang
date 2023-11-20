package service

import (
	"errors"
	"fmt"
	"hs-admin-system-ksa/main/manager"
	"hs-admin-system-ksa/main/model"
)

type DataPermService struct {
}

func NewDataPermService() DataPermService {
	return DataPermService{}
}

var dataPermManager = manager.NewDataPermManager()

func (d DataPermService) QueryPage(dataPage model.DataPage[model.AdminDataPermDto], query model.DataPermQuery) (result *model.DataPage[model.AdminDataPermDto], err error) {
	return dataPermManager.SelectPage(dataPage, query)
}

func (d DataPermService) QueryById(id int) (dataPermDto *model.AdminDataPermDto, err error) {
	dataPerm, err := dataPermManager.SelectById(id)
	if err != nil {
		return nil, err
	}
	return dataPerm.ToDto(), nil
}

func (d DataPermService) UpdateById(updateParams *model.DataPermUpdate) (err error) {
	updateById, err := dataPermManager.UpdateById(updateParams)
	if err != nil {
		return err
	}
	if updateById < 1 {
		fmt.Println("更新失败 id: ", updateParams.Id)
		return errors.New("更新失败")
	}
	return nil
}

func (d DataPermService) DeleteById(id int, userAccount string) (err error) {
	deleteById, err := dataPermManager.DeleteById(id, userAccount)
	if err != nil {
		return err
	}
	if deleteById < 1 {
		fmt.Println("删除失败 id: ", id)
		return errors.New("删除失败")
	}
	return nil
}

func (d DataPermService) Create(createParams *model.DataPermCreate) (err error) {
	create, err := dataPermManager.Create(createParams)
	if err != nil {
		return err
	}
	if create < 1 {
		fmt.Println("新增失败")
		return errors.New("新增失败")
	}
	return nil
}

func (d DataPermService) Check(checkParams *model.UserDataPermCheck) (result bool, err error) {
	list, err := dataPermManager.SelectList(model.DataPermQuery{BizType: checkParams.BizType, BizId: checkParams.BizId, UserId: checkParams.UserId})
	if err != nil {
		fmt.Println("校验用户数据权限失败")
		return false, err
	}
	if list == nil || len(list) == 0 {
		fmt.Println(fmt.Sprintf("校验用户数据权限成功  用户没有权限 [biType:%s|bizId:%s|userId:%s]",
			checkParams.BizType, checkParams.BizId, checkParams.UserId))
		return false, nil
	}
	fmt.Println("校验用户数据权限成功  用户拥有权限")
	return true, nil
}
