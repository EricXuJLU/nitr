package model

import (
	"github.com/bitcav/nitr-core/network"
	"gorm.io/gorm"
)

type Network struct {
	gorm.Model
	Name   string `json:"name"`
	IP0    string `json:"IP0"`
	IP1    string `json:"IP1"`
	IP2    string `json:"IP2"`
	MAC    string `json:"mac"`
	Active bool   `json:"active"`
}

func ParseNetworkInfo(devices network.NetworkDevices) []Network {
	var ans []Network
	for _, net := range devices {
		var tmp = Network{
			Name:   net.Name,
			MAC:    net.MAC,
			Active: net.Active,
		}
		for _, ip := range net.Addresses {
			if len(tmp.IP0) == 0 {
				tmp.IP0 = ip.IP
			} else if len(tmp.IP1) == 0 {
				tmp.IP1 = ip.IP
			} else if len(tmp.IP2) == 0 {
				tmp.IP2 = ip.IP
			} else {
				break
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}
