// package tests
//
// import (
//
//	"github.com/flipped-aurora/gin-vue-admin/server/global"
//	"time"
//
// )
//
// //import (
// //	"github.com/flipped-aurora/gin-vue-admin/server/global"
// //	"time"
// //)
//
// // 自动生成模板AFileInfo
package fileManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

type AFileInfo struct {
	global.GVA_MODEL
	Blob         []byte    `json:"blob" form:"blob" gorm:"column:blob;type:blob;size:;comment:;"`                          //blob字段
	DateModified time.Time `json:"dateModified" form:"dateModified" gorm:"column:date_modified;comment:;"`                 //dateModified字段
	IsDirectory  bool      `json:"isDirectory" form:"isDirectory" gorm:"column:is_directory;type:tinyint;comment:;"`       //isDirectory字段
	Key          string    `json:"key" form:"key" gorm:"column:key;comment:;size:500;"`                                    //key字段
	LongBlob     []byte    `json:"longBlob" form:"longBlob" gorm:"column:long_blob;type:longblob;size:;comment:;"`         //longBlob字段
	LongText     string    `json:"longText" form:"longText" gorm:"column:long_text;comment:;type:longtext;"`               //longText字段
	MediumBlob   []byte    `json:"mediumBlob" form:"mediumBlob" gorm:"column:medium_blob;type:mediumblob;size:;comment:;"` //mediumBlob字段
	MediumText   string    `json:"mediumText" form:"mediumText" gorm:"column:medium_text;comment:;type:mediumtext;"`       //mediumText字段
	Name         string    `json:"name" form:"name" gorm:"column:name;comment:;size:500;"`                                 //name字段
	Path         string    `json:"path" form:"path" gorm:"column:path;comment:;size:500;"`                                 //path字段
	Size         int64     `json:"size" form:"size" form:"size" gorm:"column:size;comment:;comment:;size:64;"`             //size字段
	Text         string    `json:"text" form:"text" gorm:"column:text;comment:;type:text;"`                                //text字段
	TinyBlob     []byte    `json:"tinyBlob" form:"tinyBlob" gorm:"column:tiny_blob;type:tinyblob;size:;comment:;"`         //tinyBlob字段
	TinyText     string    `json:"tinyText" form:"tinyText" gorm:"column:tiny_text;comment:;type:tinytext;"`               //tinyText字段
	Type         uint8     `json:"type" form:"type" gorm:"column:type;type:smallint;size:8;comment:;"`                     //type字段
}
