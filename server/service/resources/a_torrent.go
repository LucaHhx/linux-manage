package resources

import (
    "github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resources"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/filter"
)

type ATorrentService struct {
}

// CreateATorrent 创建aTorrent表记录
// Author [piexlmax](https://github.com/piexlmax)
func (aTorrentService *ATorrentService) CreateATorrent(aTorrent *resources.ATorrent) (err error) {
	err = global.GVA_DB.Create(aTorrent).Error
	return err
}

// DeleteATorrent 删除aTorrent表记录
// Author [piexlmax](https://github.com/piexlmax)
func (aTorrentService *ATorrentService)DeleteATorrent(ID string) (err error) {
	err = global.GVA_DB.Delete(&resources.ATorrent{},"id = ?",ID).Error
	return err
}

// DeleteATorrentByIds 批量删除aTorrent表记录
// Author [piexlmax](https://github.com/piexlmax)
func (aTorrentService *ATorrentService)DeleteATorrentByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]resources.ATorrent{},"id in ?",IDs).Error
	return err
}

// UpdateATorrent 更新aTorrent表记录
// Author [piexlmax](https://github.com/piexlmax)
func (aTorrentService *ATorrentService)UpdateATorrent(up request.UpdateById) (err error) {
    db := global.GVA_DB.Model(resources.ATorrent{})
	err = up.DbUpdate(db).Error
	return err
}

// GetATorrent 根据ID获取aTorrent表记录
// Author [piexlmax](https://github.com/piexlmax)
func (aTorrentService *ATorrentService)GetATorrent(ID string) (aTorrent resources.ATorrent, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&aTorrent).Error
	return
}

// GetATorrentInfoList 分页获取aTorrent表记录
// Author [piexlmax](https://github.com/piexlmax)
func (aTorrentService *ATorrentService)GetATorrentInfoList(info request.ListSearch) (list interface{}, total, groupCount int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page)

// 创建db
	db := global.GVA_DB.Model(&resources.ATorrent{})
	var aTorrents []resources.ATorrent
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
	err = db.Find(&aTorrents).Error

	return aTorrents, total, groupCount, err
}
