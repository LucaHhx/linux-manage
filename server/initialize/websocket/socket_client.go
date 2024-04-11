package websocket

import (
	room2 "github.com/flipped-aurora/gin-vue-admin/server/protobuf/pbs/room"
	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/scheduler"
	"github.com/lonng/nano/session"
)

const OrganizeName = "organize"

type SocketClient struct {
	component.Base
	timer    *scheduler.Timer
	Organize *nano.Group
}

func NewSocketClient() *SocketClient {
	return &SocketClient{
		Organize: nano.NewGroup(OrganizeName),
	}
}

func (sc *SocketClient) Join(s *session.Session, msg []byte) error {
	fakeUID := s.ID() //just use s.ID as uid !!!
	err := s.Bind(fakeUID)
	if err != nil {
		return err
	}
	s.Set(OrganizeName, sc.Organize)
	s.Push("onMembers", &room2.AllMembers{Members: sc.Organize.Members()})
	sc.Organize.Add(s)
	return s.Response(&room2.JoinResponse{Result: "success"})
}
