package dao

import "github.com/bitcav/nitr/model"

func CreateDiskInfo(disk *model.Disk) (uint, error) {
	err := defaultDB.Create(disk).Error
	return disk.ID, err
}
