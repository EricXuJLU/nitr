package dao

import "github.com/bitcav/nitr/model"

func CreateHostInfo(host *model.Host) (uint, error) {
	err := defaultDB.Create(host).Error
	return host.ID, err
}
