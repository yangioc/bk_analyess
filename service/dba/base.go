package dba

import (
	"bk_analysis/config"
	"bk_analysis/model/websocket"
	"bk_analysis/model/websocket/socketclient"
	"context"
	"fmt"

	"github.com/yangioc/bk_pack/dto"
	"github.com/yangioc/bk_pack/proto/dtomsg"
	"github.com/yangioc/bk_pack/util"
)

type IHandle interface {
	Launch() error
	OnClose(token string)

	ReadStockClosePrice()
}

type Handle struct {
	websocket      *websocket.Client
	requestTracing dto.FastSyncMap // weboskcet 請求回傳通道管理
}

func New(setting config.Env) *Handle {
	return &Handle{}
}

func (self *Handle) Launch() error {
	if self.websocket == nil {
		self.websocket = websocket.NewClient(self)
	} else if self.websocket.Handler != nil {
		self.websocket.Close(1000, "new socket connect")
	}

	return self.websocket.Launch(config.EnvInfo.Service.DBA.Addr)
}

func (self *Handle) ReceiveMessage(ctx context.Context, socketClient *socketclient.Handler, message []byte) {
	msg, err := util.MsgDecode(message)
	if err != nil {
		panic(err)
	}

	res := &dtomsg.Dto_Msg_Res{}
	if err := util.Unmarshal(msg.Payload, res); err != nil {
		panic(err)
	}

	resChan, ok := self.resChanLoadAndDelete(msg.UUID)
	if !ok {
		panic(fmt.Errorf("requestTracing not find %s", msg.UUID))
	}

	resChan <- res
}

func (self *Handle) resChanLoadAndDelete(uuid string) (chan *dtomsg.Dto_Msg_Res, bool) {
	resChan, ok := self.requestTracing.LoadAndDelete(uuid)
	if !ok {
		return nil, false
	}
	return resChan.(chan *dtomsg.Dto_Msg_Res), true
}

func (self *Handle) OnClose(token string) {}
