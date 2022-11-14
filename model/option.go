package model

type Option struct {
	ID       uint `gorm:"primarykey"`
	Interval int
}
