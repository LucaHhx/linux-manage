package control

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type ControlGrid struct {
}

func (cg ControlGrid) GetControlGridKey() string {
	return "{control_grid}"
}

func (cg ControlGrid) GetControlGrid(key string) (string, error) {
	return global.GVA_REDIS.HGet(context.Background(), cg.GetControlGridKey(), key).Result()
}

func (cg ControlGrid) SetControlGrid(key, value string) error {
	return global.GVA_REDIS.HSet(context.Background(), cg.GetControlGridKey(), key, value).Err()
}

func (cg ControlGrid) DelControlGrid(key string) error {
	return global.GVA_REDIS.HDel(context.Background(), cg.GetControlGridKey(), key).Err()
}
