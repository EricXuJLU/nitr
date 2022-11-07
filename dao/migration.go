package dao

import (
	"github.com/bitcav/nitr/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Param 0-UserName, 1-Password, 2-Address, 3-DBName
var Param []string
var defaultDB *gorm.DB

func init() {
	if len(Param[0]) == 0 {
		Param[0] = "root"
	}
	if len(Param[1]) == 0 {
		Param[1] = "password123"
	}
	if Param[1] == "NoPassword" {
		Param[1] = ""
	}
	if len(Param[2]) == 0 {
		Param[2] = "127.0.0.1:3306"
	}
	if len(Param[3]) == 0 {
		Param[3] = "nitr"
	}
	//"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := Param[0] + ":" + Param[1] + "@tcp(" + Param[2] + ")/" + Param[3] + "?charset=utf8mb4&parseTime=True&loc=Local"
	dbTmp, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Open Database Error: " + err.Error())
	}
	defaultDB = dbTmp
	err = migrate(defaultDB)
	if err != nil {
		panic("Migrate Tables Error: " + err.Error())
	}
}

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.CPU{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&model.Disk{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&model.Memory{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&model.RAM{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&model.Process{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&model.GPU{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&model.Isp{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&model.Host{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&model.Network{})
	if err != nil {
		return err
	}
	return nil
}
