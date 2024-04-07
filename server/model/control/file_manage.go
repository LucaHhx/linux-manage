package control

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/resources"
	"strings"
)

const (
	SaveOption = "save"
	LoadOption = "load"
	RmOption   = "rm"
)

type FileManageOption struct {
	MainFolder resources.AFileInfo `json:"mainFolder" form:"mainFolder"`
	FileManageOptions
}

type FileManageOptions struct {
	Hide bool `json:"hide" form:"hide"` // 是否显示隐藏文件
}

func (f *FileManageOptions) Validate(file resources.AFileInfo) bool {
	if f.Hide {
		return !strings.HasPrefix(file.Name, ".")
	}
	return true
}
