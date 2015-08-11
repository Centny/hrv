package main

import (
	"fmt"
	"github.com/Centny/gwf/netw/hrv"
	"github.com/Centny/gwf/pool"
	"github.com/Centny/gwf/util"
	"os"
	"runtime"
	"strconv"
	"time"
)

var ef func(c int) = os.Exit

func usage() {
	fmt.Println(`Usage: hrv options
  -f server configure file, running server mode.
  -s HRV server host, running client mode.
  -base reverse client base http url
  -token reverse client login token
  -name reverse client login name
  -alias reverse client login alias
  -hb 30 the HB time,default 30s
  -l if show the log
  -h show this.
		`)
}
func main() {
	var f, s string
	var base, token, name, alias string
	var log bool = false
	var hb int = 30
	olen := len(os.Args)
	for i := 1; i < olen; i++ {
		switch os.Args[i] {
		case "-f":
			if i < olen-1 {
				f = os.Args[i+1]
			}
		case "-s":
			if i < olen-1 {
				s = os.Args[i+1]
				i++
			}
		case "-base":
			if i < olen-1 {
				base = os.Args[i+1]
				i++
			}
		case "-token":
			if i < olen-1 {
				token = os.Args[i+1]
				i++
			}
		case "-name":
			if i < olen-1 {
				name = os.Args[i+1]
				i++
			}
		case "-alias":
			if i < olen-1 {
				alias = os.Args[i+1]
				i++
			}
		case "-hb":
			if i < olen-1 {
				t, err := strconv.ParseInt(os.Args[i+1], 10, 32)
				if err != nil {
					fmt.Println(err.Error())
					ef(1)
					return
				}
				hb = int(t)
				i++
			}
		case "-l":
			log = true
		case "-h":
			usage()
			ef(1)
			return
		}
	}
	runtime.GOMAXPROCS(util.CPU())
	if len(s) > 0 {
		RunHrvC(s, base, token, name, alias, log, hb)
	} else if len(f) > 0 {
		RunHrvS(f)
	} else {
		usage()
	}
}

var hs *hrv.CfgSrvH
var hc *hrv.HrvC

//run HRV server.
func RunHrvS(f string) {
	var err error
	bp := pool.NewBytePool(8, 1024000)
	hs, err = hrv.NewCfgSrvH(bp, f)
	if err != nil {
		panic(err.Error())
	}
	hs.Run()
}

//run HRV client
func RunHrvC(addr, base, token, name, alias string, log bool, hb int) {
	bp := pool.NewBytePool(8, 1024000)
	hc = hrv.NewHrvC_j(bp, addr, base)
	hc.ShowLog = log
	hc.Token = token
	hc.Name = name
	hc.Alias = alias
	hc.Start()
	go func() {
		for {
			hc.HB()
			time.Sleep(time.Duration(hb) * time.Second)
		}
	}()
	hc.Wait()
}
