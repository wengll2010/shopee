package rpc

import (
        "context"
        "log"
        //"os"
        "time"

        "google.golang.org/grpc"
        pb "shopee/usermanager/userserver/proto"
)

var (
	cc pb.UserClient
)

const (
        address     = "localhost:50001"
)

func Start() {
	go startRpcClient()
}

func startRpcClient() {
        // Set up a connection to the server.
        conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
        if err != nil {
		log.Fatalf("startRpcClient|%v|did not connect: %v", address, err)
        }
        cc = pb.NewUserClient(conn)

	log.Printf("startRpcClient|%v|success \n", address)
}

func LoginUser(x *pb.LoginUserRequest) (*pb.LoginUserReply, error) {
        ctx, _ := context.WithTimeout(context.Background(), time.Second)

        r, err := cc.LoginUser(ctx, x)
        if err != nil {
		log.Printf("LoginUser|%v|%v \n", x, err)
                return &pb.LoginUserReply{Errmsg: "error", Retcode: 2}, err
        }

        log.Printf("LoginUser|%v \n", r)
	return r, nil 
}

func GetUser(x *pb.GetUserRequest) (*pb.GetUserReply, error) {
	log.Printf("GetUser|%v", x)
        ctx, _ := context.WithTimeout(context.Background(), time.Second)

        r, err := cc.GetUser(ctx, x)
        if err != nil {
		log.Printf("GetUser|%v|%v \n", x, err)
                return &pb.GetUserReply{Errmsg: "error", Retcode: 2}, err
        }

        log.Printf("GetUser|%v \n", r)
	return r, nil 
}

func UpdateUser(x *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
        ctx, _ := context.WithTimeout(context.Background(), time.Second)

        r, err := cc.UpdateUser(ctx, x) 
        if err != nil {
		log.Printf("UpdateUser|%v|%v \n", x, err)
                return &pb.UpdateUserReply{Errmsg: "error", Retcode: 2}, err
        }

        log.Printf("UpdateUser|%v \n", r)
	return r, nil
}
