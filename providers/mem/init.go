package mem

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
    "github.com/pivotal-golang/bytefmt"
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
	Usage = fmt.Sprintf("%.2f%%", memStats.UsedPercent)

	SwapTotal = bytefmt.ByteSize(swapStats.Total)
	SwapUsed = bytefmt.ByteSize(swapStats.Used)
	SwapFree = bytefmt.ByteSize(swapStats.Free)
	SwapUsage = fmt.Sprintf("%.2f%%", swapStats.UsedPercent)
}
