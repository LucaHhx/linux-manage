// 自动生成模板{{.StructName}}
package {{.Package}}

import (
	{{ if .GvaModel }}"github.com/flipped-aurora/gin-vue-admin/server/global"{{ end }}
	{{ if or .HasTimer }}"time"{{ end }}
	{{ if .NeedJSON }}"gorm.io/datatypes"{{ end }}
)

// {{.Description}} 结构体  {{.StructName}}
type {{.StructName}} struct {
{{ if .GvaModel }} global.GVA_MODEL {{ end }}
      {{- range .Fields}}
            {{- if eq .FieldType "enum" }}
      {{.FieldName}}  uint8 `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"{{- if .PrimaryKey -}}primarykey;{{- end -}}column:{{.ColumnName}};type:smallint;size:{{.DataTypeLong}};comment:{{.Comment}};" {{- if .Require }} binding:"required"{{- end -}}`
            {{- else if eq .FieldType "bool" }}
      {{.FieldName}}  bool `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"{{- if .PrimaryKey -}}primarykey;{{- end -}}column:{{.ColumnName}};type:tinyint;comment:{{.Comment}};" {{- if .Require }} binding:"required"{{- end -}}`
            {{- else if eq .FieldType "int" }}
      {{.FieldName}}  int{{- if .DataTypeLong -}}{{.DataTypeLong}}{{- end -}}  `json:"{{.FieldJson}}" form:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"{{- if .PrimaryKey -}}primarykey;{{- end -}}column:{{.ColumnName}};comment:{{.Comment}};size:{{.DataTypeLong}};"`
            {{- else if eq .FieldType "blob" }}
      {{.FieldName}}  []byte `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"{{- if .PrimaryKey -}}primarykey;{{- end -}}column:{{.ColumnName}};type:{{.DataType}};size:{{.DataTypeLong}};comment:{{.Comment}};" {{- if .Require }} binding:"required"{{- end -}}`
            {{- else if eq .FieldType "picture" }}
      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"{{- if .PrimaryKey -}}primarykey;{{- end -}}column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}" {{- if .Require }} binding:"required"{{- end -}}`
            {{- else if eq .FieldType "video" }}
      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"{{- if .PrimaryKey -}}primarykey;{{- end -}}column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}" {{- if .Require }} binding:"required"{{- end -}}`
             {{- else if eq .FieldType "file" }}
      {{.FieldName}}  datatypes.JSON `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"{{- if .PrimaryKey -}}primarykey;{{- end -}}column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}" {{- if .Require }} binding:"required"{{- end -}}`
            {{- else if eq .FieldType "pictures" }}
      {{.FieldName}}  datatypes.JSON `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"{{- if .PrimaryKey -}}primarykey;{{- end -}}column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}" {{- if .Require }} binding:"required"{{- end -}}`
            {{- else if eq .FieldType "richtext" }}
      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"{{- if .PrimaryKey -}}primarykey;{{- end -}}column:{{.ColumnName}};comment:{{.Comment}};type:{{.DataType}};" {{- if .Require }} binding:"required"{{- end -}}`
            {{- else if eq .FieldType "json" }}
      {{.FieldName}}  datatypes.JSON `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"{{- if .PrimaryKey -}}primarykey;{{- end -}}column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}type:text;" {{- if .Require }} binding:"required"{{- end -}}`
            {{- else if ne .FieldType "string" }}
      {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"{{- if .PrimaryKey -}}primarykey;{{- end -}}column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}" {{- if .Require }} binding:"required"{{- end -}}`
            {{- else }}
      {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"{{- if .PrimaryKey -}}primarykey;{{- end -}}column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}" {{- if .Require }} binding:"required"{{- end -}}`
            {{- end }}  {{ if .FieldDesc }}//{{.FieldDesc}} {{ end }} {{- end }}
      {{- if .AutoCreateResource }}
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
      {{- end}}
}

{{ if .TableName }}
// TableName {{.Description}} {{.StructName}}自定义表名 {{.TableName}}
func ({{.StructName}}) TableName() string {
  return "{{.TableName}}"
}
{{ end }}
