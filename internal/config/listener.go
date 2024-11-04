package config

import "errors"

type TypeListen string

const (
	HTTP TypeListen = "http"
	TCP  TypeListen = "tcp"
)

type listenerConfig struct {
	TypeListen TypeListen `yaml:"type-listen" default:"http"`
	Port       int        `yaml:"port" default:"9595"`
}

func (lsnrCfg *listenerConfig) validate() error {
	if lsnrCfg.TypeListen == "" {
		return errors.New("listenerConfig : type listener is empty")
	}
	return nil
}
