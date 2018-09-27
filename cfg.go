package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// ------------------------------------------------------------
// LOAD

func LoadCfgFromArgs() (Cfg, error) {
	if len(os.Args) < 2 {
		return Cfg{}, nil
	}
	return LoadCfgFromFile(os.Args[1])
}

func LoadCfgFromFile(filename string) (Cfg, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return Cfg{}, err
	}
	cfg := Cfg{}
	err = json.Unmarshal(file, &cfg)
	return cfg, err
}

// ------------------------------------------------------------
// CFG

type Cfg struct {
	Port int            `json:"port,omitempty"`
	Cmds map[string]Cmd `json:"cmds,omitempty"`
}

func (c Cfg) GetCmd(name string) (Cmd, error) {
	if c.Cmds == nil {
		return Cmd{}, errors.New("No commands")
	}
	ans, ok := c.Cmds[name]
	if !ok {
		return Cmd{}, errors.New("No command")
	}
	return ans.copy(), nil
}

// ------------------------------------------------------------
// CMD

type Cmd struct {
	Filename string    `json:"filename,omitempty"`
	Args     []string  `json:"args,omitempty"`
	Replaces []Replace `json:"replace,omitempty"`
}

func (c Cmd) copy() Cmd {
	ans := Cmd{Filename: c.Filename, Replaces: c.Replaces}
	for _, a := range c.Args {
		ans.Args = append(ans.Args, a)
	}
	return ans
}

// ------------------------------------------------------------
// REPLACE

type Replace struct {
	Context string `json:"context,omitempty"`
	Src     string `json:"src,omitempty"`
	Dst     string `json:"dst,omitempty"`
}
