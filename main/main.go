package main

import (
	"fmt"
	"nmeter/config"
	"nmeter/metrics"
	"os"
	"path/filepath"
	"time"

	"github.com/docopt/docopt-go"
)

func main() {
	usage := `PMeter a tool to measure the TCP/UDP and other kernel, network conditions that the running host experiences

Usage:
  nmeter [options] measure <INTERFACE>

Commands:
measure     The command to start measuring the computers network activity on the specified network devices. This command accepts a list of interfaces that you wish to monitor

Options:
  -h --help     				Show this screen.
  --version     				Show version.
  -C --config_path=CONFIGPATH	Set the path to the config file [default: .nmeter/]
  --gometer_path=GOMETERPATH	Set the path to store the measured data uses the current running users home directory to create this folder [default: .nmeter/]
  --file_name=FILE_NAME  		Set the file name used to store measurements [default: network_results.txt]
  -N --measure_network     		Set if we monitor only the network interface [default: False]
  -K --measure_kernel      		Set if we monitor only the kernel [default: False]
  -U --measure_udp         		Set UDP monitoring only [default: False]
  -T --measure_tcp         		Set TCP monitoring only [default: False]
  -S --enable_std_out      		Disable printing the results to standard output [default: False]
  --interval=INTERVAL			Set the delay between measurements must follow format Ex: 100ms”, “2.3h” or “4h35m”. Valid units of time are Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h” [default: 5s]
  --measure=MEASUREMENTS   		The number of times to run the measurement [default: 1]
`
	arguments, _ := docopt.ParseDoc(usage)
	fmt.Println(arguments)
	if _, ok := arguments["measure"]; ok {
		inFace, _ := arguments.String("<INTERFACE>")
		measurementFile, _ := arguments.String("--file_name")
		measureNetwork, _ := arguments.Bool("--measure_network")
		measureKernel, _ := arguments.Bool("--measure_kernel")
		measureUdp, _ := arguments.Bool("--measure_udp")
		measureTcp, _ := arguments.Bool("--measure_tcp")
		printStdOut, _ := arguments.Bool("--enable_std_out")
		dataFolderPath, _ := arguments.String("--gometer_path")
		interval, _ := arguments.String("--interval")
		configPath, _ := arguments.String("--config_path")
		numberOfMeasurements, _ := arguments.Int("--measure")
		timeDuration, _ := time.ParseDuration(interval)
		userHome, _ := os.UserHomeDir()
		config, _ := config.ReadConfig(filepath.Join(userHome, configPath))
		if measurementFile != config.File.DataFileName {
			config.File.DataFileName = measurementFile
		}
		if len(dataFolderPath) > 0 &&  {
			config.File.DataFilePath = dataFolderPath
		}
		networkMetrics := metrics.NetworkMetrics{}
		if measureNetwork {
			if len(inFace) > 0 {
				metrics.GetInterfaceStats(inFace, &networkMetrics)
			} else {
				metrics.GetAllInterfaceStats(&networkMetrics)
			}
			metrics.RunPing(numberOfMeasurements, &networkMetrics, timeDuration)
		}
		kernelMetrics := metrics.KernelMetrics{}
		if measureKernel {
			// metrics.GetCpuInfo()
			// getCPUInfo(kernelMetrics)
			// getIOCounters()
		}

		if measureTcp {

		}

		if measureUdp {

		}
		if printStdOut {
			fmt.Printf("%+v\n", networkMetrics)
		}

	}

}
