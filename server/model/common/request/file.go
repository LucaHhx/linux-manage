package request

type FileUploadChunk struct {
	FileMd5     string `json:"fileMd5" form:"fileMd5"`
	FileName    string `json:"fileName" form:"fileName"`
	ChunkMd5    string `json:"chunkMd5" form:"chunkMd5"`
	ChunkNumber int    `json:"chunkNumber" form:"chunkNumber"`
	ChunkTotal  int    `json:"chunkTotal" form:"chunkTotal"`
	//File        multipart.FileHeader `json:"file" form:"file"`
	SavePath string `json:"savePath" form:"savePath"`
}
