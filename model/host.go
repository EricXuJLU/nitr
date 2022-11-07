package model

import (
	"github.com/bitcav/nitr-core/host"
	"gorm.io/gorm"
)

type Host struct {
	gorm.Model
	Name     string `json:"name"`
	OS       string `json:"os"`
	Platform string `json:"platform"`
	Arch     string `json:"arch"`
	Uptime   uint64 `json:"uptime"`
}

func ParseHostInfo(src host.HostInfo) (dst Host) {
	JsonConvert(src, &dst)
	return
}
