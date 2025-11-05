package dexigo

import (
	"context"
	"encoding/json"
	"errors"
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
	ticker           *time.Timer
	subscriptions    []Argument
	handlers         map[Event]func(response *WSResponse)
	ctx              context.Context
	cancel           context.CancelFunc
	errHandler       func(error)
}

func NewOkxDefault() (*Okx, error) {
	return &Okx{
		dialer: websocket.Dialer{
			EnableCompression: true,
			Proxy:             http.ProxyFromEnvironment,
			HandshakeTimeout:  45 * time.Second,
		},
		rb:       reqtango.NewRequestBuilderSimple(),
		handlers: make(map[Event]func(response *WSResponse), 1),
	}, nil
}

func (okx *Okx) Connect() error {
	publicConn, _, err := okx.dialer.Dial("wss://wsdexpri.okx.com/ws/v5/ipublic", nil)
	if err != nil {
		return errors.New("fail open public socket: " + err.Error())
	}
	okx.publicConnection = publicConn
	ctx, cancel := context.WithCancel(context.Background())
	okx.ctx = ctx
	okx.cancel = cancel
	okx.ping()
	okx.channel()
	return nil
}

func (okx *Okx) ping() {
	go func() {
		ticker := time.NewTimer(20 * time.Second)
		okx.ticker = ticker
		for {
			select {
			case <-okx.ctx.Done():
				return
			case <-ticker.C:
				if err := okx.Write(websocket.TextMessage, []byte("ping")); err != nil {
					if okx.errHandler != nil {
						okx.errHandler(errors.New("fail write ping to connection: " + err.Error()))
					}
				}
			}
		}
	}()
}

func (okx *Okx) channel() {
	go func() {
		for {
			select {
			case <-okx.ctx.Done():
				return
			default:
				response, err := okx.ReadResponse()
				if err != nil {
					if okx.errHandler != nil {
						okx.errHandler(errors.New("fail read response in handler: " + err.Error()))
					}
				}
				if response == nil {
					continue
				}
				if value, ok := okx.handlers[response.Event]; ok && value != nil {
					value(response)
				}
			}
		}
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

func (okx *Okx) AddHandler(event Event, handler func(response *WSResponse)) {
	if okx.handlers == nil {
		okx.handlers = make(map[Event]func(response *WSResponse))
	}
	okx.mu.Lock()
	defer okx.mu.Unlock()
	okx.handlers[event] = handler
}

func (okx *Okx) ErrHandler(handler func(err error)) {
	okx.errHandler = handler
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
	if okx.cancel != nil {
		okx.cancel()
	}
	if okx.ticker != nil {
		okx.ticker.Stop()
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
