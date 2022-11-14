package model

import (
	"github.com/bitcav/nitr-core/process"
	"github.com/bitcav/nitr/lib"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type Process struct {
	gorm.Model
	HostName string `json:"hostName"`
	Pid      int    `json:"pid"`
	Name     string `json:"name"`
}

func ParseProcessInfo(src process.Process) (dst Process) {
	JsonConvert(src, &dst)
	return
}

func ParseProcessesFromString(desc, hostName string) []Process {
	var ans []Process
	var processes = lib.StringsDeleteNil(strings.Split(desc, "\n"))
	for i := 1; i < len(processes); i++ {
		units := lib.StringsDeleteNil(strings.Split(processes[i], " "))
		pid, _ := strconv.Atoi(units[1])
		tmp := Process{
			HostName: hostName,
			Pid:      pid,
			Name:     strings.Join(units[10:], " "),
		}
		//fmt.Println(tmp.Pid, tmp.HostName, tmp.Name)
		ans = append(ans, tmp)
	}
	return ans
}
