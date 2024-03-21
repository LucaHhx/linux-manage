package Text

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Text"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/filter"
)

type ATextService struct {
}

// CreateAText 创建aText表记录
// Author [piexlmax](https://github.com/piexlmax)
func (aTextService *ATextService) CreateAText(aText *Text.AText) (err error) {
	err = global.GVA_DB.Create(aText).Error
	return err
}

// DeleteAText 删除aText表记录
// Author [piexlmax](https://github.com/piexlmax)
func (aTextService *ATextService) DeleteAText(ID string) (err error) {
	err = global.GVA_DB.Delete(&Text.AText{}, "id = ?", ID).Error
	return err
}

// DeleteATextByIds 批量删除aText表记录
// Author [piexlmax](https://github.com/piexlmax)
func (aTextService *ATextService) DeleteATextByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]Text.AText{}, "id in ?", IDs).Error
	return err
}

// UpdateAText 更新aText表记录
// Author [piexlmax](https://github.com/piexlmax)
func (aTextService *ATextService) UpdateAText(up request.UpdateById) (err error) {
	db := global.GVA_DB.Model(Text.AText{})
	err = up.DbUpdate(db).Error
	return err
}

// GetAText 根据ID获取aText表记录
// Author [piexlmax](https://github.com/piexlmax)
func (aTextService *ATextService) GetAText(ID string) (aText Text.AText, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&aText).Error
	return
}

// GetATextInfoList 分页获取aText表记录
// Author [piexlmax](https://github.com/piexlmax)
func (aTextService *ATextService) GetATextInfoList(info request.ListSearch) (list interface{}, total, groupCount int64, err error) {
	//jcMap := helper.GormModelJsonToColumn(Text.AText{})
	limit := info.PageSize
	offset := info.PageSize * (info.Page)

	// 创建db
	db := global.GVA_DB.Model(&Text.AText{})
	var aTexts []Text.AText
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
	err = db.Find(&aTexts).Error

	return aTexts, total, groupCount, err
}
