package dao

import "github.com/bitcav/nitr/model"

func CreateProcessInfo(process *model.Process) (uint, error) {
	err := defaultDB.Create(process).Error
	return process.ID, err
}
