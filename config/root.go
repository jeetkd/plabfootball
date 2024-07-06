package config

import (
	"github.com/naoina/toml"
	"os"
)

// Config 는 환경설정 객체
type Config struct {
	Network struct {
		Port string // :8080
	}

	Mongo struct {
		Db  string //db 이름
		Uri string //db 접속 uri
	}
}

// NewConfig 는 config 설정을 파싱해줍니다.
func NewConfig(path string) *Config {
	c := new(Config)

	if f, err := os.Open(path); err != nil {
		panic(err)
	} else if err = toml.NewDecoder(f).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}
}
