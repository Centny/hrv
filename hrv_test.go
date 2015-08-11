package main

import (
	"os"
	"testing"
	"time"
)

func TestHrv(t *testing.T) {
	ef = func(code int) {}
	os.Args = []string{"hrv", "-f", "conf.properties"}
	go main()
	time.Sleep(300 * time.Millisecond)
	os.Args = []string{"hrv", "-s", "127.0.0.1:8234", "-base", "http://localhost",
		"-token", "token", "-name", "name", "-alias", "alias", "-hb", "10", "-l"}
	go main()
	time.Sleep(300 * time.Millisecond)

	//
	os.Args = []string{"hrv"}
	main()
	os.Args = []string{"hrv", "-h"}
	main()
	os.Args = []string{"hrv", "-hb", "sss"}
	main()
	run_err()
}

func run_err() {
	defer func() {
		recover()
	}()
	os.Args = []string{"hrv", "-f", "sss"}
	main()
}
