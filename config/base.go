package config

import (
	"github.com/yangioc/bk_pack/yamlloader"
)

var EnvInfo *Env

type Env struct {
	Service Service `yamgl:"service,omitempty"`
	Nats    Nats    `yaml:"nats,omitempty"`
	Log     Log     `yaml:"log,omitempty"`
	NodeNum int64   `yaml:"nodeNum,omitempty"` // TODO: 所有服務唯一值調整到 redis或DB 在服務啟動後取得, 或是調整到通用 package 設定成 const 但會造成擴展不方便
	Arango  Arango  `yaml:"arango,omitempty"`
}

type Service struct {
	DBA DBA `yaml:"dba,omitempty"`
}

type DBA struct {
	Addr     string `yaml:"addr,omitempty"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
}

type Nats struct {
	Addr     string `yaml:"addr,omitempty"`
	UserName string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
}

type Log struct {
	Level string `yaml:"level,omitempty"`
}

type Arango struct {
	Addr     string `yaml:"addr,omitempty"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
	DataBase string `yaml:"database,omitempty"`
}

func Init(path string) error {
	if err := yamlloader.LoadYaml(path, &EnvInfo); err != nil {
		// mylog.Errorf("[Config][Init] load Error err: %v", err)
		return err
	}
	return nil
}
