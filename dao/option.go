package dao

import "github.com/bitcav/nitr/model"

func GetOption() model.Option {
	var opt model.Option
	defaultDB.First(&opt)
	return opt
}

func CreateOption(opt *model.Option) (uint, error) {
	err := defaultDB.Create(opt).Error
	return opt.ID, err
}

func UpdateOption(opt *model.Option) error {
	err := defaultDB.Save(opt).Error
	return err
}
