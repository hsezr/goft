package goft

import (
	"log"

	"github.com/go-yaml/yaml"
)

type ServerConfig struct {
	Port int32
	Name string
}

type UserConfig map[interface{}]interface{}

type SysConfig struct {
	Server *ServerConfig
	Config UserConfig
}

func NewSysConfig() *SysConfig {
	return &SysConfig{Server: &ServerConfig{Port: 8080, Name: "myweb"}}
}

func InitConfig() *SysConfig {
	config := NewSysConfig()
	if b := LoadConfigFile(); b != nil {
		err := yaml.Unmarshal(b, config)
		if err != nil {
			log.Fatal(err)
		}
	}
	return config
}

func GetConfigValue(m UserConfig, prefix []string, idx int) interface{} {
	key := prefix[idx]
	if v, ok := m[key]; ok {
		if idx == len(prefix)-1 {
			return v
		} else {
			idx = idx + 1
			if mv, ok := v.(UserConfig); ok {
				return GetConfigValue(mv, prefix, idx)
			} else {
				return nil
			}
		}
	}

	return nil
}
