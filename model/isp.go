package model

import (
	"github.com/bitcav/nitr-core/isp"
	"gorm.io/gorm"
)

type Isp struct {
	gorm.Model
	Isp string `xml:"isp,attr" json:"isp"`
	IP  string `xml:"ip,attr" json:"ip"`
	Lat string `xml:"lat,attr" json:"lat"`
	Lon string `xml:"lon,attr" json:"lon"`
}

func ParseIspInfo(src isp.Setting) (dst Isp) {
	JsonConvert(src, &dst)
	return
}