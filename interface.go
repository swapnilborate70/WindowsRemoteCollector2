package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Define CPUInfo struct
type Interface struct {
	SystemNetworkInterface                  string `json:"system.network.interface"`
	SystemNetworkInterfaceInBytesPerSec     int64  `json:"system.network.interface.in.bytes.per.sec"`
	SystemNetworkInterfaceOutputQueueLength int64  `json:"system.network.interface.output.queue.length"`
	SystemNetworkInterfaceInPacketsPerSec   int64  `json:"system.network.interface.in.packets.per.sec"`
	SystemNetworkInterfaceOutPacketsPerSec  int64  `json:"system.network.interface.out.packets.per.sec"`
	SystemNetworkInterfaceOutBytesPerSec    int64  `json:"system.network.interface.out.bytes.per.sec"`
	SystemNetworkInterfaceBytesPerSec       int64  `json:"system.network.interface.bytes.per.sec"`
}

func bufferToInt(buf bytes.Buffer) (int, error) {
	// Extract the integer from the output
	output := buf.String()
	re := regexp.MustCompile(`\d+`)
	match := re.FindString(output)
	if match == "" {
		fmt.Println("Error: No integer found in output")

	}

	// Convert string to integer
	integerValue, err := strconv.Atoi(match)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)

	}
	return integerValue, err
}

func getNamesOfInterfaces(bytes bytes.Buffer) []string {

	// Process the output to extract adapter names
	output := strings.TrimSpace(bytes.String())
	names := strings.Split(output, "\r\n")

	var namesOfInterfaces []string
	for _, name := range names {
		namesOfInterfaces = append(namesOfInterfaces, name)
	}
	return namesOfInterfaces
}

func getInputBytesPerSec(bytes bytes.Buffer) []int64 {

	// Split the output into separate lines
	outputLinesOfBytesPerSec := strings.Split(bytes.String(), "\n")

	var bytesPerSecValues []int64

	for _, line := range outputLinesOfBytesPerSec {
		// Remove any leading/trailing whitespace
		line = strings.TrimSpace(line)

		// Attempt to parse the line as an integer
		value, err := strconv.ParseInt(line, 10, 64)
		if err == nil {
			bytesPerSecValues = append(bytesPerSecValues, value)
		}
	}
	return bytesPerSecValues
}
