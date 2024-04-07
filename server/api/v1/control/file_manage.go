package control

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/control"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resources"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/service/cache"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
)

var FileCache = cache.CacheExamples.FileManage
var FileService = service.ServiceGroupApp.ControlServiceGroup.FileManage

type FileManageApi struct {
}

// Option 设置参数
func (fm *FileManageApi) Option(c *gin.Context) {
	optionReq := struct {
		Key    string                   `json:"key" form:"key"`
		Type   string                   `json:"type" form:"type"`
		Option control.FileManageOption `json:"option" form:"option"`
	}{}
	err := c.ShouldBindJSON(&optionReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var option *control.FileManageOption
	switch optionReq.Type {
	case control.SaveOption:
		// save
		err = FileCache.SetControlFileManage(optionReq.Key, optionReq.Option)
	case control.LoadOption:
		// load
		option, err = FileCache.GetControlFileManage(optionReq.Key)
	case control.RmOption:
		// rm
		err = FileCache.DelControlFileManage(optionReq.Key)
	default:
		response.FailWithMessage("type error", c)
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(option, "操作成功", c)
	}
}

func (fm *FileManageApi) GetFileItems(c *gin.Context) {
	info := struct {
		Path    string                    `json:"path" form:"path"`
		Options control.FileManageOptions `json:"options" form:"options"`
	}{}
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage("参数错误", c)
	}
	var items []*resources.AFileInfo
	items, err = FileService.GetFileItems(info.Path, info.Options)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取 %s 文件列表失败", info.Path), c)
		return
	}
	response.OkWithData(items, c)
}

// UploadFileChunk 分片上传文件
func (fm *FileManageApi) UploadFileChunk(c *gin.Context) {
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
			_ = FileService.DeleteFileChunk(file.FileMd5, file.FilePath)
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
	file, err = FileService.FindOrCreateFile(fileMd5, fileName, chunkTotal)
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
	if err = FileService.CreateFileChunk(file.ID, pathC, chunkNumber); err != nil {
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
		_ = FileService.DeleteFileChunk(file.FileMd5, file.FilePath)
		return
	}
	response.OkWithMessage("切片创建成功", c)
}

// AbortFileUpload 终止文件上传
func (fm *FileManageApi) AbortFileUpload(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	err := utils.RemoveChunk(fileMd5)
	if err != nil {
		response.FailWithDetailed(nil, "删除文件切片记录失败", c)
		return
	}
	response.OkWithMessage("删除文件切片记录成功", c)
}

// DeleteItem 删除文件
func (fm *FileManageApi) DeleteItem(c *gin.Context) {
	path := c.Query("path")
	err := os.RemoveAll(path)
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// RenameItem 重命名文件
func (fm *FileManageApi) RenameItem(c *gin.Context) {
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

// CopyItem 复制文件
func (fm *FileManageApi) CopyItem(c *gin.Context) {
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

// CreateDirectory 创建文件夹
func (fm *FileManageApi) CreateDirectory(c *gin.Context) {
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

func (fm *FileManageApi) DownloadItem(c *gin.Context) {
	var fileInfo resources.AFileInfo
	err := c.BindJSON(&fileInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var file *os.File
	file, err = os.Open(fileInfo.Key)
	if err != nil {
		response.FailWithMessage("文件打开失败", c)
		return
	}
	defer file.Close()
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(fileInfo.Name)))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		response.FailWithMessage("文件下载失败", c)
		return
	}
	//response.OkWithMessage("文件下载成功", c)
}
