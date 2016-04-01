package mem

import (
	"github.com/pivotal-golang/bytefmt"
	"github.com/shirou/gopsutil/mem"

	"iSystem/formats"
)

var (
	Total     string
	Used      string
	Free      string
	Usage     string
	SwapTotal string
	SwapUsed  string
	SwapFree  string
	SwapUsage string
)

func init() {
	memStats, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}
	swapStats, err := mem.SwapMemory()
	if err != nil {
		panic(err)
	}

	Total = bytefmt.ByteSize(memStats.Total)
	Used = bytefmt.ByteSize(memStats.Used)
	Free = bytefmt.ByteSize(memStats.Free)
	Usage = memFormat(memStats.UsedPercent)

	SwapTotal = bytefmt.ByteSize(swapStats.Total)
	SwapUsed = bytefmt.ByteSize(swapStats.Used)
	SwapFree = bytefmt.ByteSize(swapStats.Free)
	SwapUsage = memFormat(swapStats.UsedPercent)
}

func memFormat(usage float64) string {
	format := "%.2f%%"
	if usage > 70 {
		return formats.Dangerf(format, usage)
	} else if usage > 50 {
		return formats.Warningf(format, usage)
	} else {
		return formats.Successf(format, usage)
	}
}
