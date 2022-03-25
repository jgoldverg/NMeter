package metrics

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
	net "github.com/shirou/gopsutil/v3/net"
)

type KernelMetrics struct {
	CpuFrequency   float32
	EnergyConsumed float32
	CpuArch        string
}

func GetCpuInfo(KernelMetrics *KernelMetrics) {

}

func getVmemory() {
	v, _ := mem.VirtualMemory()
	fmt.Printf("%+v\n", v)
}

func getProtoCounters() {
	var empty = []string{}
	counters, _ := net.ProtoCounters(empty)
	for _, v := range counters {
		fmt.Printf("Interfaces:\n %+v\n", v)
	}
}
