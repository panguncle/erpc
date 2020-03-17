package main

import (
	"context"
	"time"

	"github.com/henrylee2cn/erpc/v6"
)

//go:generate go build $GOFILE

func main() {
	// PReader: 也不知道干嘛的, 先跳过了
	defer erpc.SetLoggerLevel("INFO")()
	// PReader: 从README粘过来的注释: Use peer to provide the same API package for the server and client
	cli := erpc.NewPeer(erpc.PeerConfig{PrintDetail: true})
	sess, stat := cli.Dial(":9090")
	if !stat.OK() {
		erpc.Fatalf("%v", stat)
	}

	var result string
	sess.Call("/test/ok", "test1", &result)
	erpc.Infof("test sync1: %v", result)
	result = ""
	stat = sess.Call("/test/timeout", "test2", &result).Status()
	erpc.Infof("test sync2: server context timeout: %v", stat)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	result = ""
	callCmd := sess.AsyncCall(
		"/test/timeout",
		"test3",
		&result,
		make(chan erpc.CallCmd, 1),
		erpc.WithContext(ctx),
	)
	select {
	case <-callCmd.Done():
		cancel()
		erpc.Infof("test async1: %v", result)
	case <-ctx.Done():
		erpc.Warnf("test async1: client context timeout: %v", ctx.Err())
	}

	time.Sleep(time.Second * 6)
	result = ""
	stat = sess.Call("/test/ok", "test4", &result).Status()
	erpc.Warnf("test sync3: disconnect due to server session timeout: %v", stat.Cause())

	sess, stat = cli.Dial(":9090")
	if !stat.OK() {
		erpc.Fatalf("%v", stat)
	}
	sess.AsyncCall(
		"/test/break",
		nil,
		nil,
		make(chan erpc.CallCmd, 1),
	)
}
