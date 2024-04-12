package torrent

import (
	"fmt"
	"github.com/anacrolix/torrent"
	"github.com/flipped-aurora/gin-vue-admin/server/enum"
)

type DownloadTask interface {
	StartTorrent()
	SetProgress()
	StopTorrent()
	FinishTorrent()
	GetKey() int
}

type MagnetClient struct {
	client  *torrent.Client
	taskMap map[int]DownloadTask
}

func (m *MagnetClient) GetClient() *torrent.Client {
	return m.client
}

func createClient() *torrent.Client {
	config := torrent.NewDefaultClientConfig()
	config.DataDir = enum.Torrent_Dir
	c, err := torrent.NewClient(config)
	if err != nil {
		panic(err)
	}
	return c
}

func NewMagnetClient() *MagnetClient {
	return &MagnetClient{
		client:  createClient(),
		taskMap: make(map[int]DownloadTask),
	}
}

func (m *MagnetClient) Download(task DownloadTask) error {
	m.taskMap[task.GetKey()] = task
	go task.StartTorrent()
	return nil
}

func (m *MagnetClient) Stop(id int) error {
	task, ok := m.taskMap[id]
	if !ok {
		return fmt.Errorf("task not found")
	}
	task.StopTorrent()
	return nil
}
