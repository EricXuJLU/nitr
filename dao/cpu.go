package dao

import (
	"github.com/bitcav/nitr/model"
)

func CreateCPUInfo(cpu *model.CPU) (uint, error) {
	err := defaultDB.Create(cpu).Error
	return cpu.ID, err
}
