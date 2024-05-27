$cpuMetrics = @{
    "cpu" = @{
        "system.cpu.core" = (Get-WmiObject -Class Win32_Processor).Name
        "system.cpu.core.user.percent" = (Get-Counter '\Processor(_Total)\% User Time').CounterSamples.CookedValue
        "system.cpu.core.percent" = (Get-Counter '\Processor(_Total)\% Processor Time').CounterSamples.CookedValue
        "system.cpu.core.interrupt.percent" = (Get-Counter '\Processor(_Total)\% Interrupt Time').CounterSamples.CookedValue
        "system.cpu.core.idle.percent" = (Get-Counter '\Processor(_Total)\% Idle Time').CounterSamples.CookedValue
    }
}

# Get Disk metrics
$diskMetrics = @{
    "disk" = @{
        "system.disk.read.bytes.per.sec" = (Get-Counter '\LogicalDisk(*)\Disk Read Bytes/sec').CounterSamples.CookedValue
        "system.disk.write.bytes.per.sec" = (Get-Counter '\LogicalDisk(*)\Disk Write Bytes/sec').CounterSamples.CookedValue
        "system.disk.bytes.per.sec" = (Get-Counter '\LogicalDisk(*)\Disk Bytes/sec').CounterSamples.CookedValue
        "system.disk.write.ops.per.sec" = (Get-Counter '\LogicalDisk(*)\Disk Write Operations/sec').CounterSamples.CookedValue
        # Add more disk metrics as needed
    }
}

# Get Interface metrics
$interfaceMetrics = @{
    "interface" = @{}
}

# Get network adapter statistics
$networkAdapters = Get-NetAdapterStatistics | Select-Object -Property Name, ReceivedBytes, SentBytes, ReceivedPackets, SentPackets, ReceivedUnicastPackets, SentUnicastPackets

# Add metrics for each network adapter
foreach ($adapter in $networkAdapters) {
    $interfaceMetrics["interface"][$adapter.Name] = @{
        "system.network.interface.in.bytes.per.sec" = $adapter.ReceivedBytes
        "system.network.interface.out.bytes.per.sec" = $adapter.SentBytes
        "system.network.interface.in.packets.per.sec" = $adapter.ReceivedPackets
        "system.network.interface.out.packets.per.sec" = $adapter.SentPackets
        "system.network.interface.in.unicast.packets.per.sec" = $adapter.ReceivedUnicastPackets
        "system.network.interface.out.unicast.packets.per.sec" = $adapter.SentUnicastPackets
        # Add more interface metrics as needed
    }
}



# Get Memory metrics
$memoryMetrics = @{
    "memory" = @{
        "system.memory.total.bytes" = (Get-WmiObject -Class Win32_OperatingSystem).TotalVisibleMemorySize * 1KB
        "system.memory.used.bytes" = (Get-WmiObject -Class Win32_OperatingSystem).FreePhysicalMemory * 1KB
        "system.memory.free.bytes" = (Get-WmiObject -Class Win32_OperatingSystem).FreePhysicalMemory * 1KB
        "system.memory.used.percent" = (1 - (Get-WmiObject -Class Win32_OperatingSystem).FreePhysicalMemory / (Get-WmiObject -Class Win32_OperatingSystem).TotalVisibleMemorySize) * 100
        "system.memory.free.percent" = (Get-WmiObject -Class Win32_OperatingSystem).FreePhysicalMemory / (Get-WmiObject -Class Win32_OperatingSystem).TotalVisibleMemorySize * 100
    }
}

# Combine all metrics
$allMetrics = @{
    "cpu" = $cpuMetrics["cpu"]
    "disk" = $diskMetrics["disk"]
    "interface" = $interfaceMetrics["interface"]
    "memory" = $memoryMetrics["memory"]
}

# Convert to JSON
$jsonOutput = $allMetrics | ConvertTo-Json

# Output JSON
Write-Output $jsonOutput
