package main

import (
	"flag"
	"plabfootball/cmd/app"
	"plabfootball/config"
)

// 디버깅용 절대경로 in MAC OS
//var abPath = "/Users/jeongseong-won/GolandProjects/plabfootball/cmd/config.toml"

// 디버깅용 절대경로 in Windows OS
var abPath = "/Users/lovet/GolandProjects/plabfootball/cmd/config.toml"

// 터미널용 상대경로
// var rePath = "./config.toml"

var pathFlag = flag.String("config", abPath, "you need to set toml path")

func main() {
	flag.Parse()
	// config 파일 가져옴.
	c := config.NewConfig(*pathFlag)
	// 앱 세팅.
	a := app.NewApp(c)
	//서버 시작.
	a.Router.Engin.Run(c.Network.Port)
}
