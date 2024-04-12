package enum

import "time"

const (
	Torrent_State_Prepare        = 1 // 准备
	Torrent_State_Downloading    = 2 // 下载中
	Torrent_State_Downloaded     = 3 // 下载完成
	Torrent_State_DownloadFailed = 4 // 下载失败
	Torrent_State_Drop           = 5 // 删除
	Torrent_State_Stop           = 6 // 暂停
)
const Torrent_Dir = "./torrent/"

var Torrent_Interval = 3 * time.Second
