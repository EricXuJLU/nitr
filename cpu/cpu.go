package cpu

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber"
	"github.com/jaypipes/ghw"
	"github.com/shirou/gopsutil/cpu"
)

type CPU struct {
	Brand     string    `json:"brand"`
	Cores     uint32    `json:"cores"`
	Threads   uint32    `json:"threads"`
	Frecuency float64   `json:"frecuency"`
	Usage     float64   `json:"usage"`
	UsageEach []float64 `json:"usageEach"`
}

func CpuUsage() float64 {
	duration := 500 * time.Millisecond
	cpuUsage, err := cpu.Percent(duration, false)
	if err != nil {
		panic(err)
	}
	return cpuUsage[0]
}

func cpuUsageEach() []float64 {
	duration := 500 * time.Millisecond
	cpuUsage, err := cpu.Percent(duration, true)
	if err != nil {
		panic(err)
	}
	return cpuUsage
}

func cores() uint32 {
	cpu, err := ghw.CPU()
	if err != nil {
		fmt.Printf("Error getting CPU info: %v", err)
	}

	return cpu.TotalCores
}

func threads() uint32 {
	cpu, err := ghw.CPU()
	if err != nil {
		fmt.Printf("Error getting CPU info: %v", err)
	}

	return cpu.TotalThreads
}

func frecuency() float64 {
	cpu, err := cpu.Info()
	if err != nil {
		fmt.Printf("Error getting CPU info: %v", err)
	}
	return cpu[0].Mhz
}

func model() string {
	cpu, err := ghw.CPU()
	if err != nil {
		fmt.Printf("Error getting CPU info: %v", err)
	}

	return cpu.Processors[0].Model
}

func Check() CPU {
	return CPU{
		Brand:     model(),
		Cores:     cores(),
		Threads:   threads(),
		Frecuency: frecuency(),
		Usage:     CpuUsage(),
		UsageEach: cpuUsageEach(),
	}
}

func Data(c *fiber.Ctx) {
	c.JSON(Check())
}
