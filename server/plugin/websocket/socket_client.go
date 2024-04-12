package websocket

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/protobuf/pbs/common"
	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/pipeline"
	"github.com/lonng/nano/scheduler"
	"github.com/lonng/nano/session"
	"time"
)

const OrganizeName = "organize"

type (
	SocketClient struct {
		component.Base
		timer    *scheduler.Timer
		Organize *nano.Group
		Pip      pipeline.Pipeline
		stats    *Stats
	}

	Stats struct {
		component.Base
		timer         *scheduler.Timer
		outboundBytes int
		inboundBytes  int
	}
)

func NewSocketClient() *SocketClient {
	stats := &Stats{}
	client := &SocketClient{
		Organize: nano.NewGroup(OrganizeName),
		Pip:      pipeline.New(),
		stats:    stats,
	}
	client.Pip.Outbound().PushBack(stats.Outbound)
	client.Pip.Inbound().PushBack(stats.Inbound)
	return client
}

func (stats *Stats) Outbound(s *session.Session, msg *pipeline.Message) error {
	stats.outboundBytes += len(msg.Data)
	return nil
}

func (stats *Stats) Inbound(s *session.Session, msg *pipeline.Message) error {
	stats.inboundBytes += len(msg.Data)
	return nil
}

func (stats *Stats) AfterInit() {
	stats.timer = scheduler.NewTimer(time.Minute, func() {
		println("OutboundBytes", stats.outboundBytes)
		println("InboundBytes", stats.outboundBytes)
	})
}

// AfterInit component lifetime callback
func (sc *SocketClient) AfterInit() {
	session.Lifetime.OnClosed(func(s *session.Session) {
		sc.Organize.Leave(s)
	})
	sc.timer = scheduler.NewTimer(time.Minute, func() {
		println(fmt.Sprintf("UserCount: Time=%s, Count=%d",
			time.Now().String(), sc.Organize.Count()))
	})
}

func (sc *SocketClient) Create(s *session.Session, msg []byte) error {
	fakeUID := s.ID() //just use s.ID as uid !!!
	err := s.Bind(fakeUID)
	if err != nil {
		return err
	}
	s.Set(OrganizeName, sc.Organize)
	err = sc.Organize.Add(s)
	if err != nil {
		return err
	}
	err = sc.Organize.Broadcast("onMessage", SendOk(&common.Message{Message: "连接成功"}))
	if err != nil {
		return err
	}
	return s.Response(SendOk(&common.Message{Message: "连接成功"}))
}

func (sc *SocketClient) Message(s *session.Session, msg *common.Message) error {
	fmt.Println("message", msg.Message)
	s.Push("onMessage", SendOk(&common.Message{Message: fmt.Sprintf("Push收到消息：%s", msg.Message)}))
	return sc.Organize.Broadcast("onMessage", SendOk(&common.Message{Message: fmt.Sprintf("Broadcast收到消息：%s", msg.Message)}))
}

func (sc *SocketClient) Process(s *session.Session, msg *common.Message) error {
	fmt.Println("message", msg.Message)
	s.Push("onMessage", SendOk(&common.Message{Message: fmt.Sprintf("Push收到消息：%s", msg.Message)}))
	return sc.Organize.Broadcast("onMessage", SendOk(&common.Message{Message: fmt.Sprintf("Broadcast收到消息：%s", msg.Message)}))
}
