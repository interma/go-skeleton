package main

import (
	// standard library
	"fmt"
	"net/http"

	// vendor package
	"github.com/beego/mux"
	"go.uber.org/zap"

	// my own package
	"github.com/interma/go-skeleton/pkg/add"
	"github.com/interma/go-skeleton/pkg/version"
)

// run a basic http server
func run_webserver() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	mx := mux.New()
	mx.Handler("GET", "/", http.FileServer(http.Dir(".")))
	sugar.Fatal(http.ListenAndServe("127.0.0.1:8080", mx))
}

func main() {
	ver := version.Version
	fmt.Println(ver)

	sum := add.Add(1, 2)
	fmt.Println(sum)
}
