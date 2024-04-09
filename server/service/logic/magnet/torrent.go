package magnet

import (
	"github.com/anacrolix/torrent"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resources"
)

type Torrent struct {
	cl *torrent.Client
	jl *resources.ATorrent
}
