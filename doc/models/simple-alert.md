
# Simple Alert

Heuristic alert thresholds used when a Marvis subscription is unavailable

## Structure

`SimpleAlert`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ArpFailure` | [`*models.SimpleAlertArpFailure`](../../doc/models/simple-alert-arp-failure.md) | Optional | Thresholds for ARP failure heuristic alerts |
| `DhcpFailure` | [`*models.SimpleAlertDhcpFailure`](../../doc/models/simple-alert-dhcp-failure.md) | Optional | Thresholds for DHCP failure heuristic alerts |
| `DnsFailure` | [`*models.SimpleAlertDnsFailure`](../../doc/models/simple-alert-dns-failure.md) | Optional | Thresholds for DNS failure heuristic alerts |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    simpleAlert := models.SimpleAlert{
        ArpFailure:           models.ToPointer(models.SimpleAlertArpFailure{
            ClientCount:          models.ToPointer(26),
            Duration:             models.ToPointer(60),
            IncidentCount:        models.ToPointer(226),
        }),
        DhcpFailure:          models.ToPointer(models.SimpleAlertDhcpFailure{
            ClientCount:          models.ToPointer(246),
            Duration:             models.ToPointer(60),
            IncidentCount:        models.ToPointer(6),
        }),
        DnsFailure:           models.ToPointer(models.SimpleAlertDnsFailure{
            ClientCount:          models.ToPointer(252),
            Duration:             models.ToPointer(60),
            IncidentCount:        models.ToPointer(0),
        }),
    }

}
```

