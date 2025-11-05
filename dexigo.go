package dexigo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/JustBugLord/reqtango"
	"github.com/gorilla/websocket"
)

type Okx struct {
	dialer           websocket.Dialer
	mu               sync.Mutex
	publicConnection *websocket.Conn
	rb               *reqtango.RequestBuilder
	subscriptions    []Argument
	handlerCtx       context.Context
	handlerCancel    context.CancelFunc
}

func NewOkxDefault() (*Okx, error) {
	return &Okx{
		dialer: websocket.Dialer{
			EnableCompression: true,
			Proxy:             http.ProxyFromEnvironment,
			HandshakeTimeout:  45 * time.Second,
		},
		rb: reqtango.NewRequestBuilderSimple(),
	}, nil
}

func (okx *Okx) Connect() error {
	publicConn, _, err := okx.dialer.Dial("wss://wsdexpri.okx.com/ws/v5/ipublic", nil)
	if err != nil {
		return errors.New("fail open public socket: " + err.Error())
	}
	okx.publicConnection = publicConn
	if err := okx.ping(); err != nil {
		return errors.New("fail send ping: " + err.Error())
	}
	return nil
}

func (okx *Okx) ping() error {
	return func() error {
		if r := recover(); r != nil {
			return errors.New(fmt.Sprintf("private connection panic: %v", r))
		}
		go func() {
			for {
				if err := okx.Write(websocket.TextMessage, []byte("ping")); err != nil {
					panic("fail write ping to connection: " + err.Error())
				}
				time.Sleep(20 * time.Second)
			}
		}()
		return nil
	}()
}

func (okx *Okx) Subscribe(tokens ...Argument) error {
	if tokens == nil || len(tokens) == 0 {
		return errors.New("tokens is empty")
	}
	okx.subscriptions = append(okx.subscriptions, tokens...)
	return okx.WriteRequest(WSRequest{
		Op:   Subscribe,
		Args: tokens,
	})
}

//func (okx *Okx) Unsubscribe(tokens ...Argument) error {
//	if tokens == nil || len(tokens) == 0 {
//		return errors.New("tokens is empty")
//	}
//	okx.removeSubscriptions(tokens...)
//	return okx.WriteRequest(WSRequest{
//		Op:   Unsubscribe,
//		Args: tokens,
//	})
//}

func (okx *Okx) SetHandler(handler func(response *WSResponse)) error {
	if handler != nil && okx.handlerCancel != nil {
		okx.handlerCancel()
	}
	okx.mu.Lock()
	defer okx.mu.Unlock()
	ctx, cancel := context.WithCancel(context.Background())
	okx.handlerCtx = ctx
	okx.handlerCancel = cancel
	return func() error {
		if r := recover(); r != nil {
			return errors.New(fmt.Sprintf("private connection panic: %v", r))
		}
		go func() {
			for {
				select {
				case <-okx.handlerCtx.Done():
					return
				default:
					response, err := okx.ReadResponse()
					if err != nil {
						panic("fail read response in handler: " + err.Error())
					}
					switch response.Event {
					case Update:
						handler(response)
					}
				}
			}
		}()
		return nil
	}()
}

func (okx *Okx) WriteRequest(wsRequest WSRequest) error {
	bytes, err := json.Marshal(wsRequest)
	if err != nil {
		return errors.New("fail to marshal WSRequest: " + err.Error())
	}
	return okx.Write(websocket.TextMessage, bytes)
}

func (okx *Okx) ReadResponse() (*WSResponse, error) {
	_, bytes, err := okx.Read()
	if err != nil {
		return nil, err
	}
	if bytes == nil {
		return nil, errors.New("response is empty")
	}
	if string(bytes) == "pong" {
		return &WSResponse{
			Event: Pong,
		}, nil
	}
	var response WSResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, errors.New("fail to unmarshal WSResponse: " + err.Error())
	}
	if response.Event == "" {
		response.Event = Update
	}
	return &response, nil
}

func (okx *Okx) Write(msgType int, bytes []byte) error {
	okx.mu.Lock()
	defer okx.mu.Unlock()
	if err := okx.publicConnection.WriteMessage(msgType, bytes); err != nil {
		return errors.New("fail write to public socket: " + err.Error())
	}
	return nil
}

func (okx *Okx) Read() (int, []byte, error) {
	msgType, bytes, err := okx.publicConnection.ReadMessage()
	if err != nil {
		return 0, nil, errors.New("fail read from public socket: " + err.Error())
	}
	return msgType, bytes, nil
}

func (okx *Okx) Close() {
	if okx.publicConnection != nil {
		okx.publicConnection.Close()
	}
	if okx.handlerCancel != nil {
		okx.handlerCancel()
	}
}

func (okx *Okx) removeSubscriptions(source ...Argument) {
	result := make([]Argument, 0, len(source))
	for _, item := range okx.subscriptions {
		contains := false
		for _, arg := range source {
			if item.TokenAddress == arg.TokenAddress {
				contains = true
				break
			}
		}
		if !contains {
			result = append(result, item)
		}
	}
	okx.subscriptions = result
}
