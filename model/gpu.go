package model

import (
	"github.com/bitcav/nitr-core/gpu"
	"gorm.io/gorm"
)

type GPU struct {
	gorm.Model
	HostName string `json:"hostName"`
	Brand    string `json:"brand"`
	GPUModel string `json:"model"`
}

func ParseGPUInfo(src gpu.GPU) (dst GPU) {
	JsonConvert(src, &dst)
	return
}
