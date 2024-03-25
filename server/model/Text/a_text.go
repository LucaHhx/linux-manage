// 自动生成模板AText
package Text

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// aText表 结构体  AText
type AText struct {
	global.GVA_MODEL
	Bool       bool    `json:"bool" form:"bool" gorm:"column:bool;type:tinyint;comment:;"`                             //bool字段
	Type       int     `json:"type" form:"type" form:"type" gorm:"column:type;comment:;size:;"`                        //type字段
	Int32      int32   `json:"int32" form:"int32" form:"int32" gorm:"column:int32;comment:;size:32;"`                  //int32字段
	Int64      int64   `json:"int64" form:"int64" form:"int64" gorm:"column:int64;comment:;size:64;"`                  //int64字段
	Varchar    string  `json:"varchar" form:"varchar" gorm:"column:varchar;comment:;size:255;"`                        //varchar字段
	Text       string  `json:"text" form:"text" gorm:"column:text;comment:;type:text;"`                                //text字段
	TinyText   string  `json:"tinyText" form:"tinyText" gorm:"column:tiny_text;comment:;type:tinytext;"`               //tinyText字段
	MediumText string  `json:"mediumText" form:"mediumText" gorm:"column:medium_text;comment:;type:mediumtext;"`       //mediumText字段
	LongText   string  `json:"longText" form:"longText" gorm:"column:long_text;comment:;type:longtext;"`               //longText字段
	Blob       []byte  `json:"blob" form:"blob" gorm:"column:blob;type:blob;size:;comment:;"`                          //blob字段
	Tinyblob   []byte  `json:"tinyblob" form:"tinyblob" gorm:"column:tinyblob;type:tinyblob;size:;comment:;"`          //tinyblob字段
	MediumBlob []byte  `json:"mediumBlob" form:"mediumBlob" gorm:"column:medium_blob;type:mediumblob;size:;comment:;"` //mediumBlob字段
	LongBlob   []byte  `json:"longBlob" form:"longBlob" gorm:"column:long_blob;type:longblob;size:;comment:;"`         //longBlob字段
	Float      float64 `json:"float" form:"float" gorm:"column:float;comment:;"`                                       //float字段
}

// TableName aText表 AText自定义表名 a_text
func (AText) TableName() string {
	return "a_text"
}
