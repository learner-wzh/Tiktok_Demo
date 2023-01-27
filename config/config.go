package config

import (
	"github.com/BurntSushi/toml"
	"strings"
)

type Mysql struct {
	Host      string
	Port      int
	Database  string
	Username  string
	Password  string
	Charset   string
	ParseTime bool `toml:"parse_time"`
	Loc       string
}

type Server struct {
	IP   string
	Port int
}

type Path struct {
	FfmpegPath       string `toml:"ffmpeg_path"`
	StaticSourcePath string `toml:"static_source_path"`
}

type Config struct {
	DB     Mysql `toml:"mysql"`
	Server `toml:"server"`
	Path   `toml:"path"`
}

var Info Config

// 包初始化加载时候会调用的函数
func init() {
	if _, err := toml.DecodeFile("config/config.toml", &Info); err != nil {
		panic(err)
	}
	//去除左右的空格
	strings.Trim(Info.Server.IP, " ")
	strings.Trim(Info.DB.Host, " ")
}