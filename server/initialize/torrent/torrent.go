package torrent

import (
	"fmt"
	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/storage"
	"github.com/dustin/go-humanize"
	"github.com/flipped-aurora/gin-vue-admin/server/enum"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resources"
	"math"
	"path/filepath"
	"time"
)

type Torrent struct {
	//client   *torrent.Client
	aTorrent       *resources.ATorrent
	torrent        *torrent.Torrent
	bytesCompleted int64
}

func AddTorrent(aTorrent *resources.ATorrent, client *torrent.Client) (*Torrent, error) {
	spec, err := torrent.TorrentSpecFromMagnetUri("magnet:?xt=urn:btih:702a42c6d4d19e3c7844726b722e4e98618438bf")
	if err != nil {
		return nil, err
	}
	spec.Storage = storage.NewFile(filepath.Join(enum.Torrent_Dir, aTorrent.Name))
	t, _, err := client.AddTorrentSpec(spec)
	if err != nil {
		return nil, err
	}
	nt := &Torrent{
		aTorrent: aTorrent,
		torrent:  t,
	}
	MagnetTaskInstance.Set(nt)
	return nt, nil
}

func (t *Torrent) StopTorrent() {
	fmt.Println("closed")
	t.torrent.Drop()
}

func (t *Torrent) StartTorrent() {
	select {
	case <-t.torrent.GotInfo():
		fmt.Println("downloaded")
		go t.SetProgress()
		t.torrent.DownloadAll()
	}

	select {
	case <-t.torrent.Closed():
		fmt.Println("closed")
		if t.torrent.BytesCompleted() == t.torrent.Length() {
			t.aTorrent.State = enum.Torrent_State_Downloaded
		}
	}
}

func (t *Torrent) SetProgress() {
	for range time.Tick(enum.Torrent_Interval) {
		nowBytes := t.torrent.BytesCompleted()
		allBytes := t.torrent.Length()
		if nowBytes == allBytes {

			t.StopTorrent()
		}
		gap := nowBytes - t.bytesCompleted
		speed := math.Floor(float64(gap * int64(time.Second) / int64(enum.Torrent_Interval)))
		t.aTorrent.Progress = int32(math.Floor(float64(nowBytes) / float64(allBytes) * 100))
		t.bytesCompleted = nowBytes
		fmt.Println(fmt.Sprintf("%s %s/%s  %s  %d", t.aTorrent.Name, humanize.Bytes(uint64(nowBytes)),
			humanize.Bytes(uint64(allBytes)), humanize.Bytes(uint64(speed)), t.aTorrent.Progress))
	}
}
