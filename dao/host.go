package dao

import "github.com/bitcav/nitr/model"

func CreateHostInfo(host *model.Host) (uint, error) {
	err := defaultDB.Create(host).Error
	return host.ID, err
}

func DeleteHostByName(name string) error {
	err := defaultDB.Model(&model.Host{}).Where("name=?", name).Delete(&model.Host{}).Error
	return err
}
