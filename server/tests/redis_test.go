package fileManage

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"testing"
)

func TestRedis(t *testing.T) {
	core.Init()
	global.GVA_REDIS.Set(context.Background(), "test", "test", 0)
}
