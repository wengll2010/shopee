package rpc

import (
	"fmt"
        "context"
        "log"
	"net"
        //"os"
        "time"
	"crypto/md5"
    	"encoding/hex"
	"unicode"
	"errors"
	"math/rand"

        "google.golang.org/grpc"
        pb "shopee/usermanager/userserver/proto"
        pbdb "shopee/usermanager/userdbproxyserver/proto"
        pbcache "shopee/usermanager/usercacheproxyserver/proto"
)

var (
	ccdb pbdb.DBUserClient
	cccache pbcache.CacheUserClient
)

// 服务端监听端口
const (
        port = ":50001"
)

// server is used to implement usermanager.UserServer.
type server struct {
        pb.UnimplementedUserServer
}

const (
        addressdb     = "localhost:50151"
        addresscache  = "localhost:50651"
)

func Start() {
	go startDBRpcClient()
	go startCacheRpcClient()
	go startRpcServer()
}

func startDBRpcClient() {
        // Set up a connection to the server.
        conn, err := grpc.Dial(addressdb, grpc.WithInsecure(), grpc.WithBlock())
        if err != nil {
                log.Fatalf("startDBRpcClient|%v|did not connect: %v", addressdb, err)
        }
        ccdb = pbdb.NewDBUserClient(conn)

	log.Printf("startDBRpcClient|%v|success", addressdb)
}

func startCacheRpcClient() {
        // Set up a connection to the server.
        conn, err := grpc.Dial(addresscache, grpc.WithInsecure(), grpc.WithBlock())
        if err != nil {
                log.Fatalf("startCacheRpcClient|%v|did not connect: %v", addresscache, err)
        }
        cccache = pbcache.NewCacheUserClient(conn)

	log.Printf("startCacheRpcClient|%v|success", addresscache)
}

func startRpcServer() {
        lis, err := net.Listen("tcp", port)
        if err != nil {
		log.Fatalf("startRpcServer|%v|failed to listen: %v", port, err)
        }
        s := grpc.NewServer()
        pb.RegisterUserServer(s, &server{})
        if err := s.Serve(lis); err != nil {
                log.Fatalf("startRpcServer|%v|failed to serve: %v", port, err)
        }

	log.Printf("startRpcServer|%v|success", port)
}

func md5V(str string) string  {
	h := md5.New()
    	h.Write([]byte(str))
    	return hex.EncodeToString(h.Sum(nil))
}

func generatetoken(name string) string {
	md5str := fmt.Sprintf("shopee|%s", name)
	return md5V(md5str) 
} 

func IsLetterNumber(s string) bool {
	for _, r := range s {
		log.Println("TTTTTT:", r)
        	if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
            		return false
        	}
    	}

   	 return true
}

//rpc server
func (s *server) LoginUser(ctx context.Context, in *pb.LoginUserRequest) (*pb.LoginUserReply, error) {
        log.Printf("LoginUser.Received|%v|%v \n", in.GetName(), in.GetPassword())

	// step0 判断name合法
	flag := IsLetterNumber(in.GetName())
	if !flag {
                log.Printf("LoginUser|%v \n", in.GetName())
        	return &pb.LoginUserReply{Errmsg: "error name", Retcode: 2}, errors.New("error name") 
	}
	
	// step1 判断密码
	ctxdb, _ := context.WithTimeout(context.Background(), time.Second)
        r, err := ccdb.GetDBUser(ctxdb, &pbdb.GetDBUserRequest{Name: in.GetName()})
        if err != nil {
                log.Printf("LoginUser.GetDBUser|%v|%v \n", in.GetName(), err)
        	return &pb.LoginUserReply{Errmsg: "error user", Retcode: 2}, err
        } else {
		md5pwd := generatetoken(in.GetPassword()) 
		if md5pwd != r.GetPassword() {
			log.Printf("LoginUser.GetDBUser|%v|%v \n", in, r)
        		return &pb.LoginUserReply{Errmsg: "error password", Retcode: 2}, err
		}
	}

	// step2 生成token
	ctxcache, _ := context.WithTimeout(context.Background(), time.Second)
	token := generatetoken(in.GetName())
        rc, err := cccache.SetUserToken(ctxcache, &pbcache.SetUserTokenRequest{Name: in.GetName(), Token: token})
        if err != nil {
                log.Printf("LoginUser.SetUserToken|%v|%v \n", in.GetName(), err)
        	return &pb.LoginUserReply{Errmsg: "token error", Retcode: 2}, err
	}

	log.Printf("LoginUser|%v|%v \n", in, rc)
        return &pb.LoginUserReply{Errmsg: "ok", Retcode: 1, Token: token}, nil
}

