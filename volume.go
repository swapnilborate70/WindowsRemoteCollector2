package main

// Define volume struct
type Volume struct {
	SystemDiskVolume               string  `json:"system.disk.volume"`
	SystemDiskVolumeFreeBytes      int64   `json:"system.disk.volume.free.bytes"`
	SystemDiskVolumeUsedBytes      int64   `json:"system.disk.volume.used.bytes"`
	SystemDiskVolumeCapacityBytes  int64   `json:"system.disk.volume.capacity.bytes"`
	SystemDiskVolumeUsedPercentage float64 `json:"system.disk.volume.used.percent"`
	SystemDiskVolumeFreePercent    float64 `json:"system.disk.volume.free.percent"`
}
