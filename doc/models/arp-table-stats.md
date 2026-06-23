
# Arp Table Stats

ARP table usage and capacity statistics

## Structure

`ArpTableStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ArpTableCount` | `*int` | Optional | Number of ARP table entries currently present on the device |
| `MaxEntriesSupported` | `*int` | Optional | Supported ARP table capacity for the device |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    arpTableStats := models.ArpTableStats{
        ArpTableCount:        models.ToPointer(76),
        MaxEntriesSupported:  models.ToPointer(68),
    }

}
```

