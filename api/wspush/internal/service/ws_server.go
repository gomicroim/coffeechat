package service

import (
	"context"
	"encoding/json"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gomicroim/gomicroim/pkg/log"
	"github.com/gomicroim/gomicroim/pkg/rescue"
	pb "github.com/gomicroim/gomicroim/protos/wspush"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
	"user/api/user"
	"wspush/api/wspush"
	"wspush/internal/biz"
	"wspush/internal/conf"
)

var (
	upGrader = websocket.Upgrader{
		ReadBufferSize: 4096,
		// websocket写缓冲区，对象池重复使用，不限制数据发送大小，可以比该值大
		WriteBufferSize: 4096,
		// 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	} // use default options
)

type WsServer struct {
	auth user.AuthClient
	*khttp.Server
}

func NewWsServer(server *conf.Server, client user.AuthClient) *WsServer {
	w := &WsServer{
		auth: client,
	}

	router := mux.NewRouter()
	router.HandleFunc("/ws", w.onWsHandler)

	w.Server = khttp.NewServer(khttp.Address(server.Ws.Addr))
	w.Server.HandlePrefix("/", router)

	return w
}

func (w *WsServer) onWsHandler(res http.ResponseWriter, req *http.Request) {
	// 提取token
	jwtToken := ""
	temp := req.Header.Get("authorization")
	prefix := "Bearer "
	if strings.Contains(temp, prefix) {
		jwtToken = temp[len(prefix):]
	}

	if jwtToken == "" {
		responseError(res, int(wspush.ErrorReason_TOKEN_MISS), wspush.ErrorReason_TOKEN_MISS.String())
		log.L.Warn(wspush.ErrorReason_TOKEN_MISS.String())
		return
	}

	// valid token
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	validRes, err := w.auth.TokenValid(ctx, &user.TokenValidRequest{
		Token:          jwtToken,
		IsRefreshToken: false,
	})
	if err != nil {
		responseError(res, int(wspush.ErrorReason_TOKEN_INVALID), err.Error())
		log.L.Warn(err.Error(), zap.Error(err), zap.String("token", jwtToken))
		return
	}

	// save client type
	clientInfo := biz.ClientInfo{
		UserId:     validRes.ClientInfo.UserId,
		DeviceId:   validRes.ClientInfo.DeviceId,
		ClientType: parseClientType(validRes.ClientInfo.ClientType),
		Domain:     validRes.ClientInfo.Domain,
		AppVersion: validRes.ClientInfo.AppVersion,
	}

	wsConn, err := upGrader.Upgrade(res, req, nil)
	if err != nil {
		log.L.Error("Upgrade failed", zap.Any("ClientInfo", validRes.ClientInfo), zap.Error(err))
		return
	}

	requestId := req.URL.Query().Get("request_id")
	logger := log.With(log.L,
		zap.Int64("userId", clientInfo.UserId),
		zap.String("deviceId", clientInfo.DeviceId),
		zap.Int32("appVersion", clientInfo.AppVersion),
		zap.String("requestId", requestId),
		zap.Any("clientType", clientInfo.ClientType),
		zap.Int32("socketFd", biz.FD(wsConn)),
	)

	go rescue.WithRecover(func() {
		deviceConn := biz.NewConnection(wsConn, clientInfo, logger)
		curUser, remain := biz.Users.Add(deviceConn)

		// Handle user offline
		defer func() {
			remain = biz.Users.Remove(deviceConn)
			curUser.SetOffline(ctx)
			logger.Info("user offline", zap.Int("online devices", remain))
		}()

		// Handle user online
		setOnline := func() bool {
			newCtx, cancel := context.WithTimeout(ctx, time.Second*3)
			defer cancel()
			return curUser.SetOnline(newCtx)
		}
		if !setOnline() {
			wsConn.Close()
			return
		}

		logger.Info("user online", zap.Int("online devices", remain))
		deviceConn.OnLoop()
	})
}

func parseClientType(clientType string) pb.IMClientType {
	if strings.ToLower(clientType) == "android" {
		return pb.IMClientType_ClientTypeAndroid
	} else if strings.ToLower(clientType) == "ios" {
		return pb.IMClientType_ClientTypeIos
	} else if strings.ToLower(clientType) == "web" {
		return pb.IMClientType_ClientTypeWeb
	}
	return pb.IMClientType_ClientTypeDefault
}

func responseError(res http.ResponseWriter, code int, msg string) {
	data := map[string]interface{}{
		"code": code,
		"msg":  msg,
	}
	buffer, err := json.Marshal(data)
	if err != nil {
		log.L.Warn(err.Error())
		return
	}

	res.WriteHeader(http.StatusBadRequest)
	_, _ = res.Write(buffer)
}
