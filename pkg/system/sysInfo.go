package system

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// SystemInfo struct contains all
// information about the system
type SysInfo struct {
	Hostname  string `bson:hostname`
	Platform  string `bson:platform`
	CPU       string `bson:cpu`
	RAM_All   uint64 `bson:ram_all`
	RAM_USED  uint64 `json:"ram_used"`
	SWAP_All  uint64 `json:"swap_all"`
	SWAP_USED uint64 `json:"swap_used"`
	Disk      uint64 `bson:disk`
}

// This function returns all information
// about the system and parses them into
// the SysInfo struct
func GetSystemInformation() *SysInfo {
	hostStat, err := host.Info()
	if err != nil {
		panic(err.Error())
	}
	cpuStat, err := cpu.Info()
	if err != nil {
		panic(err.Error())
	}
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		panic(err.Error())
	}
	swap, err := mem.SwapMemory()
	if err != nil {
		panic(err.Error())
	}
	diskStat, err := disk.Usage("/") // If you're in Unix change this "\\" for "/"
	if err != nil {
		panic(err.Error())
	}
	info := new(SysInfo)
	info.Hostname = hostStat.Hostname
	info.Platform = hostStat.Platform
	info.CPU = cpuStat[0].ModelName
	info.RAM_All = vmStat.Total / 1024 / 1024
	info.RAM_USED = vmStat.Used / 1024 / 1024
	info.SWAP_All = swap.Total / 1024 / 1024
	info.SWAP_USED = swap.Used / 1024 / 1024
	info.Disk = diskStat.Total / 1024 / 1024 / 1024

	return info
}
