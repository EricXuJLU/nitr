package dao

import "github.com/bitcav/nitr/model"

func CreateRAMInfo(ram *model.RAM) (uint, error) {
	err := defaultDB.Create(ram).Error
	return ram.ID, err
}
