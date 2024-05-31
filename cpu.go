package main

import (
	"bytes"
	"encoding/csv"
)

type CPUInfo struct {
	SystemCPUCore                 string  `json:"system.cpu.core"`
	SystemCPUCoreUserPercent      float64 `json:"system.cpu.core.user.percent"`
	SystemCPUCorePercent          float64 `json:"system.cpu.core.percent"`
	SystemCPUCoreInterruptPercent float64 `json:"system.cpu.core.interrupt.percent"`
	SystemCPUCoreIdlePercent      float64 `json:"system.cpu.core.idle.percent"`
}

func parseCSV(csvData []byte) ([][]string, error) {
	// Remove leading/trailing whitespace and any empty lines
	csvReader := csv.NewReader(bytes.NewReader(bytes.TrimSpace(csvData)))
	csvReader.FieldsPerRecord = -1 // Allow variable number of fields

	// Read CSV records until EOF
	var records [][]string
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break // Reached end of file
			}
			return nil, err // Return if any other error occurs
		}
		records = append(records, record)
	}

	return records, nil

}
