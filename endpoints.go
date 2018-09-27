package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

// ------------------------------------------------------------
// NOT-FOUND

func epNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

// ------------------------------------------------------------
// RUN

func epRun(w http.ResponseWriter, r *http.Request) {
	name := getUrlParam(r, "cmd")
	cmd_cfg, err := global_cfg.GetCmd(name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cmd_cfg = replaceHeaders(cmd_cfg, r)
	fmt.Println("run cmd", name, "cfg", cmd_cfg)
	var args []string
	if cmd_cfg.Args != nil {
		args = cmd_cfg.Args
	}

	cmd := exec.Command(cmd_cfg.Filename, args...)
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		fmt.Println("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, response{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
}

// ------------------------------------------------------------
// REPLACE

func replaceHeaders(cmd Cmd, req *http.Request) Cmd {
	if len(cmd.Replaces) < 1 {
		return cmd
	}
	for _, r := range cmd.Replaces {
		if r.Context == "header" {
			cmdReplace(r.Src, req.Header.Get(r.Dst), &cmd)
		}
	}
	return cmd
}

func cmdReplace(src, dst string, cmd *Cmd) {
	cmd.Filename = strings.Replace(cmd.Filename, src, dst, -1)
	if len(cmd.Args) < 1 {
		return
	}
	for i, a := range cmd.Args {
		cmd.Args[i] = strings.Replace(a, src, dst, -1)
	}
}

// ------------------------------------------------------------
// SUPPORT

// Utility to get a chi param and unencode it.
// Return blank if something goes wrong.
func getUrlParam(r *http.Request, key string) string {
	_path := chi.URLParam(r, key)
	path, err := url.QueryUnescape(_path)
	if err != nil {
		return ""
	}
	return path
}

func epFakeFmt() {
	fmt.Println("")
}

type response struct {
	Error string `json:"error,omitempty"`
}
