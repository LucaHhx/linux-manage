package resources

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resources"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/service/cache/control"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
)

type AFileManageApi struct {
}

var aFileManageService = service.ServiceGroupApp.ResourcesServiceGroup.AFileManageService

// GetFileItems 获取目录下文件
func (afm *AFileManageApi) GetFileItems(c *gin.Context) {
	parentDirectory := c.Query("parentDirectory")
	if parentDirectory == "" {
		response.FailWithMessage("未找到指定目录", c)
	}
	items, err := aFileManageService.GetFileItems(parentDirectory)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取 %s 文件列表失败", parentDirectory), c)
		return
	}
	response.OkWithData(items, c)
}

// UploadFileChunk 分片上传文件
func (afm *AFileManageApi) UploadFileChunk(c *gin.Context) {
	var (
		err        error
		f          multipart.File
		pathC      string
		file       example.ExaFile
		FileHeader *multipart.FileHeader
	)

	fileMd5 := c.Request.FormValue("fileMd5")
	fileName := c.Request.FormValue("fileName")
	chunkMd5 := c.Request.FormValue("chunkMd5")
	chunkNumber, _ := strconv.Atoi(c.Request.FormValue("chunkNumber"))
	chunkTotal, _ := strconv.Atoi(c.Request.FormValue("chunkTotal"))

	defer func(err2 error) {
		if err2 != nil {
			global.GVA_LOG.Error(err2.Error())
			aFileManageService.DeleteFileChunk(file.FileMd5, file.FilePath)
		}
	}(err)

	_, FileHeader, err = c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}

	f, err = FileHeader.Open()
	if err != nil {
		global.GVA_LOG.Error("文件读取失败!", zap.Error(err))
		response.FailWithMessage("文件读取失败", c)
		return
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	cen, _ := io.ReadAll(f)
	if !utils.CheckMd5(cen, chunkMd5) {
		err = fmt.Errorf("md5检查失败")
		response.FailWithMessage("检查md5失败", c)
		return
	}
	file, err = aFileManageService.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		err = fmt.Errorf("查找或创建记录失败,%v", zap.Error(err))
		response.FailWithMessage("查找或创建记录失败", c)
		return
	}
	pathC, err = utils.BreakPointContinue(cen, fileName, chunkNumber, chunkTotal, fileMd5)
	if err != nil {
		err = fmt.Errorf("断点续传失败!,%v", zap.Error(err))
		response.FailWithMessage("断点续传失败", c)
		return
	}
	if err = aFileManageService.CreateFileChunk(file.ID, pathC, chunkNumber); err != nil {
		err = fmt.Errorf("创建文件记录失败!,%v", zap.Error(err))
		response.FailWithMessage("创建文件记录失败", c)
		return
	}
	if chunkNumber == chunkTotal-1 {
		var newPath string
		filePath := c.Request.FormValue("savePath")
		newPath, err = utils.MakeFileToPath(fileName, fileMd5, filePath)
		if err != nil {
			err = fmt.Errorf("%s 文件创建失败!,%v", newPath, zap.Error(err))
			response.FailWithDetailed(nil, "上传文件失败", c)
		} else {
			response.OkWithMessage("上传文件成功", c)
		}
		aFileManageService.DeleteFileChunk(file.FileMd5, file.FilePath)
		return
	}
	response.OkWithMessage("切片创建成功", c)
}

// AbortFileUpload 终止文件上传
func (afm *AFileManageApi) AbortFileUpload(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	err := utils.RemoveChunk(fileMd5)
	if err != nil {
		response.FailWithDetailed(nil, "删除文件切片记录失败", c)
		return
	}
	response.OkWithMessage("删除文件切片记录成功", c)
}

// DeleteItem 删除文件
func (afm *AFileManageApi) DeleteItem(c *gin.Context) {
	path := c.Query("path")
	err := os.RemoveAll(path)
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// RenameItem 重命名文件
func (afm *AFileManageApi) RenameItem(c *gin.Context) {
	reReq := struct {
		File struct {
			Name string `json:"name" form:"name"`
			Path string `json:"path" form:"path"`
		}
		NewName string `json:"newName" form:"newName"`
		NewPath string `json:"newPath" form:"newPath"`
	}{}
	err := c.ShouldBindJSON(&reReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = os.Rename(filepath.Join(reReq.File.Path, reReq.File.Name), filepath.Join(reReq.NewPath, reReq.NewName))
	if err != nil {
		response.FailWithMessage("重命名失败", c)
		return
	}
	response.OkWithMessage("重命名成功", c)
}

// CreateDirectory 创建文件夹
func (afm *AFileManageApi) CreateDirectory(c *gin.Context) {
	dirReq := struct {
		Path string `json:"path" form:"path"`
		Name string `json:"name" form:"name"`
	}{}
	err := c.ShouldBindJSON(&dirReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	dirPath := filepath.Join(dirReq.Path, dirReq.Name)
	err = os.Mkdir(dirPath, os.ModePerm)
	if err != nil {
		response.FailWithMessage("创建文件夹失败", c)
		return
	}
	response.OkWithMessage("创建文件夹成功", c)
}

// CopyItem 复制文件
func (afm *AFileManageApi) CopyItem(c *gin.Context) {
	reReq := struct {
		File    string `json:"file" form:"file"`
		NewName string `json:"newName" form:"newName"`
		NewPath string `json:"newPath" form:"newPath"`
	}{}
	err := c.ShouldBindJSON(&reReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.CopyFile(reReq.File, filepath.Join(reReq.NewPath, reReq.NewName))
	if err != nil {
		response.FailWithMessage("复制失败", c)
		return
	}
	response.OkWithMessage("复制成功", c)
}

// AddTag 添加标签
func (afm *AFileManageApi) AddTag(c *gin.Context) {
	var file resources.AFileInfo
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if file.Tag == "Root" {
		err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
			err = tx.Where("tag = ?", "Root").Delete(&resources.AFileInfo{}).Error
			if err != nil {
				return err
			}
			err = tx.Save(&file).Error
			if err != nil {
				return err
			}
			return nil
		})
	} else {
		err = global.GVA_DB.Save(&file).Error
	}
	if err != nil {
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (afm *AFileManageApi) GetMainList(c *gin.Context) {
	list := make([]resources.AFileInfo, 0)
	err := global.GVA_DB.Find(&list, "tag = ?", "Main").Error
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(list, c)
}

func (afm *AFileManageApi) GetRoot(c *gin.Context) {
	fi := resources.AFileInfo{}
	err := global.GVA_DB.First(&fi, "tag = ?", "Root").Error
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	if fi.Key == "" {
		fi.Key = "/"
	}
	response.OkWithData(fi, c)
}

func (afm *AFileManageApi) UpAttribute(c *gin.Context) {
	attribute := struct {
		Key string      `json:"key" form:"key"`
		Val interface{} `json:"val" form:"val"`
	}{}
	err := c.ShouldBindJSON(&attribute)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%s 属性修改失败", attribute.Key), c)
	}
	err = control.SetAttribute(attribute.Key, attribute.Val)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%s 属性修改失败", attribute.Key), c)
	}
	response.OkWithData(fmt.Sprintf("%s 属性修改成功", attribute.Key), c)
}
