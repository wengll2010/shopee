package main

import (
	"log"
	"flag"
	"shopee/usermanager/userdbproxyserver/g"
	"shopee/usermanager/userdbproxyserver/db"
	"shopee/usermanager/userdbproxyserver/rpc"
	"shopee/usermanager/userdbproxyserver/http"
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

	// db Init
	db.Init()

	// rpc
	rpc.Start()

	// http
	http.Start()

	select {}
}
