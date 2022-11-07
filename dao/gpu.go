package dao

import "github.com/bitcav/nitr/model"

func CreateGPUInfo(gpu *model.GPU) (uint, error) {
	err := defaultDB.Create(gpu).Error
	return gpu.ID, err
}
