package hy2

import (
	"github.com/EbadiDev/Arch-Server/conf"
	vCore "github.com/EbadiDev/Arch-Server/core"
	"go.uber.org/zap"
)

var _ vCore.Core = (*Hysteria2)(nil)

type Hysteria2 struct {
	Hy2nodes map[string]Hysteria2node
	Auth     *Arch-Server
	Logger   *zap.Logger
}

func init() {
	vCore.RegisterCore("hysteria2", New)
}

func New(c *conf.CoreConfig) (vCore.Core, error) {
	loglever := "error"
	if c.Hysteria2Config.LogConfig.Level != "" {
		loglever = c.Hysteria2Config.LogConfig.Level
	}
	log, err := initLogger(loglever, "console")
	if err != nil {
		return nil, err
	}
	return &Hysteria2{
		Hy2nodes: make(map[string]Hysteria2node),
		Auth: &Arch-Server{
			usersMap: make(map[string]int),
		},
		Logger: log,
	}, nil
}

func (h *Hysteria2) Protocols() []string {
	return []string{
		"hysteria2",
	}
}

func (h *Hysteria2) Start() error {
	return nil
}

func (h *Hysteria2) Close() error {
	for _, n := range h.Hy2nodes {
		err := n.Hy2server.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *Hysteria2) Type() string {
	return "hysteria2"
}
