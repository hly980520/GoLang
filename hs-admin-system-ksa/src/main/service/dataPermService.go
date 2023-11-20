package dataPermService

import (
	"errors"
	"fmt"
	dataPermManager "hs-admin-system-ksa/main/manager"
	"hs-admin-system-ksa/main/model"
	"hs-admin-system-ksa/main/params"
)

func QueryPage(dataPage model.DataPage[model.AdminDataPermDto], query params.DataPermQuery) (result *model.DataPage[model.AdminDataPermDto], err error) {
	return dataPermManager.SelectPage(dataPage, query)
}

func QueryById(id int) (dataPermDto *model.AdminDataPermDto, err error) {
	dataPerm, err := dataPermManager.SelectById(id)
	if err != nil {
		return nil, err
	}
	return dataPerm.ToDto(), nil
}

func UpdateById(updateParams *params.DataPermUpdate) (err error) {
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

func DeleteById(id int, userAccount string) (err error) {
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

func Create(createParams *params.DataPermCreate) (err error) {
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
