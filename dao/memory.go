package dao

import "github.com/bitcav/nitr/model"

func CreateMemoryInfo(memory *model.Memory) (uint, error) {
	err := defaultDB.Create(memory).Error
	return memory.ID, err
}
