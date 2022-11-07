package service

import (
	"github.com/bitcav/go-memdev"
	"github.com/bitcav/nitr-core/cpu"
	"github.com/bitcav/nitr-core/disk"
	"github.com/bitcav/nitr-core/gpu"
	"github.com/bitcav/nitr-core/host"
	"github.com/bitcav/nitr-core/isp"
	"github.com/bitcav/nitr-core/network"
	"github.com/bitcav/nitr-core/process"
	"github.com/bitcav/nitr-core/ram"
	"github.com/bitcav/nitr/dao"
	"github.com/bitcav/nitr/model"
	"log"
)

func UpdateCPUStatus(hostName string) {
	info := cpu.Info()
	gCpu := model.ParseCPUInfo(info)
	gCpu.HostName = hostName
	_, err := dao.CreateCPUInfo(&gCpu)
	if err != nil {
		log.Println("CPU Error: " + err.Error())
	}
}

func UpdateDiskStatus(hostName string) {
	infos := disk.Info()
	for _, info := range infos {
		gDisk := model.ParseDiskInfo(info)
		gDisk.HostName = hostName
		_, err := dao.CreateDiskInfo(&gDisk)
		if err != nil {
			log.Println("Disk Error: " + err.Error())
		}
	}
}

func UpdateGPUStatus(hostName string) {
	infos := gpu.Info()
	for _, info := range infos {
		gGPU := model.ParseGPUInfo(info)
		gGPU.HostName = hostName
		_, err := dao.CreateGPUInfo(&gGPU)
		if err != nil {
			log.Println("GPU Error: " + err.Error())
		}
	}
}

func GetHostName() string {
	info := host.Info()
	return info.Name
}

func UpdateHostStatus(name string) {
	err := dao.DeleteHostByName(name)
	if err != nil {
		log.Fatal("Delete Host Error: ", err)
	}
	info := host.Info()
	gHost := model.ParseHostInfo(info)
	_, err = dao.CreateHostInfo(&gHost)
	if err != nil {
		log.Println("Host Error: " + err.Error())
	}
}

func UpdateIspStatus(hostName string) {
	info := isp.Info()
	gIsp := model.ParseIspInfo(info)
	gIsp.HostName = hostName
	_, err := dao.CreateIspInfo(&gIsp)
	if err != nil {
		log.Println("Isp Error: " + err.Error())
	}
}

func UpdateMemoryStatus(hostName string) {
	infos, err := memdev.Info()
	if err != nil {
		log.Println("Memory Error: " + err.Error())
		return
	}
	for _, info := range infos {
		gMemory := model.ParseMemoryInfo(info)
		gMemory.HostName = hostName
		_, err := dao.CreateMemoryInfo(&gMemory)
		if err != nil {
			log.Println("Memory Error: " + err.Error())
		}
	}
}

func UpdateNetworkStatus(hostName string) {
	infos := network.Info()
	nets := model.ParseNetworkInfo(infos)
	for _, info := range nets {
		info.HostName = hostName
		_, err := dao.CreateNetworkInfo(&info)
		if err != nil {
			log.Println("Network Error: " + err.Error())
		}
	}
}

func UpdateProcessStatus(hostName string) {
	infos := process.Info()
	for _, info := range infos {
		gProcess := model.ParseProcessInfo(info)
		gProcess.HostName = hostName
		_, err := dao.CreateProcessInfo(&gProcess)
		if err != nil {
			log.Println("Process Error: " + err.Error())
		}
	}
}

func UpdateRAMStatus(hostName string) {
	info := ram.Info()
	gRAM := model.ParseRAMInfo(info)
	gRAM.HostName = hostName
	_, err := dao.CreateRAMInfo(&gRAM)
	if err != nil {
		log.Println("RAM Error: " + err.Error())
	}
}
