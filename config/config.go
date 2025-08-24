package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Viewer Viewer
	SystemConfig SystemConfig
}


type Viewer struct {
	Title string
	Description  string
	Logo  string
	Navigation  []string
	Bilibili string
	Avatar string
	UserName string
	UserDesc string
}
type SystemConfig struct {
	AppName             string
	Version             float32
	CurrentDir          string
	CdnURL string
	QiniuAccessKey string
	QiniuSecretKey string
	Valine bool
	ValineAppid string
	ValineAppkey string
	ValineServerURL string
}


var Cfg *tomlConfig

func init() {
	Cfg = new(tomlConfig)
	Cfg.SystemConfig.AppName = "sorcerer-go-blog"
	Cfg.SystemConfig.Version = 1.0
	Cfg.SystemConfig.CurrentDir, _ = os.Getwd()
	_, err := toml.DecodeFile("config/config.toml", Cfg)
	if err != nil {
		panic(err)
	}
}