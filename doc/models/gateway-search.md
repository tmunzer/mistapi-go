
# Gateway Search

Gateway record returned by device search endpoints

## Structure

`GatewaySearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Clustered` | `*bool` | Optional | Whether the gateway is part of a gateway cluster |
| `EvpnMissingLinks` | `*bool` | Optional | Whether EVPN topology links are missing for this gateway |
| `EvpntopoId` | `*string` | Optional | EVPN topology ID associated with this gateway |
| `ExtIp` | `*string` | Optional | External IP address observed for gateway management traffic |
| `Hostname` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Ip` | `*string` | Optional | Management IP address currently reported for the gateway |
| `LastConfigStatus` | `*string` | Optional | Most recent configuration status reported for the gateway |
| `LastHostname` | `*string` | Optional | Most recent hostname detected for the gateway |
| `LastTroubleCode` | `*string` | Optional | Most recent trouble code reported for the gateway |
| `LastTroubleTimestamp` | `*int` | Optional | Timestamp when the most recent gateway trouble code was reported |
| `Mac` | `*string` | Optional | Gateway MAC address reported in search results |
| `Managed` | `*bool` | Optional | Whether the gateway is managed by Mist. Deprecated in favor of `mist_configured` |
| `MistConfigured` | `*bool` | Optional | Whether the gateway can be configured by Mist. Replaces `managed` for adopted devices and `disable_auto_config` for claimed devices |
| `Model` | `*string` | Optional | Gateway model reported for this search result |
| `Node` | `*string` | Optional | Gateway cluster node associated with this search result |
| `Node0Mac` | `*string` | Optional | Cluster node0 MAC address reported for an HA gateway |
| `Node1Mac` | `*string` | Optional | Cluster node1 MAC address reported for an HA gateway |
| `NumMembers` | `*int` | Optional | Number of members in the gateway cluster |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Role` | `*string` | Optional | Gateway cluster role reported for this search result |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `T128agentVersion` | `*string` | Optional | Session Smart Router agent version reported by the gateway |
| `TimeDrifted` | `*bool` | Optional | Whether the gateway clock has drifted from the expected time |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Type` | `string` | Required, Constant | Device Type. enum: `gateway`<br><br>**Value**: `"gateway"` |
| `Uptime` | `*int` | Optional | Device uptime for the gateway, in seconds |
| `Version` | `*string` | Optional | Software version currently running on the gateway |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    gatewaySearch := models.GatewaySearch{
        Clustered:            models.ToPointer(false),
        EvpnMissingLinks:     models.ToPointer(false),
        EvpntopoId:           models.ToPointer("evpntopo_id2"),
        ExtIp:                models.ToPointer("ext_ip6"),
        Hostname:             []string{
            "hostname8",
        },
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Type:                 "gateway",
    }

}
```

