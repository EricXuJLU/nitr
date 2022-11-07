package model

import (
	"github.com/bitcav/go-memdev"
	"gorm.io/gorm"
)

type Memory struct {
	gorm.Model
	HostName     string `json:"hostName"`
	Bank         string `json:"bank"`
	Size         int    `json:"size"`
	Unit         string `json:"unit"`
	Type         string `json:"type"`
	FormFactor   string `json:"formFactor"`
	Manufacturer string `json:"manufacturer"`
	Speed        int    `json:"speed"`
	DataWidth    int    `json:"dataWidth"`
	TotalWidth   int    `json:"totalWidth"`
}

func ParseMemoryInfo(src memdev.Memory) (dst Memory) {
	JsonConvert(src, &dst)
	return
}
