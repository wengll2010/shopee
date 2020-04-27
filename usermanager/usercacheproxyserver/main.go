package main

import (
	"log"
	"flag"
	"shopee/usermanager/usercacheproxyserver/g"
	"shopee/usermanager/usercacheproxyserver/redis"
	"shopee/usermanager/usercacheproxyserver/rpc"
	"shopee/usermanager/usercacheproxyserver/http"
	"os"
)

func main() {
	g.Version = Version

	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	if *version {
		log.Printf("version %s\n", Version)
		os.Exit(0)
	}
	
	// global config
	g.ParseConfig(*cfg)

	// redis Init
	redi.Init()

	// rpc
	rpc.Start()

	// http
	http.Start()

	select {}
}
