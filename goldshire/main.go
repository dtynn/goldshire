package main

import (
	"flag"
	"os"

	"github.com/labstack/echo"

	"github.com/dtynn/goldshire"
)

// flags
var cfgPath string

func main() {
	flag.StringVar(&cfgPath, "config", "default.json", "path to config file")

	cfg, err := goldshire.GetConfig(cfgPath)
	if err != nil {
		goldshire.GetLogger().Errorf("fail to load config at %s", err)
		os.Exit(2)
	}

	store := goldshire.NewGitlabStore(cfg.Gitlab)

	e := echo.New()
	mux := e.Group("")
	goldshire.RegisterApis(mux, cfg, store)

	e.Run(cfg.Listen)
}
