package providers

import (
	"fmt"

	"github.com/gosuri/uitable"
	"github.com/shirou/gopsutil/mem"
)

const (
	_         = iota
	KB uint64 = 1 << (10 * iota)
	MB
	GB
	TB
)

func MemStats() string {
	memStats, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}
	memTable := uitable.New()
	memTable.AddRow("MEMORY")
	memTable.AddRow("Total(MB)", "Used(MB)", "Free(MB)", "Usage")
	memTable.AddRow(memStats.Total/MB, memStats.Used/MB, memStats.Free/MB, fmt.Sprintf("%.2f%%", memStats.UsedPercent))

	return memTable.String()
}
