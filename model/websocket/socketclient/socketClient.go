package socketclient

import (
	"context"
	"time"

	"github.com/yangioc/bk_pack/log"
	"go.uber.org/atomic"
	"nhooyr.io/websocket"
)

const (
	// 讀取限制 超過此大小連線將會被異常中斷
	readLimit int64 = 1 * 1024 * 1024 // 1M
)

type SocketManagerCallBack interface {
	// 連線通知分為兩部份驗證前, 驗證後
	// 驗證前: 連線通知只通知道 socket Manager 但目前沒有需要處理的事
	// 驗證後: 驗證流程必須 收到 register 或 reconnect 才會完成不再由這邊做通知
	// OnNewConnect(token string, socketClient *Handler)
	OnClose(token string)
	ReceiveMessage(ctx context.Context, socketClient *Handler, message []byte)
}

type IHandle interface {
	Send(ctx context.Context, message []byte) error
	GetToken() string
	AddWeight(weight int64)
}

type Handler struct {
	httpCtx    context.Context // http 升級訊號的 ctx
	cancelFunc func()          // 關閉這次 http 連線的方法

	conn *websocket.Conn

	// receiveCallBack       ReceiveCallBack
	socketManagerCallBack SocketManagerCallBack

	weight atomic.Int64 // 房間負載
	token  string       // Socket Token
	// shutdownChan chan struct{}
}

// Websocket Client 建立新物件
//
// @params context.Context	此連線的 Context
// @params *websocket.Conn	Websocket操作界面
// @return Handler			client物件
func New(ctx context.Context, conn *websocket.Conn, callBack SocketManagerCallBack) *Handler {
	client := &Handler{
		conn:                  conn,
		socketManagerCallBack: callBack,
	}

	conn.SetReadLimit(readLimit)
	client.httpCtx, client.cancelFunc = context.WithCancel(ctx)
	go client.read(client.httpCtx)

	return client
}

// Websocket Client 啟動監聽
//
// @params context.Context client 啟動監聽
func (self *Handler) Listen() error {
	self.listenHandle()
	log.Infof("[SocketClient][%v] close done.", self.token)
	return nil
}

// Websocket Client 識別編號
//
// @return string 識別編號
func (self *Handler) GetToken() string {
	return self.token
}

// Websocket Client 識別編號
//
// @params string 識別編號
func (self *Handler) SetToken(token string) {
	self.token = token
}

// 監聽所有關閉 websocket 的訊號
func (self *Handler) listenHandle() {
	<-self.httpCtx.Done()

	// 關閉後續處理
	self.socketManagerCallBack.OnClose(self.GetToken())
}

// WebSocket 監聽
func (self *Handler) read(ctx context.Context) {
	defer self.close(websocket.StatusInternalError, "read err")
	for {
		select {
		case <-ctx.Done():
			return

		default:
			handleCtx := context.TODO()
			_, msg, err := self.conn.Read(handleCtx)

			// server 斷線，先關閉 read()，由pingLoop()發動重新連線
			if err != nil {
				if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
					log.Errorf("[SocketClient][%v] connect closed: %v", self.token, err)
				} else {
					log.Errorf("[SocketClient][%v] read error: %v", self.token, err)
				}
				return
			}

			// 訊息轉拋
			go self.socketManagerCallBack.ReceiveMessage(handleCtx, self, msg)
		}
	}
}

func (self *Handler) Ping() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for range ticker.C {
		if err := self.conn.Ping(self.httpCtx); err != nil {
			log.Errorf("[SocketClient][%v] ping error: %v", self.token, err)
			return
		}
	}
}

func (self *Handler) Close(code websocket.StatusCode, errorMsg string) {
	self.close(code, errorMsg)
}

// 關閉 weboscket connect 並關閉 socket client ctx
func (self *Handler) close(code websocket.StatusCode, errorMsg string) {
	self.conn.Close(code, errorMsg)
	self.cancelFunc()
}
