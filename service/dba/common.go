package dba

import (
	"context"
	"time"

	"github.com/yangioc/bk_pack/proto/dtomsg"
	"github.com/yangioc/bk_pack/util"
	"google.golang.org/protobuf/proto"
)

func (self *Handle) CommonCreate(uuid, reqType, requestStr string, payload []byte) error {
	dbaReq, err := proto.Marshal(&dtomsg.Dto_Msg{
		Type:    reqType,
		Request: requestStr,
		Data:    payload,
	})
	if err != nil {
		return err
	}

	msg, err := util.MsgEncode(&dtomsg.Dto_Base{
		UUID:           uuid,
		StartTime:      util.ServerTimeNow().UnixMicro(),
		ExpirationTime: util.ServerTimeNow().Add(5 * time.Second).UTC().UnixMicro(),
		Payload:        dbaReq,
	})
	if err != nil {
		return err
	}

	if err := self.websocket.Send(context.TODO(), msg); err != nil {
		return err
	}

	return nil
}
