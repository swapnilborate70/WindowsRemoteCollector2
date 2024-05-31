package main

import (
	"fmt"
	// // "bytes"
	// "strings"
	"bufio"
	"bytes"
)

type Memory struct {
	SystemMemoryName       string  `json:"system.memory.name"`
	SystemMemoryUsedSpace  float64 `json:"system.memory.used.space"`
	SystemMemoryFreeSpace  float64 `json:"system.memory.free.space"`
	SystemMemoryTotalSpace float64 `json:"system.memory.total.space"`
}

func parseMemoryOutput(stdout bytes.Buffer) ([]string, []float64, []float64, []float64) {
	var deviceLocators []string
	var usedSpaceBytes []float64
	var freeSpaceBytes []float64
	var capacityBytes []float64

	scanner := bufio.NewScanner(&stdout)
	var line1, line2, line3 string
	var err error

	for scanner.Scan() {
		line1 = scanner.Text()
		if line1 == "" {
			fmt.Println("Error: missing device locator")
			continue
		}

		if !scanner.Scan() {
			fmt.Println("Error: missing used space line")
			break
		}
		line2 = scanner.Text()
		if line2 == "" {
			fmt.Println("Error: empty used space line")
			continue
		}

		if !scanner.Scan() {
			fmt.Println("Error: missing free space line")
			break
		}
		line3 = scanner.Text()
		if line3 == "" {
			fmt.Println("Error: empty free space line")
			continue
		}

		fmt.Println("Parsed lines:", line1, line2, line3)

		// Parse device locator
		deviceLocators = append(deviceLocators, line1)

		// Parse used space
		var used, free, capacity float64
		used, err = parseFloat(line2)
		if err != nil {
			fmt.Println("Error parsing used space:", err)
			continue
		}
		usedSpaceBytes = append(usedSpaceBytes, used)

		// Parse free space
		free, err = parseFloat(line3)
		if err != nil {
			fmt.Println("Error parsing free space:", err)
			continue
		}
		freeSpaceBytes = append(freeSpaceBytes, free)

		// Total space is constant (last line)
		if !scanner.Scan() {
			fmt.Println("Error: missing total space line")
			break
		}
		line4 := scanner.Text()
		capacity, err = parseFloat(line4)
		if err != nil {
			fmt.Println("Error parsing total space:", err)
			continue
		}
		capacityBytes = append(capacityBytes, capacity)
	}

	return deviceLocators, usedSpaceBytes, freeSpaceBytes, capacityBytes
}

func parseFloat(value string) (float64, error) {
	var result float64
	_, err := fmt.Sscanf(value, "%f", &result)
	return result, err
}
