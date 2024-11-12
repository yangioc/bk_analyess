package websocket

import (
	"bk_analysis/model/websocket/socketclient"
	"context"
	"fmt"

	"github.com/yangioc/bk_pack/log"
	"nhooyr.io/websocket"
)

type IClient interface {
	socketclient.IHandle

	Launch(addr string) error
}

type Client struct {
	*socketclient.Handler
	callback socketclient.SocketManagerCallBack
}

func NewClient(callback socketclient.SocketManagerCallBack) *Client {
	return &Client{
		callback: callback,
	}
}

func (socket *Client) Launch(addr string) error {
	if len(addr) == 0 {
		return fmt.Errorf("[Websocket][Launch] addr Error addr: %v", addr)
	}

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	conn, _, err := websocket.Dial(ctx, addr, nil)
	if err != nil {
		return fmt.Errorf("[Websocket][Launch] Listen Error addr: %v", addr)
	}

	socket.Handler = socketclient.New(ctx, conn, socket.callback)
	log.Infof("websocket conn to: %s", addr)
	defer log.Infof("websocket diconn: %s", addr)
	return socket.Handler.Listen()
}
