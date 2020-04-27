package http

import (
	"shopee/usermanager/accessserver/rpc"
	"net/http"
	"encoding/json"
	pb "shopee/usermanager/userserver/proto"
	"log"
)

func userCommonRoutes() {

	http.HandleFunc("/loginuser", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var x *pb.LoginUserRequest
		err := decoder.Decode(&x)
		if err != nil {
			log.Printf("loginuser err: %v \n", err)
			RenderUserJson(w, "", &pb.LoginUserReply{Errmsg: "error", Retcode: 2, Token: ""})
		} else {	
			rp, _ := rpc.LoginUser(x)
			token := rp.GetToken()
			rp.Token = ""
			RenderUserJson(w, token, rp)
		}
	})

	http.HandleFunc("/getuser", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("getuser body: %v", r.Body)
		decoder := json.NewDecoder(r.Body)
		var x *pb.GetUserRequest
		err := decoder.Decode(&x)
		if err != nil {
			log.Printf("getuser err: %v \n", err)
			RenderUserJson(w, "", &pb.GetUserReply{Errmsg: "error", Retcode: 2})
		} else {	
			x.Token = r.Header.Get("Authorization")
			rp, _ := rpc.GetUser(x) 
			RenderUserJson(w, "", rp)
		}
	})

	http.HandleFunc("/updateuser", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var x *pb.UpdateUserRequest
		err := decoder.Decode(&x)
		if err != nil {
			log.Printf("updateuser err: %v \n", err)
			RenderUserJson(w, "", &pb.UpdateUserReply{Errmsg: "error", Retcode: 2})
		} else {	
			x.Token = r.Header.Get("Authorization")
			rp, _ := rpc.UpdateUser(x)
			RenderUserJson(w, "", rp)
		}
	})
}
