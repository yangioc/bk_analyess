package socketclient

import (
	"context"
	"errors"
	"fmt"

	"nhooyr.io/websocket"
)

// 發送訊息 *獨立 goroutine 發送
//
// @params context.Context
// @params message			發送訊息
// @return error			結果回傳
func (self *Handler) Send(ctx context.Context, message []byte) error {
	if self == nil {
		return errors.New("socket Handler error token")
	} else if self.conn == nil {
		return errors.New(fmt.Sprintf("socket disconnect token: %v", self.token))
	}
	err := self.conn.Write(ctx, websocket.MessageBinary, message)
	return err
}

// 增加 client 負載權重
//
// @params int	client 權重值
func (self *Handler) AddWeight(weight int64) {
	self.weight.Add(weight)
}

// 取得 client 負載權重
//
// @rturn int	client 權重值
func (self *Handler) GetWeight() int64 {
	return self.weight.Load()
}
