package main

import (
	"quantos-protocol/pkg/cfg"
	_ "quantos-protocol/pkg/cfg"
)

func init() {
	cfg.WriteConfig(cfg.ConfigPath, new(cfg.ChainCfg))
}
