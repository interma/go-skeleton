package main

import (
	// standard library
	"fmt"
	"net/http"
	"flag"

	// vendor package
	"github.com/beego/mux"
	"github.com/golang/glog"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"

	// my own package
	"github.com/interma/go-skeleton/pkg/add"
	"github.com/interma/go-skeleton/pkg/version"
)

// run a basic http server
func runWebserver() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	mx := mux.New()
	mx.Handler("GET", "/", http.FileServer(http.Dir(".")))
	sugar.Fatal(http.ListenAndServe("127.0.0.1:8080", mx))
}

func main() {
	flag.Parse()

	ver := version.Version
	//fmt.Println(ver)
	logrus.Infof("Version: %v", ver)
	logrus.WithFields(logrus.Fields{
		"Version": ver,
	}).Info("A log example")

	//another log
	//go run main.go -log_dir=./log/ -v=4 -alsologtostderr
	glog.Info("hello, glog")
	glog.V(4).Info("another V log, by glog!")
	glog.Flush()

	sum := add.Add(1, 2)
	fmt.Println(sum)
}
