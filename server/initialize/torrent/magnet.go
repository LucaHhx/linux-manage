package torrent

import (
	cmap "github.com/orcaman/concurrent-map/v2"
)

type MagnetTask struct {
	tMap cmap.ConcurrentMap[uint, *Torrent]
}

func sharding(key uint) uint32 {
	return uint32(key)
}
func NewMagnetTask() *MagnetTask {
	return &MagnetTask{
		tMap: cmap.NewWithCustomShardingFunction[uint, *Torrent](sharding),
	}
}

var MagnetTaskInstance = NewMagnetTask()

func (m *MagnetTask) Set(t *Torrent) {
	m.tMap.Set(t.aTorrent.ID, t)
}
