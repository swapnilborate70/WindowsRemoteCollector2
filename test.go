package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"strconv"
	"strings"

	"github.com/masterzen/winrm"
	"golang.org/x/net/context"
)

func main() {
	// Replace these with your own values
	endpoint := "192.168.2.102"
	username := "033-nitesh"
	password := "9835"
	port := 5985

	// Create a WinRM client
	endpointConfig := winrm.NewEndpoint(endpoint, port, false, false, nil, nil, nil, 0)
	client, err := winrm.NewClient(endpointConfig, username, password)
	if err != nil {
		fmt.Println("Error creating WinRM client:", err)

	}

	// Execute the PowerShell command remotely to get user percentage for each CPU core
	var stdout, stderr bytes.Buffer
	ctx := context.Background()

	CpuCommand1 := `C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe -Command "& {$userPercentages = (Get-Counter -Counter '\Processor(*)\% User Time').CounterSamples | Select-Object -ExpandProperty CookedValue; $userPercentages}"`
	client.RunWithContext(ctx, CpuCommand1, &stdout, &stderr)

	// fmt.Println("stdout:", strings.TrimSpace(stdout.String()))
	// fmt.Println("stderr:", strings.TrimSpace(stderr.String()))

	// Parse the output
	userPercentageStr := strings.TrimSpace(stdout.String())
	userPercentages := strings.Split(userPercentageStr, "\n")

	// Convert the string percentages to float64
	var userPercentageFloats []float64
	for _, val := range userPercentages {
		percent, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
		if err != nil {
			fmt.Println("Error parsing float:", err)
			continue
		}
		userPercentageFloats = append(userPercentageFloats, percent)
	}

	// // Use the parsed float values
	// fmt.Println("User percentages:", userPercentageFloats)

	stdout.Reset()
	stderr.Reset()

	CpuCommand2 := `powershell (Get-WmiObject -Class Win32_Processor).NumberOfLogicalProcessors`
	client.RunWithContext(ctx, CpuCommand2, &stdout, &stderr)

	// Parse the output and convert to int
	numLogicalProcessorsStr := strings.TrimSpace(stdout.String())
	numLogicalProcessors, err := strconv.Atoi(numLogicalProcessorsStr)
	if err != nil {
		fmt.Println("Error parsing output:", err)
		return
	}

	stdout.Reset()
	stderr.Reset()

	CpuCommand3 := `C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe -Command "& {$cpuCores = (Get-WmiObject Win32_Processor).NumberOfLogicalProcessors; $cores = Get-WmiObject -Class Win32_PerfFormattedData_PerfOS_Processor | Where-Object {$_.Name -notlike '*_Total'}; $cpuCorePercentages = @(); foreach ($core in $cores) { $corePercent = [math]::Round(($core.PercentProcessorTime / $cpuCores), 2); $cpuCorePercentages += $corePercent }; $cpuCorePercentages }"`

	client.RunWithContext(ctx, CpuCommand3, &stdout, &stderr)

	// fmt.Println("cmd :", strings.TrimSpace(stdout.String()))

	// Parse the output and convert to string
	cpucorepercentString := strings.TrimSpace(stdout.String())

	// Split the string into individual CPU core percentages
	cpucorepercentStrings := strings.Split(cpucorepercentString, "\n")

	// Create a slice to store CPU core percentages
	cpucorepercent := make([]string, len(cpucorepercentStrings))
	cpucorepercentFloats := make([]float64, len(cpucorepercentStrings))

	// Convert the string CPU core percentages to float
	for i, val := range cpucorepercentStrings {
		valNew, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
		if err != nil {
			fmt.Println("Error parsing float:", err)
			continue
		}
		cpucorepercentFloats[i] = valNew
	}

	// // Print the result to verify
	// fmt.Println(cpucorepercentFloats)

	stdout.Reset()
	stderr.Reset()

	CpuCommand4 := `C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe -Command "& {Get-Counter -Counter '\Processor(*)\% Interrupt Time' | Select-Object -ExpandProperty CounterSamples | Select-Object -ExpandProperty CookedValue}"`

	client.RunWithContext(ctx, CpuCommand4, &stdout, &stderr)

	// Parse the output and convert to string
	cpucoreinterruptpercentString := strings.TrimSpace(stdout.String())

	// Split the string into individual CPU core interrupt percentages
	cpucoreinterruptpercentStrings := strings.Split(cpucoreinterruptpercentString, "\n")

	// Create a slice to store CPU core interrupt percentages
	cpucoreinterruptpercent := make([]string, len(cpucoreinterruptpercentStrings))

	cpucoreinterruptpercentFloats := make([]float64, len(cpucoreinterruptpercentStrings))

	// Convert the string CPU core interrupt percentages to float
	for i, val := range cpucoreinterruptpercentStrings {
		valNew, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
		if err != nil {
			fmt.Println("Error parsing float:", err)
			continue
		}
		cpucoreinterruptpercentFloats[i] = valNew
	}

	// // Print the result to verify
	// fmt.Println(cpucoreinterruptpercentFloats)

	stdout.Reset()
	stderr.Reset()

	CpuCommand5 := `C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe -Command "& {Get-Counter -Counter '\Processor(*)\% Idle Time' | Select-Object -ExpandProperty CounterSamples | Select-Object -ExpandProperty CookedValue}"`

	client.RunWithContext(ctx, CpuCommand5, &stdout, &stderr)

	// Parse the output and convert to string
	cpucoreidlepercentString := strings.TrimSpace(stdout.String())

	// Split the string into individual CPU core idle percentages
	cpucoreidlepercentStrings := strings.Split(cpucoreidlepercentString, "\n")

	// Create a slice to store CPU core idle percentages
	cpucoreidlepercent := make([]string, len(cpucoreidlepercentStrings))

	cpucoreidlepercentFloats := make([]float64, len(cpucoreidlepercentStrings))

	// Convert the string CPU core idle percentages to float
	for i, val := range cpucoreidlepercentStrings {
		valNew, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
		if err != nil {
			fmt.Println("Error parsing float:", err)
			continue
		}
		cpucoreidlepercentFloats[i] = valNew
	}

	// // Print the result to verify
	// fmt.Println(cpucoreidlepercentFloats)

	stdout.Reset()
	stderr.Reset()

	VolumeCommand1 := `C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe -Command "& {Get-WmiObject -Class Win32_Volume | Select-Object -ExpandProperty Name}"`
	client.RunWithContext(ctx, VolumeCommand1, &stdout, &stderr)

	// Parse the command output to extract the names
	namesOfVolumes := strings.Split(strings.TrimSpace(stdout.String()), "\n")

	stdout.Reset()
	stderr.Reset()

	VolumeCommand2 := `C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe -Command "& {Get-WmiObject Win32_Volume | Select-Object FreeSpace}"`
	client.RunWithContext(ctx, VolumeCommand2, &stdout, &stderr)

	// Parse the command output to extract the free space values
	// Split the output by newline characters
	lines := strings.Split(strings.TrimSpace(stdout.String()), "\n")

	// Extract the free space values (as integers), skipping invalid lines
	var freeSpaceValues []int64
	for _, line := range lines {
		// Remove any leading/trailing whitespace
		line = strings.TrimSpace(line)

		// Attempt to parse the line as an integer
		value, err := strconv.ParseInt(line, 10, 64)
		if err == nil {
			freeSpaceValues = append(freeSpaceValues, value)
		}
	}

	// // Print the free space values
	// fmt.Println("Free space of volumes in bytes:", freeSpaceValues)

	stdout.Reset()
	stderr.Reset()

	VolumeCommand3 := `C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe -Command "& {Get-WmiObject Win32_Volume | Select-Object @{Name='UsedSpace';Expression={$_.Capacity - $_.FreeSpace}} | ForEach-Object { $_.UsedSpace }}"`
	client.RunWithContext(ctx, VolumeCommand3, &stdout, &stderr)

	// Parse the command output to extract the used space values
	// Split the output by newline characters
	parsedlines := strings.Split(strings.TrimSpace(stdout.String()), "\n")

	// Extract the used space values (as integers), skipping invalid lines
	var UsedSpaceValues []int64
	for _, line := range parsedlines {
		// Remove any leading/trailing whitespace
		line = strings.TrimSpace(line)

		// Attempt to parse the line as an integer
		value, err := strconv.ParseInt(line, 10, 64)
		if err == nil {
			UsedSpaceValues = append(UsedSpaceValues, value)
		}
	}

	// // Print the free space values
	// fmt.Println("used space of volumes in bytes:", UsedSpaceValues)

	stdout.Reset()
	stderr.Reset()

	VolumeCommand4 := `C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe -Command "& {Get-WmiObject Win32_Volume | Select-Object -ExpandProperty Capacity}"`
	client.RunWithContext(ctx, VolumeCommand4, &stdout, &stderr)

	// Parse the command output to extract the capacity values
	// Split the output by newline characters
	parsedlinesOfCapacity := strings.Split(strings.TrimSpace(stdout.String()), "\n")

	// Extract the used space values (as integers), skipping invalid lines
	var CapacityBytesValues []int64
	for _, line := range parsedlinesOfCapacity {
		// Remove any leading/trailing whitespace
		line = strings.TrimSpace(line)

		// Attempt to parse the line as an integer
		value, err := strconv.ParseInt(line, 10, 64)
		if err == nil {
			CapacityBytesValues = append(CapacityBytesValues, value)
		}
	}

	// // Print the free space values
	// fmt.Println("used space of volumes in bytes:", UsedSpaceValues)

	VolumeCommand5 := `C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe -Command "& {Get-WmiObject Win32_Volume | Select-Object @{Name='UsedPercent';Expression={($_.Capacity - $_.FreeSpace) / $_.Capacity * 100 -as [decimal]}} | ForEach-Object { '{0}' -f $_.UsedPercent }}"`

	client.RunWithContext(ctx, VolumeCommand5, &stdout, &stderr)

	// Split the output into separate lines
	outputLines := strings.Split(stdout.String(), "\n")

	// Extract values with decimals
	var volumeUsedPercent []float64
	for _, line := range outputLines {
		line = strings.TrimSpace(line) // Remove any leading or trailing spaces
		if strings.Contains(line, ".") {
			if value, err := strconv.ParseFloat(line, 64); err == nil {
				volumeUsedPercent = append(volumeUsedPercent, value)
			}
		}
	}

	// // Print extracted values
	// fmt.Println("Values of Used percent:")
	// for _, values := range volumeUsedPercent {
	// 	fmt.Println(values)
	// }

	stdout.Reset()
	stderr.Reset()

	VolumeCommand6 := `C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe -Command "& {Get-WmiObject Win32_Volume | Select-Object @{Name='FreePercent';Expression={($_.FreeSpace / $_.Capacity) * 100 -as [decimal]}} | ForEach-Object { '{0}' -f $_.FreePercent }}"`

	client.RunWithContext(ctx, VolumeCommand6, &stdout, &stderr)

	// Split the output into separate lines
	outputLinesOfFreePercent := strings.Split(stdout.String(), "\n")

	// Extract values with decimals
	var volumeFreePercent []float64
	for _, line := range outputLinesOfFreePercent {
		line = strings.TrimSpace(line) // Remove any leading or trailing spaces
		if strings.Contains(line, ".") {
			if value, err := strconv.ParseFloat(line, 64); err == nil {
				volumeFreePercent = append(volumeFreePercent, value)
			}
		}
	}

	// // Print extracted values
	// fmt.Println("Values of free percent:")
	// for _, valuesOfFreePercent := range volumeFreePercent {
	// 	fmt.Println(valuesOfFreePercent)
	// }

	stdout.Reset()
	stderr.Reset()

	VolumeCommand10 := `C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe -Command "& {(Get-WmiObject Win32_Volume).Count}"`

	client.RunWithContext(ctx, VolumeCommand10, &stdout, &stderr)

	// Parse the output and trim it
	countStringOfNoOfVolume := strings.TrimSpace(stdout.String())

	// // Print the count
	// fmt.Println("Count of volumes:", countStringOfNoOfVolume)

	// Convert the count string to an integer
	countOfNoOfVolume, err := strconv.Atoi(countStringOfNoOfVolume)

	InterfaceCommand10 := `C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe -Command "& {(Get-NetAdapter).Count}"`

	client.RunWithContext(ctx, InterfaceCommand10, &stdout, &stderr)

	noOfInterfaces, err := bufferToInt(stdout)

	stdout.Reset()
	stderr.Reset()

	InterfaceCommand1 := `C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe -Command "& {Get-Counter -Counter "\Network Interface(*)\Bytes Received/sec" | Select-Object -ExpandProperty CounterSamples | Select-Object -ExpandProperty CookedValue}"`

	client.RunWithContext(ctx, InterfaceCommand1, &stdout, &stderr)

	interfacesBytesPerSec := getInputBytesPerSec(stdout)

	stdout.Reset()
	stderr.Reset()

	InterfaceCommand2 := `C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe -Command "& {Get-NetAdapter | Select-Object -ExpandProperty Name}"`

	client.RunWithContext(ctx, InterfaceCommand2, &stdout, &stderr)

	namesOfInterfaces := getNamesOfInterfaces(stdout)

	// Create a slice to store CPU info
	cpuInfoList := make([]CPUInfo, numLogicalProcessors)
	diskVolumeInfoList := make([]Volume, countOfNoOfVolume)
	interfacesList := make([]Interface, noOfInterfaces)

	fmt.Println()
	fmt.Println(len(interfacesList))
	fmt.Println(interfacesList)

	for i := 0; i < noOfInterfaces; i++ {
		interfaces := Interface{
			SystemNetworkInterface: namesOfInterfaces[i],
		}

		if i < len(interfacesBytesPerSec) {
			interfaces.SystemNetworkInterfaceInBytesPerSec = interfacesBytesPerSec[i]
		}

		interfacesList[i] = interfaces
	}

	// Convert CPU info to JSON
	interfacesJSON, err := json.Marshal(interfacesList)
	if err != nil {
		fmt.Println("Error encoding CPU info to JSON:", err)
		return
	}

	fmt.Println(string(interfacesJSON))

	for i := 0; i < countOfNoOfVolume; i++ {
		disks := Volume{
			SystemDiskVolume: namesOfVolumes[i],
		}

		if i < len(freeSpaceValues) {
			disks.SystemDiskVolumeFreeBytes = freeSpaceValues[i]
		}

		if i < len(UsedSpaceValues) {
			disks.SystemDiskVolumeUsedBytes = UsedSpaceValues[i]
		}

		if i < len(CapacityBytesValues) {
			disks.SystemDiskVolumeCapacityBytes = CapacityBytesValues[i]
		}

		if i < len(volumeUsedPercent) {
			disks.SystemDiskVolumeUsedPercentage = volumeUsedPercent[i]
		}

		if i < len(volumeFreePercent) {
			disks.SystemDiskVolumeFreePercent = volumeFreePercent[i]
		}

		diskVolumeInfoList[i] = disks
	}

	// Convert CPU info to JSON
	diskVolumeJSON, err := json.Marshal(diskVolumeInfoList)
	if err != nil {
		fmt.Println("Error encoding CPU info to JSON:", err)
		return
	}

	fmt.Println(string(diskVolumeJSON))

	// Iterate through each CPU core
	for i := 0; i < numLogicalProcessors; i++ {
		cpu := CPUInfo{
			SystemCPUCore: strconv.Itoa(i),
		}

		// Extract user percentage for the current CPU core
		if i < len(userPercentages) {
			cpu.SystemCPUCoreUserPercent = userPercentageFloats[i]
		}

		// Extract cpu core percentage for the current CPU core
		if i < len(cpucorepercent) {
			cpu.SystemCPUCorePercent = cpucorepercentFloats[i]
		}

		// Extract cpu core interrupt percentage for the current CPU core
		if i < len(cpucoreinterruptpercent) {
			cpu.SystemCPUCoreInterruptPercent = cpucoreinterruptpercentFloats[i]
		}

		// Extract cpu core idle percentage for the current CPU core
		if i < len(cpucoreidlepercent) {
			cpu.SystemCPUCoreIdlePercent = cpucoreidlepercentFloats[i]
		}

		cpuInfoList[i] = cpu
	}

	// Convert CPU info to JSON
	cpuJSON, err := json.Marshal(cpuInfoList)
	if err != nil {
		fmt.Println("Error encoding CPU info to JSON:", err)
		return
	}

	fmt.Println(string(cpuJSON))
}
