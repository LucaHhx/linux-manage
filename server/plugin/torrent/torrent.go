package torrent

import (
	"fmt"
	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/storage"
	"github.com/dustin/go-humanize"
	"github.com/flipped-aurora/gin-vue-admin/server/enum"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resources"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/websocket"
	pbs "github.com/flipped-aurora/gin-vue-admin/server/protobuf/pbs/resources"
	"math"
	"path/filepath"
	"time"
)

type DownloadTask struct {
	//client   *torrent.Client
	aTorrent       *resources.ATorrent
	torrent        *torrent.Torrent
	bytesCompleted int64
	closed         bool
}

func NewDownloadTask(id int) (*DownloadTask, error) {
	client := global.GVA_TORRENT.GetClient()
	var aTorrent *resources.ATorrent
	err := global.GVA_DB.Where("id = ?", id).First(&aTorrent).Error
	if err != nil {
		return nil, err
	}
	global.GVA_DB.Model(aTorrent).Update("state", enum.Torrent_State_Prepare)
	defer func(err error) {
		if err != nil {
			global.GVA_DB.Model(aTorrent).Update("state", enum.Torrent_State_DownloadFailed)
		}
	}(err)
	spec, err := torrent.TorrentSpecFromMagnetUri(aTorrent.Magnet)
	if err != nil {
		return nil, err
	}
	spec.Storage = storage.NewFile(filepath.Join(enum.Torrent_Dir, aTorrent.Name))
	t, _, err := client.AddTorrentSpec(spec)
	if err != nil {
		return nil, err
	}
	return &DownloadTask{
		aTorrent: aTorrent,
		torrent:  t,
		closed:   false,
	}, nil
}

func (t *DownloadTask) StartTorrent() {
	select {
	case <-t.torrent.GotInfo():
		global.GVA_DB.Model(t.aTorrent).Update("state", enum.Torrent_State_Downloading)
		go t.SetProgress()
		t.torrent.DownloadAll()
	}

	select {
	case <-t.torrent.Closed():
		t.closed = true
		fmt.Println("closed")
		if t.torrent.BytesCompleted() == t.torrent.Length() {
			t.aTorrent.State = enum.Torrent_State_Downloaded
		} else {
			t.aTorrent.State = enum.Torrent_State_Stop
		}
	}
}

func (t *DownloadTask) SetProgress() {
	for range time.Tick(enum.Torrent_Interval) {
		nowBytes := t.torrent.BytesCompleted()
		allBytes := t.torrent.Length()
		gap := nowBytes - t.bytesCompleted
		speed := math.Floor(float64(gap * int64(time.Second) / int64(enum.Torrent_Interval)))
		t.aTorrent.Progress = int32(math.Floor(float64(nowBytes) / float64(allBytes) * 100))
		t.bytesCompleted = nowBytes
		downloadInfo := &pbs.DownloadInfo{
			Id:          int32(t.aTorrent.ID),
			Progress:    int32(math.Floor(float64(nowBytes) / float64(allBytes) * 100)),
			FinishBytes: humanize.Bytes(uint64(nowBytes)),
			TotalBytes:  humanize.Bytes(uint64(allBytes)),
			Speed:       humanize.Bytes(uint64(speed)),
			IsFinished:  nowBytes == allBytes || t.closed,
		}
		if nowBytes == allBytes || t.closed {
			global.GVA_SOCKET.Organize.Broadcast("downloadInfo", websocket.SendOk(downloadInfo))
			if !t.closed {
				t.FinishTorrent()
			}
			return
		}
		global.GVA_SOCKET.Organize.Broadcast("downloadInfo", websocket.SendOk(downloadInfo))
	}
}

func (t *DownloadTask) StopTorrent() {
	global.GVA_DB.Model(t.aTorrent).Update("state", enum.Torrent_State_Stop)
	t.torrent.Drop()
}

func (t *DownloadTask) FinishTorrent() {
	t.torrent.Drop()
}

func (t *DownloadTask) GetKey() int {
	return int(t.aTorrent.ID)
}