func (s *server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserReply, error) {
        log.Printf("GetUser.Received|%v|%v \n", in.GetName(), in.GetToken())

	// step0 判断name合法
	flag := IsLetterNumber(in.GetName())
	if !flag {
                log.Printf("LoginUser|%v \n", in.GetName())
        	return &pb.GetUserReply{Errmsg: "error name", Retcode: 2}, errors.New("error name") 
	}

	// step1 token验证
	ctxcache, _ := context.WithTimeout(context.Background(), time.Second)
        rc, err := cccache.GetUserToken(ctxcache, &pbcache.GetUserTokenRequest{Name: in.GetName(), Token: in.GetToken()})
        if err != nil {
                log.Printf("GetUser.GetUserToken|%v|%v|%v \n", in.GetName(), in.GetToken(), err)
        	return &pb.GetUserReply{Errmsg: "token error", Retcode: 2}, err
        } else {
		timenow := time.Now().Unix()
		if timenow > int64(rc.GetExpiretime()) {
                	log.Printf("GetUser.GetUserToken|%v|%v|%v \n", in.GetName(), in.GetToken(), rc.GetExpiretime())
        		return &pb.GetUserReply{Errmsg: "token expire", Retcode: 2}, err
		}
	}

	// step2 查询用户信息
	ctxdb, _ := context.WithTimeout(context.Background(), time.Second)
        r, err := ccdb.GetDBUser(ctxdb, &pbdb.GetDBUserRequest{Name: in.GetName()})
        if err != nil {
                log.Printf("GetUser.GetDBUser|%v|%v \n", in.GetName(), err)
        	return &pb.GetUserReply{Errmsg: "error", Retcode: 2}, err
        }

        log.Printf("GetUser|%v|%v \n", in.GetName(), r)
        return &pb.GetUserReply{Errmsg: "ok", Retcode: 1, Name: r.GetName(), Nick: r.GetNick(), Picture: r.GetPicture(), Regtime: r.GetRegtime()}, nil
}

func (s *server) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
        log.Printf("UpdateUser.Received|%v|%v|%v|%v \n", in.GetName(), in.GetToken(), in.GetNick(), in.GetPicture())

	// step0 判断name合法
	flag := IsLetterNumber(in.GetName())
	if !flag {
                log.Printf("LoginUser|%v \n", in.GetName())
        	return &pb.UpdateUserReply{Errmsg: "error name", Retcode: 2}, errors.New("error name") 
	}

	// step1 token验证
	ctxcache, _ := context.WithTimeout(context.Background(), time.Second)
        rc, err := cccache.GetUserToken(ctxcache, &pbcache.GetUserTokenRequest{Name: in.GetName(), Token: in.GetToken()})
        if err != nil {
                log.Printf("UpdateUser.GetUserToken|%v|%v|%v \n", in.GetName(), in.GetToken(), err)
        	return &pb.UpdateUserReply{Errmsg: "token error", Retcode: 2}, err
        } else {
		timenow := time.Now().Unix()
		if timenow > int64(rc.GetExpiretime()) {
                	log.Printf("UpdateUser.GetUserToken|%v|%v|%v \n", in.GetName(), in.GetToken(), rc.GetExpiretime())
        		return &pb.UpdateUserReply{Errmsg: "token expire", Retcode: 2}, err
		}
	}

	// step2 修改用户信息
	ctxdb, _ := context.WithTimeout(context.Background(), time.Second)
        r, err := ccdb.UpdateDBUser(ctxdb, &pbdb.UpdateDBUserRequest{Name: in.GetName(), Nick: in.GetNick(), Picture: in.GetPicture()})
        if err != nil {
                log.Printf("UpdateUser.UpdateDBUser|%v|%v|%v|%v \n", in.GetName(), in.GetNick(), in.GetPicture(), err)
        	return &pb.UpdateUserReply{Errmsg: "error", Retcode: 2}, err
        }

        log.Printf("GetUser|%v|%v \n", in, r)
        return &pb.UpdateUserReply{Errmsg: "ok", Retcode: 1}, nil
}

//test
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
func RandStringRunes() string {
    b := make([]rune, 10)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func TestAddUser(usernum int) {
        log.Printf("TestAddUser.Received|%d \n", usernum)
	rand.Seed(time.Now().UnixNano())

	ctxdb, _ := context.WithTimeout(context.Background(), 100*time.Second)
	regtime := int32(time.Now().Unix()) + 3600*24*7
	for i := 0; i <= usernum; i++ {
		name := RandStringRunes() 
		password := md5V(fmt.Sprintf("%s123456", name))   //密码都是123456，客户端传上来的password=md5(用户名+密码)
		md5pwd := generatetoken(password)
		nick := "testnick" 
		picture := "http://testwll.com/pic/1.jpg"
        	_, err := ccdb.AddDBUser(ctxdb, &pbdb.AddDBUserRequest{Name: name, Password: md5pwd, Nick: nick, Picture: picture, Regtime: regtime})
                log.Printf("TestAddUser.AddDBUser|%v|%v|%v|%v|%v|%v \n", name, password, md5pwd, nick, picture, err)
        }
}

