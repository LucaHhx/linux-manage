package control

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type Grid struct {
}

func (cg Grid) GetControlGridKey() string {
	return "{control_grid}"
}

func (cg Grid) GetControlGrid(key string) (string, error) {
	return global.GVA_REDIS.HGet(context.Background(), cg.GetControlGridKey(), key).Result()
}

func (cg Grid) SetControlGrid(key, value string) error {
	return global.GVA_REDIS.HSet(context.Background(), cg.GetControlGridKey(), key, value).Err()
}

func (cg Grid) DelControlGrid(key string) error {
	return global.GVA_REDIS.HDel(context.Background(), cg.GetControlGridKey(), key).Err()
}
