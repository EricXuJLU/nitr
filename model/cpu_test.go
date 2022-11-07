package model

import (
	"fmt"
	"github.com/bitcav/nitr-core/cpu"
	"testing"
)

func TestJsonConvert(t *testing.T) {
	src := cpu.CPU{
		Vendor:     "vendorxx",
		Model:      "Modelxx",
		Cores:      8,
		Threads:    16,
		ClockSpeed: 20000000,
		Usage:      0.3,
		UsageEach:  []float64{0.1, 0.5},
	}
	var dst CPU
	JsonConvert(src, &dst)
	if dst.Cores != src.Cores {
		t.Error(dst)
	}
	fmt.Printf("%#v", dst)
}
