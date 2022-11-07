package dao

import (
	"github.com/bitcav/nitr/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// Param 0-UserName, 1-Password, 2-Address, 3-DBName
var Param []string
var defaultDB *gorm.DB

func init() {
	Param = make([]string, 4)
	if len(os.Args) > 1 {
		Param[0] = os.Args[1]
	} else {
		Param[0] = "root"
	}
	if len(os.Args) > 2 {
		Param[1] = os.Args[2]
	} else {
		Param[1] = "PassWord_123"
	}
	if Param[1] == "NoPassword" {
		Param[1] = ""
	}
	if len(os.Args) > 3 {
		Param[2] = os.Args[3]
	} else {
		Param[2] = "127.0.0.1:3306"
	}
	if len(os.Args) > 4 {
		Param[3] = os.Args[4]
	} else {
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
