package redi

import (
	"log"
	"strings"
	"time"

	"shopee/usermanager/usercacheproxyserver/g"
	pbredis "shopee/usermanager/usercacheproxyserver/proto"
	"github.com/garyburd/redigo/redis"
)

var RedisConnPool *redis.Pool

func Init() {
	redisConfig := g.Config().RedisMaster
	auth, addr := formatRedisAddr(redisConfig.Addr)
	log.Printf("redis config: %v|%v", auth, addr)
	RedisConnPool = &redis.Pool{
		MaxIdle:     redisConfig.MaxIdle,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				log.Fatalln("dail redis fail|", err)
				return nil, err
			}
			if auth != "" {
				if _, err := c.Do("AUTH", auth); err != nil {
					_ = c.Close()
					log.Fatalln("auth redis fail|", err)
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: PingRedis,
	}

	log.Println("open redis success")
}

func formatRedisAddr(addrConfig string) (string, string) {
	if redisAddr := strings.Split(addrConfig, "@"); len(redisAddr) == 1 {
		return "", redisAddr[0]
	} else {
		return strings.Join(redisAddr[0:len(redisAddr)-1], "@"), redisAddr[len(redisAddr)-1]
	}
}

func PingRedis(c redis.Conn, t time.Time) error {
	_, err := c.Do("ping")
	if err != nil {
		log.Println("[ERROR] ping redis fail", err)
	}
	return err
}


func SetUserToken(in  *pbredis.SetUserTokenRequest) error {
	rc := RedisConnPool.Get()
	defer rc.Close()

	_, err := rc.Do("HSet", in.GetName(), in.GetToken(), in.GetExpiretime())
   	if err != nil {
      		log.Println("SetUserToken error|", err)
      		return err
   	}

        return nil 
}

func GetUserToken(in  *pbredis.GetUserTokenRequest) (*pbredis.GetUserTokenReply, error) {
        ret := pbredis.GetUserTokenReply{Errmsg: "ok", Retcode: 1}

	rc := RedisConnPool.Get()
	defer rc.Close()

	r, err := redis.Int(rc.Do("HGet", in.GetName(), in.GetToken()))
   	if err != nil {
      		log.Println("GetUserToken error|", err)
		ret.Retcode = 2
                ret.Errmsg = "error"
      		return &ret, err
   	}

	ret.Expiretime = int32(r)
        return &ret, nil
}
