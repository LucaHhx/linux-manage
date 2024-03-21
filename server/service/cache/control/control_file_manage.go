package control

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

const ControlFileManage = "{control_file_manage}"

func GetAttribute(key string) (string, error) {
	return global.GVA_REDIS.HGet(context.Background(), ControlFileManage, key).Result()
}

func SetAttribute(key string, value interface{}) error {
	return global.GVA_REDIS.HSet(context.Background(), ControlFileManage, key, value).Err()
}

func DelAttribute(key string) error {
	return global.GVA_REDIS.HDel(context.Background(), ControlFileManage, key).Err()
}
