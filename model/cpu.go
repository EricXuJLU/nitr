package model

import (
	"encoding/json"
	"github.com/bitcav/nitr-core/cpu"
	"gorm.io/gorm"
)

type CPU struct {
	gorm.Model
	HostName   string  `json:"hostName"`
	Vendor     string  `json:"vendor"`
	CPUModel   string  `json:"Model"`
	Cores      uint32  `json:"cores"`
	Threads    uint32  `json:"threads"`
	ClockSpeed float64 `json:"clockSpeed"`
	Usage      float64 `json:"usage"`
}

func ParseCPUInfo(src cpu.CPU) (dst CPU) {
	JsonConvert(src, &dst)
	return
}

func JsonConvert(src interface{}, dst interface{}) {
	by, err := json.Marshal(src)
	if err != nil {
		panic("Json Marshal Error")
	}
	err = json.Unmarshal(by, dst)
	if err != nil {
		panic("Json Unmarshal Error")
	}
}
