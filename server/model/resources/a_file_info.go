// 自动生成模板AFileInfo
package resources

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// aFileInfo表 结构体  AFileInfo
type AFileInfo struct {
	global.GVA_MODEL
	Key               string    `json:"key" form:"key" gorm:"column:key;comment:全路径;size:500;" binding:"required"`                               //全路径
	Path              string    `json:"path" form:"path" gorm:"column:path;comment:路径;size:500;"`                                                //路径
	Name              string    `json:"name" form:"name" gorm:"column:name;comment:名称;size:500;" binding:"required"`                             //名称
	DateModified      time.Time `json:"dateModified" form:"dateModified" gorm:"column:date_modified;comment:修改日期;"`                              //修改日期
	Size              int64     `json:"size" form:"size" form:"size" gorm:"column:size;comment:尺寸;size:64;"`                                     //尺寸
	IsDirectory       bool      `json:"isDirectory" form:"isDirectory" gorm:"column:is_directory;type:tinyint;comment:目录;"`                      //目录
	HasSubDirectories bool      `json:"hasSubDirectories" form:"hasSubDirectories" gorm:"column:has_sub_directories;type:tinyint;comment:有子目录;"` //有子目录
	Thumbnail         string    `json:"thumbnail" form:"thumbnail" gorm:"column:thumbnail;comment:图标路径;size:500;"`                               //图标路径
	Tag               string    `json:"tag" form:"tag" gorm:"column:tag;comment:标签;size:255;"`                                                   //标签
	IsMain            bool      `json:"isMain" form:"isMain" gorm:"column:is_main;type:tinyint;comment:主目录;"`                                    //主目录
}

// TableName aFileInfo表 AFileInfo自定义表名 a_file_info
func (AFileInfo) TableName() string {
	return "a_file_info"
}
