// 自动生成模板ATorrent
package resources

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// aTorrent表 结构体  ATorrent
type ATorrent struct {
 global.GVA_MODEL 
      State  uint8 `json:"state" form:"state" gorm:"column:state;type:smallint;size:8;comment:状态;"`  //状态 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名称;size:255;" binding:"required"`  //名称 
      Progress  int32`json:"progress" form:"progress" form:"progress" gorm:"column:progress;comment:进度;size:32;"`  //进度 
      Magnet  string `json:"magnet" form:"magnet" gorm:"column:magnet;comment:链接;type:text;" binding:"required"`  //链接 
}


// TableName aTorrent表 ATorrent自定义表名 a_torrent
func (ATorrent) TableName() string {
  return "a_torrent"
}

