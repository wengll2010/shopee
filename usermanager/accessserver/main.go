package main

import (
	"log"
	"flag"
	"shopee/usermanager/accessserver/g"
	"shopee/usermanager/accessserver/rpc"
	"shopee/usermanager/accessserver/http"
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

	// rpc
	rpc.Start()

	// http
	http.Start()

	select {}
}
