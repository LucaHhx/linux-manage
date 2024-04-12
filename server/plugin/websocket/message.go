package websocket

import (
	"github.com/flipped-aurora/gin-vue-admin/server/protobuf/pbs/common"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func SendOk(data proto.Message) *common.BasicRep {
	rep := &common.BasicRep{
		Head: &common.RepHead{
			Code: 0,
			Msg:  "ok",
		},
		Data: &anypb.Any{},
	}
	a, err := anypb.New(data)
	if err != nil {
		rep.Head.Code = 1
		rep.Head.Msg = err.Error()
		return rep
	}
	rep.Data = a
	return rep
}

func SendErr(code common.Code, msg string) *common.BasicRep {
	return &common.BasicRep{
		Head: &common.RepHead{
			Code: code,
			Msg:  msg,
		},
	}
}
