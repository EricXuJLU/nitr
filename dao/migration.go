package dao

import (
	"github.com/bitcav/nitr/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// Param 0-UserName, 1-Password, 2-Address, 3-DBName
var Param = []string{"root", "PassWord_123", "127.0.0.1:3306", "nitr"}
var defaultDB *gorm.DB

const PianYi = 2

func init() {
	for i := 0; i < len(Param); i++ {
		if len(os.Args) > i+PianYi {
			Param[i] = os.Args[i+PianYi]
		} else {
			break
		}
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
