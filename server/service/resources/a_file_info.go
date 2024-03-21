package resources

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resources"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/filter"
)

type AFileInfoService struct {
}

// CreateAFileInfo 创建aFileInfo表记录
// Author [piexlmax](https://github.com/piexlmax)
func (aFileInfoService *AFileInfoService) CreateAFileInfo(aFileInfo *resources.AFileInfo) (err error) {
	err = global.GVA_DB.Create(aFileInfo).Error
	return err
}

// DeleteAFileInfo 删除aFileInfo表记录
// Author [piexlmax](https://github.com/piexlmax)
func (aFileInfoService *AFileInfoService) DeleteAFileInfo(ID string) (err error) {
	err = global.GVA_DB.Delete(&resources.AFileInfo{}, "id = ?", ID).Error
	return err
}

// DeleteAFileInfoByIds 批量删除aFileInfo表记录
// Author [piexlmax](https://github.com/piexlmax)
func (aFileInfoService *AFileInfoService) DeleteAFileInfoByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]resources.AFileInfo{}, "id in ?", IDs).Error
	return err
}

// UpdateAFileInfo 更新aFileInfo表记录
// Author [piexlmax](https://github.com/piexlmax)
func (aFileInfoService *AFileInfoService) UpdateAFileInfo(up request.UpdateById) (err error) {
	db := global.GVA_DB.Model(resources.AFileInfo{})
	err = up.DbUpdate(db).Error
	return err
}

// GetAFileInfo 根据ID获取aFileInfo表记录
// Author [piexlmax](https://github.com/piexlmax)
func (aFileInfoService *AFileInfoService) GetAFileInfo(ID string) (aFileInfo resources.AFileInfo, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&aFileInfo).Error
	return
}

// GetAFileInfoInfoList 分页获取aFileInfo表记录
// Author [piexlmax](https://github.com/piexlmax)
func (aFileInfoService *AFileInfoService) GetAFileInfoInfoList(info request.ListSearch) (list interface{}, total, groupCount int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page)

	// 创建db
	db := global.GVA_DB.Model(&resources.AFileInfo{})
	var aFileInfos []resources.AFileInfo
	filter.NewParsing(info.Conditions, db).Apply()
	err = db.Count(&total).Error
	//分组
	if len(info.Groups) > 0 {
		var groupLists []response.GroupList
		info.Groups[0].DbGroup(db)
		err = db.Find(&groupLists).Error
		groupCount = int64(len(groupLists))
		for i, groupList := range groupLists {
			groupLists[i].Summary = []int{groupList.Count}
		}
		return groupLists, total, groupCount, err
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	for _, sort := range info.Sorts {
		sort.DbSort(db)
	}
	err = db.Find(&aFileInfos).Error

	return aFileInfos, total, groupCount, err
}
