package server

import (
	"github.com/bhaney/rtsp-simple-server/internal/core"
)

func New(configArgs []string) (*core.Core, error) {
	return core.New(configArgs)
}
