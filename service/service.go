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

func UpdateCPUStatus() {
	info := cpu.Info()
	gCpu := model.ParseCPUInfo(info)
	_, err := dao.CreateCPUInfo(&gCpu)
	if err != nil {
		log.Println("CPU Error: " + err.Error())
	}
}

func UpdateDiskStatus() {
	infos := disk.Info()
	for _, info := range infos {
		gDisk := model.ParseDiskInfo(info)
		_, err := dao.CreateDiskInfo(&gDisk)
		if err != nil {
			log.Println("Disk Error: " + err.Error())
		}
	}
}

func UpdateGPUStatus() {
	infos := gpu.Info()
	for _, info := range infos {
		gGPU := model.ParseGPUInfo(info)
		_, err := dao.CreateGPUInfo(&gGPU)
		if err != nil {
			log.Println("GPU Error: " + err.Error())
		}
	}
}

func UpdateHostStatus() {
	info := host.Info()
	gHost := model.ParseHostInfo(info)
	_, err := dao.CreateHostInfo(&gHost)
	if err != nil {
		log.Println("Host Error: " + err.Error())
	}
}

func UpdateIspStatus() {
	info := isp.Info()
	gIsp := model.ParseIspInfo(info)
	_, err := dao.CreateIspInfo(&gIsp)
	if err != nil {
		log.Println("Isp Error: " + err.Error())
	}
}

func UpdateMemoryStatus() {
	infos, err := memdev.Info()
	if err != nil {
		log.Println("Memory Error: " + err.Error())
		return
	}
	for _, info := range infos {
		gMemory := model.ParseMemoryInfo(info)
		_, err := dao.CreateMemoryInfo(&gMemory)
		if err != nil {
			log.Println("Memory Error: " + err.Error())
		}
	}
}

func UpdateNetworkStatus() {
	infos := network.Info()
	nets := model.ParseNetworkInfo(infos)
	for _, info := range nets {
		_, err := dao.CreateNetworkInfo(&info)
		if err != nil {
			log.Println("Network Error: " + err.Error())
		}
	}
}

func UpdateProcessStatus() {
	infos := process.Info()
	for _, info := range infos {
		gProcess := model.ParseProcessInfo(info)
		_, err := dao.CreateProcessInfo(&gProcess)
		if err != nil {
			log.Println("Process Error: " + err.Error())
		}
	}
}

func UpdateRAMStatus() {
	info := ram.Info()
	gRAM := model.ParseRAMInfo(info)
	_, err := dao.CreateRAMInfo(&gRAM)
	if err != nil {
		log.Println("RAM Error: " + err.Error())
	}
}
