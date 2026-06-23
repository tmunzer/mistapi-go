
# Marvis Auto Operations

Marvis automatic remediation operation toggles

## Structure

`MarvisAutoOperations`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApInsufficientCapacity` | `*bool` | Optional | Whether Marvis may remediate AP insufficient-capacity issues automatically<br><br>**Default**: `false` |
| `ApLoop` | `*bool` | Optional | Whether Marvis may remediate AP loop issues automatically<br><br>**Default**: `false` |
| `ApNonCompliant` | `*bool` | Optional | Whether Marvis may remediate AP non-compliance automatically<br><br>**Default**: `false` |
| `BouncePortForAbnormalPoeClient` | `*bool` | Optional | Whether Marvis may bounce switch ports for abnormal PoE clients<br><br>**Default**: `false` |
| `DisablePortWhenDdosProtocolViolation` | `*bool` | Optional | Whether Marvis may disable a port when DDOS protocol violations are detected<br><br>**Default**: `false` |
| `DisablePortWhenRogueDhcpServerDetected` | `*bool` | Optional | Whether Marvis may disable a port when a rogue DHCP server is detected<br><br>**Default**: `false` |
| `GatewayNonCompliant` | `*bool` | Optional | Whether Marvis may remediate non-compliant gateways automatically<br><br>**Default**: `false` |
| `SwitchMisconfiguredPort` | `*bool` | Optional | Whether Marvis may remediate misconfigured switch ports automatically<br><br>**Default**: `false` |
| `SwitchPortStuck` | `*bool` | Optional | Whether Marvis may remediate stuck switch ports automatically<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    marvisAutoOperations := models.MarvisAutoOperations{
        ApInsufficientCapacity:                 models.ToPointer(false),
        ApLoop:                                 models.ToPointer(false),
        ApNonCompliant:                         models.ToPointer(false),
        BouncePortForAbnormalPoeClient:         models.ToPointer(false),
        DisablePortWhenDdosProtocolViolation:   models.ToPointer(false),
        DisablePortWhenRogueDhcpServerDetected: models.ToPointer(false),
        GatewayNonCompliant:                    models.ToPointer(false),
        SwitchMisconfiguredPort:                models.ToPointer(false),
        SwitchPortStuck:                        models.ToPointer(false),
    }

}
```

