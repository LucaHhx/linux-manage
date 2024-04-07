package control

import (
	"context"
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/control"
)

type FileManage struct {
}

func (fm FileManage) GetControlFileManageKey() string {
	return "{control_file_manage}"
}

func (fm FileManage) GetControlFileManage(key string) (*control.FileManageOption, error) {
	bytes, err := global.GVA_REDIS.HGet(context.Background(), fm.GetControlFileManageKey(), key).Bytes()
	if err != nil {
		return nil, nil
	}
	var option *control.FileManageOption
	err = json.Unmarshal(bytes, &option)
	return option, err
}

func (fm FileManage) SetControlFileManage(key string, value control.FileManageOption) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return global.GVA_REDIS.HSet(context.Background(), fm.GetControlFileManageKey(), key, bytes).Err()
}

func (fm FileManage) DelControlFileManage(key string) error {
	return global.GVA_REDIS.HDel(context.Background(), fm.GetControlFileManageKey(), key).Err()
}
