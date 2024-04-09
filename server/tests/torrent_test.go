package fileManage

import (
	torr "github.com/anacrolix/torrent"
	"github.com/flipped-aurora/gin-vue-admin/server/enum"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize/torrent"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resources"
	"testing"
)

func TestTorr(t *testing.T) {

	cli := createClient()
	addTorrent1, err := torrent.AddTorrent(&resources.ATorrent{
		GVA_MODEL: global.GVA_MODEL{
			ID: 1,
		},
		Name:   "test",
		Magnet: "magnet:?xt=urn:btih:702a42c6d4d19e3c7844726b722e4e98618438bf",
	}, cli)
	if err != nil {
		return
	}

	addTorrent2, err := torrent.AddTorrent(&resources.ATorrent{
		GVA_MODEL: global.GVA_MODEL{
			ID: 2,
		},
		Name:   "test2",
		Magnet: "magnet:?xt=urn:btih:702a42c6d4d19e3c7844726b722e4e98618438bf",
	}, cli)
	if err != nil {
		return
	}

	go addTorrent1.StartTorrent()
	go addTorrent2.StartTorrent()

	cli.WaitAll()
}

func createClient() *torr.Client {
	config := torr.NewDefaultClientConfig()
	config.DataDir = enum.Torrent_Dir
	c, err := torr.NewClient(config)
	if err != nil {
		panic(err)
	}
	return c
}
