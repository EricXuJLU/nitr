package model

import (
	"github.com/bitcav/nitr-core/ram"
	"gorm.io/gorm"
)

type RAM struct {
	gorm.Model
	Total uint64 `json:"total"`
	Free  uint64 `json:"free"`
	Usage uint64 `json:"usage"`
}

func ParseRAMInfo(src ram.RAM) (dst RAM) {
	JsonConvert(src, &dst)
	return
}
