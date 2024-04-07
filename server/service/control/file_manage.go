package control

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/control"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resources"
	"github.com/flipped-aurora/gin-vue-admin/server/service/cache"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/helper"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

var FileCache = cache.CacheExamples.FileManage

type FileManage struct {
}

func (fm *FileManage) GetFileInfos() map[string]resources.AFileInfo {
	list := make([]resources.AFileInfo, 0)
	err := global.GVA_DB.Model(&resources.AFileInfo{}).Find(&list).Error
	if err != nil {
		return map[string]resources.AFileInfo{}
	}
	return helper.ListToMap(list, func(t resources.AFileInfo) (string, bool) {
		return t.Key, true
	})
}

func (fm *FileManage) GetFileItems(folderPath string, option control.FileManageOptions) ([]*resources.AFileInfo, error) {
	fileMap := fm.GetFileInfos()
	dir, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}
	list := make([]*resources.AFileInfo, 0)
	for _, entry := range dir {
		var (
			info              os.FileInfo
			hasSubDirectories bool
			subDir            []os.DirEntry
		)
		info, err = entry.Info()
		if err != nil {
			continue
		}
		if info.IsDir() {
			subDir, err = os.ReadDir(filepath.Join(folderPath, entry.Name()))
			if err != nil {
				continue
			}
			hasDirs := lo.Filter(subDir, func(entry os.DirEntry, index int) bool {
				return entry.IsDir()
			})
			hasSubDirectories = len(hasDirs) > 0
		}
		var newFile resources.AFileInfo
		if fileInfo, ok := fileMap[filepath.Join(folderPath, entry.Name())]; ok {
			newFile = fileInfo
		} else {
			newFile = resources.AFileInfo{
				Key:               filepath.Join(folderPath, entry.Name()),
				Path:              folderPath,
				Name:              entry.Name(),
				Size:              info.Size(),
				DateModified:      info.ModTime(),
				IsDirectory:       entry.IsDir(),
				HasSubDirectories: hasSubDirectories,
			}
		}
		if option.Validate(newFile) {
			list = append(list, &newFile)
		}

	}
	return list, nil
}

// FindOrCreateFile 查找或者创建文件
func (fm *FileManage) FindOrCreateFile(fileMd5 string, fileName string, chunkTotal int) (file example.ExaFile, err error) {
	var cfile example.ExaFile
	cfile.FileMd5 = fileMd5
	cfile.FileName = fileName
	cfile.ChunkTotal = chunkTotal

	if errors.Is(global.GVA_DB.Where("file_md5 = ? AND is_finish = ?", fileMd5, true).First(&file).Error, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Where("file_md5 = ? AND file_name = ?", fileMd5, fileName).Preload("ExaFileChunk").FirstOrCreate(&file, cfile).Error
		return file, err
	}
	cfile.IsFinish = true
	cfile.FilePath = file.FilePath
	err = global.GVA_DB.Create(&cfile).Error
	return cfile, err
}

// CreateFileChunk 创建文件切片记录
func (fm *FileManage) CreateFileChunk(id uint, fileChunkPath string, fileChunkNumber int) error {
	var chunk example.ExaFileChunk
	chunk.FileChunkPath = fileChunkPath
	chunk.ExaFileID = id
	chunk.FileChunkNumber = fileChunkNumber
	err := global.GVA_DB.Create(&chunk).Error
	return err
}

// DeleteFileChunk 删除文件切片记录
func (fm *FileManage) DeleteFileChunk(fileMd5 string, filePath string) error {
	err := utils.RemoveChunk(fileMd5)
	if err != nil {
		return err
	}
	var chunks []example.ExaFileChunk
	var file example.ExaFile
	err = global.GVA_DB.Where("file_md5 = ? ", fileMd5).First(&file).
		Updates(map[string]interface{}{
			"IsFinish":  true,
			"file_path": filePath,
		}).Error
	if err != nil {
		return err
	}
	err = global.GVA_DB.Where("exa_file_id = ?", file.ID).Delete(&chunks).Unscoped().Error
	return err
}
