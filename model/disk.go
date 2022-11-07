package model

import (
	"github.com/bitcav/nitr-core/disk"
	"gorm.io/gorm"
)

type Disk struct {
	gorm.Model
	HostName   string  `json:"hostName"`
	Mountpoint string  `json:"mountPoint"`
	Free       uint64  `json:"free"`
	Size       uint64  `json:"size"`
	Used       uint64  `json:"used"`
	Percent    float64 `json:"percent"`
}

func ParseDiskInfo(src disk.Disk) (dst Disk) {
	JsonConvert(src, &dst)
	return
}
