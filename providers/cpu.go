package providers

import (
	"github.com/gosuri/uitable"
	"github.com/shirou/gopsutil/cpu"
)

func CpuStats() string {
	cpuStats, err := cpu.CPUInfo()
	if err != nil {
		panic(err)
	}
	cpuTable := uitable.New()
	cpuTable.AddRow("#", "CPU MODEL", "CORES")
	for i, stat := range cpuStats {
		cpuTable.AddRow(i, stat.ModelName, stat.Cores)
	}

	return cpuTable.String()
}
