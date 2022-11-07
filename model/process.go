package model

import (
	"github.com/bitcav/nitr-core/process"
	"gorm.io/gorm"
)

type Process struct {
	gorm.Model
	Pid  int    `json:"pid"`
	Name string `json:"name"`
}

func ParseProcessInfo(src process.Process) (dst Process) {
	JsonConvert(src, &dst)
	return
}
