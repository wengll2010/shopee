package rpc

import (
        "context"
        "log"
	"net"
        //"os"
        //"time"

	"shopee/usermanager/usercacheproxyserver/redis"
        "google.golang.org/grpc"
        pbredis "shopee/usermanager/usercacheproxyserver/proto"
)

const (
        port = ":50651"
)

// server is used to implement proto.CacheUserServer.
type server struct {
        pbredis.UnimplementedCacheUserServer
}

func Start() {
	go startRpcServer()
}

func startRpcServer() {
        lis, err := net.Listen("tcp", port)
        if err != nil {
                log.Fatalf("startRpcServer|%v|failed to listen: %v", port, err)
        }
        s := grpc.NewServer()
        pbredis.RegisterCacheUserServer(s, &server{})
        if err := s.Serve(lis); err != nil {
		log.Fatalf("startRpcServer|%v|failed to serve: %v", port, err)
        }

	log.Printf("startRpcServer|%v|success \n", port)
}

func (s *server) SetUserToken(ctx context.Context, in *pbredis.SetUserTokenRequest) (*pbredis.SetUserTokenReply, error) {
        log.Printf("SetUserToken.Received|%v|%v|%v \n", in.GetName(), in.GetToken(), in.GetExpiretime())
        return &pbredis.SetUserTokenReply{Errmsg: "ok", Retcode: 1}, nil
}

// GetUserToken implements proto.UserServer
func (s *server) GetUserToken(ctx context.Context, in *pbredis.GetUserTokenRequest) (*pbredis.GetUserTokenReply, error) {
        log.Printf("GetUserToken.Received|%v|%v \n", in.GetName(), in.GetToken())

	r, err := redi.GetUserToken(in)
	if err != nil {
		r.Retcode = 2
		r.Errmsg = err.Error()
	} else {
		r.Retcode = 1
		r.Errmsg = "OK"
	}

	return r, err
}
