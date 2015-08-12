Http Reverse Require
======
hrv command provide the http reverse connection to require inner netwok http server. 

it map http://\<public server address\>/\<prefix\>/\<url path\> to http://\<inner server address\>/\<url path\>

![hrv](https://raw.githubusercontent.com/Centny/hrv/master/hrv.png)

## Install
```
go get github.com/Centny/gwf
go get github.com/Centny/hrv
```

## Running Server

* edit configure file `conf.properties`

```
#the web server prefix
PRE=
#if show debug log
LOG=1
#the header name to transfter to reverse client.
HEADERS=
#the web base directory
WWW=.
#the port for HRV
ADDR=:8234
#the port for http server
HTTP=:8123
#
#prefix A- to configure arguments to external transfter arguments,
#it will append to require arguments and send to reverse client.
#
A-a1=1
#prefix A- to configure arguments to external transfter headers
#it will append to require headers and send to reverse client.
H-h1=2
#
#prefix P- to confgireu pattern to match uri.
#the uri which match the pattern list will reverse to client.
P-p1=3
#
#prefix T- to configure user login tokan and name
T-name=token


```

* run server

```
hrv -f conf.properties
```

## Running Standalone Client

```
hrv -s 127.0.0.1:8234 -base <server base which will be reversed> -token <login token> -name <login name> -alias <login alias> -hb 60 -l
```
run `hrf -h` to show help.


## Running Emmed Client

```
import "github.com/Centny/gwf/netw/hrv"
....
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
```
see doc on `github.com/Centny/gwf/netw/hrv` for detail.

## Example

![hrv](https://raw.githubusercontent.com/Centny/hrv/master/hrv.png)