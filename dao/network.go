package dao

import "github.com/bitcav/nitr/model"

func CreateNetworkInfo(network *model.Network) (uint, error) {
	err := defaultDB.Create(network).Error
	return network.ID, err
}
