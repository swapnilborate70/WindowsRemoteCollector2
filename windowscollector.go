package main

type Collector struct {
	Memory    []Memory    `json:"memory"`
	Interface []Interface `json:"interface"`
	CPU       []CPUInfo   `json:"cpu"`
	Volume    []Volume    `json:"volume"`
}
