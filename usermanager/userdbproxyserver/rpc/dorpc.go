package rpc

import (
        "context"
        "log"
	"net"
        //"os"
        //"time"

	"shopee/usermanager/userdbproxyserver/db"
        "google.golang.org/grpc"
        pbdb "shopee/usermanager/userdbproxyserver/proto"
)

const (
        port = ":50151"
)

// server is used to implement proto.DBUserServer.
type server struct {
        pbdb.UnimplementedDBUserServer
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
        pbdb.RegisterDBUserServer(s, &server{})
        if err := s.Serve(lis); err != nil {
                log.Fatalf("startRpcServer|%v|failed to serve: %v", port, err)
        }

	log.Printf("startRpcServer|%v|success \n", port)
}

func (s *server) AddDBUser(ctx context.Context, in *pbdb.AddDBUserRequest) (*pbdb.AddDBUserReply, error) {
        log.Printf("AddDBUser.Received|%v \n", in)

	err := db.AddDBUser(in)
	if err != nil {
        	return &pbdb.AddDBUserReply{Errmsg: err.Error(), Retcode: 1}, err
	} else {
        	return &pbdb.AddDBUserReply{Errmsg: "ok", Retcode: 1}, nil
	}
}

func (s *server) GetDBUser(ctx context.Context, in *pbdb.GetDBUserRequest) (*pbdb.GetDBUserReply, error) {
        log.Printf("GetDBUser.Received|%v|%v \n", in.GetName())

	r, err := db.GetDBUser(in.GetName())
	if err != nil {
		r.Retcode = 2
		r.Errmsg = err.Error()
	} else {
		r.Retcode = 1
		r.Errmsg = "OK"
	}

	return r, err
}

func (s *server) UpdateDBUser(ctx context.Context, in *pbdb.UpdateDBUserRequest) (*pbdb.UpdateDBUserReply, error) {
        log.Printf("UpdateDBUser.Received|%v|%v|%v \n", in.GetName(), in.GetNick(), in.GetPicture())

	err := db.UpdateDBUser(in)
	if err != nil {
        	return &pbdb.UpdateDBUserReply{Errmsg: err.Error(), Retcode: 1}, err
	} else {
        	return &pbdb.UpdateDBUserReply{Errmsg: "ok", Retcode: 1}, nil
	}
}
