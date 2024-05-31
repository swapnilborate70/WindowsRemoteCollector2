package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	// // "bytes"
	// "strings"
)

type Interface struct {
	SystemNetworkInterface                  string  `json:"system.network.interface"`
	SystemNetworkInterfaceInBytesPerSec     float64 `json:"system.network.interface.in.bytes.per.sec"`
	SystemNetworkInterfaceOutputQueueLength float64 `json:"system.network.interface.output.queue.length"`
	SystemNetworkInterfaceInPacketsPerSec   float64 `json:"system.network.interface.in.packets.per.sec"`
	SystemNetworkInterfaceOutPacketsPerSec  float64 `json:"system.network.interface.out.packets.per.sec"`
	SystemNetworkInterfaceOutBytesPerSec    float64 `json:"system.network.interface.out.bytes.per.sec"`
	SystemNetworkInterfaceBytesPerSec       float64 `json:"system.network.interface.bytes.per.sec"`
}

func parseInterfaceOutput(output string) ([]string, []float64) {
	lines := strings.Split(output, "\n")
	names := make([]string, 0)
	values := make([]float64, 0)

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		names = append(names, parts[0])
		valueStr := strings.TrimSpace(parts[1])
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			fmt.Println("Error parsing value:", err)
			continue
		}
		values = append(values, (value))
	}

	return names, values
}

func getNamesOfInterfaces(bytes bytes.Buffer) []string {
	output := strings.TrimSpace(bytes.String())
	names := strings.Split(output, "\n")

	for i, name := range names {
		names[i] = strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) {
				return unicode.ToLower(r)
			}
			if r == '(' || r == '{' {
				return '[' // Replace brackets with square brackets
			}
			if r == ')' || r == '}' {
				return ']' // Replace brackets with square brackets
			}
			return r
		}, name)
	}

	return names
}
