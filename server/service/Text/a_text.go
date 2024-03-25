package Text

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Text"
	TextReq "github.com/flipped-aurora/gin-vue-admin/server/model/Text/request"
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
func (aTextService *ATextService) UpdateAText(aText Text.AText) (err error) {
	err = global.GVA_DB.Save(&aText).Error
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
func (aTextService *ATextService) GetATextInfoList(info TextReq.ATextSearch) (list []Text.AText, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page)
	// 创建db
	db := global.GVA_DB.Model(&Text.AText{})

	for _, condition := range info.Conditions {
		condition.DbCondition(db)
	}
	var aTexts []Text.AText
	err = db.Count(&total).Error
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	for _, sort := range info.Sorts {
		sort.DbSort(db)
	}
	err = db.Find(&aTexts).Error
	return aTexts, total, err
}
