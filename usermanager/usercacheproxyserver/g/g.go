package g

import (
	"log"
	"runtime"
)

var (
	Version    string
)

func VersionMsg() string {
	return Version 
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
