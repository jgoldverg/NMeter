package metrics

import (
	"fmt"

	"runtime"
	"time"

	"github.com/go-ping/ping"
	net "github.com/shirou/gopsutil/v3/net"
)

type Metrics struct {
	NetworkMetrics NetworkMetrics
	KernelMetrics  KernelMetrics
}

type NetworkMetrics struct {
	Interfaces            []string
	Rtt                   float32
	Bandwidth             float32
	BandwidthDelayProduct float32
	PacketLossRate        float32
	LinkCapacity          float32
	BytesSent             float64
	BytesRecv             float64
	PacketsSent           int64
	PacketsRecv           int64
	DropIn                int32
	DropOut               int32
	NicSpeed              int32
	NicMtu                int32
	StartTime             JSONTime
	EndTime               JSONTime
	Count                 int32
	Latency               []float32
	PingStats             []ping.Statistics
}

type JSONTime struct {
	time.Time
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", t.Format("Mon Jan _2"))
	return []byte(stamp), nil
}

func RunPing(sendCount int, networkMetrics *NetworkMetrics, duration time.Duration) {
	// var pingStats = make([]ping.Statistics, sendCount, sendCount)
	pinger, err := ping.NewPinger("www.google.com")
	if err != nil {
		panic(err)
	}
	pinger.Count = sendCount
	if runtime.GOOS == "darwin" {
		pinger.SetPrivileged(true)
	}
	pinger.Interval = duration

	// Listen for Ctrl-C.
	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}
	pinger.OnDuplicateRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
	}
	pinger.OnFinish = func(stats *ping.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		panic(err)
	}

	if networkMetrics.PingStats == nil {
		networkMetrics.PingStats = make([]ping.Statistics, 0)
		networkMetrics.PingStats = append(networkMetrics.PingStats, *pinger.Statistics())
	} else {
		networkMetrics.PingStats = append(networkMetrics.PingStats, *pinger.Statistics())
	}
	fmt.Printf("The Stat struct is = %+v\n", networkMetrics.PingStats[len(networkMetrics.PingStats)-1])
}

func GetAllInterfaceStats(networkMetrics *NetworkMetrics) {
	info, _ := net.Interfaces()
	fmt.Printf("Interfaces:\n %+v\n", info)
}

func GetInterfaceStats(inter string, networkMetrics *NetworkMetrics) {
}

func GetNetInfo() {
	info, _ := net.IOCounters(true)
	for index, v := range info {
		fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	}
}

func getIOCounters() {
	info, _ := net.IOCounters(true)
	fmt.Printf("NIC:\n %+v\n", info)
}
