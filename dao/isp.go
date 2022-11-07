package dao

import "github.com/bitcav/nitr/model"

func CreateIspInfo(isp *model.Isp) (uint, error) {
	err := defaultDB.Create(isp).Error
	return isp.ID, err
}
