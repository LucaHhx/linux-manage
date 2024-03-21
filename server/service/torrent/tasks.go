package torrent

import (
	"github.com/anacrolix/torrent"
)

type Task struct {
	Torrent *torrent.Client
}

type TorrentTask struct {
	TaskMap map[int]*torrent.Client
}
