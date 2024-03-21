package {{.Package}}

import (
    "github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/{{.Package}}"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/filter"
)

type {{.StructName}}Service struct {
}

{{- $db := "" }}
{{- if eq .BusinessDB "" }}
 {{- $db = "global.GVA_DB" }}
{{- else}}
 {{- $db =  printf "global.MustGetGlobalDBByDBName(\"%s\")" .BusinessDB   }}
{{- end}}

// Create{{.StructName}} 创建{{.Description}}记录
// Author [piexlmax](https://github.com/piexlmax)
func ({{.Abbreviation}}Service *{{.StructName}}Service) Create{{.StructName}}({{.Abbreviation}} *{{.Package}}.{{.StructName}}) (err error) {
	err = {{$db}}.Create({{.Abbreviation}}).Error
	return err
}

// Delete{{.StructName}} 删除{{.Description}}记录
// Author [piexlmax](https://github.com/piexlmax)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Delete{{.StructName}}({{.PrimaryField.FieldJson}} string{{- if .AutoCreateResource -}},userID uint{{- end -}}) (err error) {
	{{- if .AutoCreateResource }}
	err = {{$db}}.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&{{.Package}}.{{.StructName}}{}).Where("{{.PrimaryField.ColumnName}} = ?", {{.PrimaryField.FieldJson}}).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&{{.Package}}.{{.StructName}}{},"{{.PrimaryField.ColumnName}} = ?",{{.PrimaryField.FieldJson}}).Error; err != nil {
              return err
        }
        return nil
	})
    {{- else }}
	err = {{$db}}.Delete(&{{.Package}}.{{.StructName}}{},"{{.PrimaryField.ColumnName}} = ?",{{.PrimaryField.FieldJson}}).Error
	{{- end }}
	return err
}

// Delete{{.StructName}}ByIds 批量删除{{.Description}}记录
// Author [piexlmax](https://github.com/piexlmax)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Delete{{.StructName}}ByIds({{.PrimaryField.FieldJson}}s []string {{- if .AutoCreateResource }},deleted_by uint{{- end}}) (err error) {
	{{- if .AutoCreateResource }}
	err = {{$db}}.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&{{.Package}}.{{.StructName}}{}).Where("{{.PrimaryField.ColumnName}} in ?", {{.PrimaryField.FieldJson}}s).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("{{.PrimaryField.ColumnName}} in ?", {{.PrimaryField.FieldJson}}s).Delete(&{{.Package}}.{{.StructName}}{}).Error; err != nil {
            return err
        }
        return nil
    })
    {{- else}}
	err = {{$db}}.Delete(&[]{{.Package}}.{{.StructName}}{},"{{.PrimaryField.ColumnName}} in ?",{{.PrimaryField.FieldJson}}s).Error
    {{- end}}
	return err
}

// Update{{.StructName}} 更新{{.Description}}记录
// Author [piexlmax](https://github.com/piexlmax)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Update{{.StructName}}(up request.UpdateById) (err error) {
    db := global.GVA_DB.Model({{.Package}}.{{.StructName}}{})
	err = up.DbUpdate(db).Error
	return err
}

// Get{{.StructName}} 根据{{.PrimaryField.FieldJson}}获取{{.Description}}记录
// Author [piexlmax](https://github.com/piexlmax)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Get{{.StructName}}({{.PrimaryField.FieldJson}} string) ({{.Abbreviation}} {{.Package}}.{{.StructName}}, err error) {
	err = {{$db}}.Where("{{.PrimaryField.ColumnName}} = ?", {{.PrimaryField.FieldJson}}).First(&{{.Abbreviation}}).Error
	return
}

// Get{{.StructName}}InfoList 分页获取{{.Description}}记录
// Author [piexlmax](https://github.com/piexlmax)
func ({{.Abbreviation}}Service *{{.StructName}}Service)Get{{.StructName}}InfoList(info request.ListSearch) (list interface{}, total, groupCount int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page)

// 创建db
	db := global.GVA_DB.Model(&{{.Package}}.{{.StructName}}{})
	var {{.Abbreviation}}s []{{.Package}}.{{.StructName}}
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
	err = db.Find(&{{.Abbreviation}}s).Error

	return {{.Abbreviation}}s, total, groupCount, err
}
