package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"time"
)

var (
	global_cfg = Cfg{} // This should be supplied to endpoints in a context, but this works for now
)

func main() {
	cfg, err := LoadCfgFromArgs()
	if err != nil {
		panic(err)
	}
	if len(cfg.Cmds) < 1 {
		panic("No commands. Specify a valid config file as the first argument.")
	}
	global_cfg = cfg

	mux := chi.NewRouter()
	mux.HandleFunc("/run/{cmd}", epRun)
	mux.NotFound(epNotFound)

	url := ":" + strconv.Itoa(cfg.Port)
	fmt.Println("Running at", url)
	server := &http.Server{Addr: url, Handler: mux}
	shutdownArgs := shutdownArgs{5 * time.Second, server}
	go gracefulShutdown(shutdownArgs)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Println(err)
	}
}
