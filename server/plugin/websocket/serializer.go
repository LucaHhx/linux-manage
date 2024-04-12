package websocket

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/protobuf/pbs/common"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// ErrWrongValueType is the error used for marshal the value with protobuf encoding.
var ErrWrongValueType = errors.New("protobuf: convert on wrong type value")

// Serializer implements the serialize.Serializer interface
type Serializer struct{}

// NewSerializer returns a new Serializer.
func NewSerializer() *Serializer {
	return &Serializer{}
}

// Marshal returns the protobuf encoding of v.
func (s *Serializer) Marshal(v interface{}) ([]byte, error) {
	pb, ok := v.(proto.Message)
	if !ok {
		return nil, ErrWrongValueType
	}
	return proto.Marshal(pb)
}

// Unmarshal parses the protobuf-encoded data and stores the result
// in the value pointed to by v.
func (s *Serializer) Unmarshal(data []byte, v interface{}) error {
	basicReq := &common.BasicReq{}
	err := proto.Unmarshal(data, basicReq)
	if err != nil {
		return err
	}

	pb, ok := v.(proto.Message)
	if !ok {
		return ErrWrongValueType
	}
	return anypb.UnmarshalTo(basicReq.Data, pb, proto.UnmarshalOptions{})
}
