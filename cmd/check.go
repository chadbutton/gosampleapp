package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Checks the health of the machine and prints out a simple report.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		Check()
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func printReport() {
	var appName = viper.Get("app.name").(string)
	var colorGreen = color.New(color.FgGreen)
	var colorRed = color.New(color.FgRed)
	var checkmark = "[pass]"
	var xmark = "[fail]"
	var meetsOsRequirements = true
	var useConsoleGraphics = viper.GetBool("console.showGraphics")

	if useConsoleGraphics {
		checkmark = "✔"
		xmark = "✘"
	}

	fmt.Println("")
	fmt.Println("OS Report Summary for: " + appName)
	fmt.Println("=================")

	vm, _ := mem.VirtualMemory()
	fmt.Printf("Memory Total: %v, Available: %v, Used: %v, Used Percent: %.2f%%\n", vm.Total, vm.Available, vm.Used, vm.UsedPercent)

	meetsOsRequirements = meetsOsRequirements && vm.Total > 10000000000

	cpus, _ := cpu.Info()
	for _, c := range cpus {
		fmt.Printf("Cpu Model: %v, Cores: %v\n", c.ModelName, c.Cores)
		meetsOsRequirements = meetsOsRequirements && c.Cores > 2
	}

	partitions, _ := disk.Partitions(false)
	for _, p := range partitions {
		if p.Device != "" {
			fmt.Printf("Disk: %v", p.Device)
			d, _ := disk.Usage(p.Device)
			fmt.Printf("  Total %v, Free: %v, Used: %v, UsedPercent: %.2f%%\n", d.Total, d.Free, d.Used, d.UsedPercent)
			meetsOsRequirements = meetsOsRequirements && d.UsedPercent > 50
		}
	}

	if meetsOsRequirements {
		colorGreen.Println(checkmark, " meets all OS requirements!")
	} else {
		colorRed.Println(xmark, " does not meet all OS memory and cpu requirements, see support manager for assistance")
	}
}

func Check() {
	printReport()
}
