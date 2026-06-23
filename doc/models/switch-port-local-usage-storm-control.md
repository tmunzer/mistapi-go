
# Switch Port Local Usage Storm Control

Storm-control settings for this local port configuration

## Structure

`SwitchPortLocalUsageStormControl`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisablePort` | `*bool` | Optional | Whether to disable the port when storm control is triggered<br><br>**Default**: `false` |
| `NoBroadcast` | `*bool` | Optional | Whether to disable storm control on broadcast traffic<br><br>**Default**: `false` |
| `NoMulticast` | `*bool` | Optional | Whether to disable storm control on multicast traffic<br><br>**Default**: `false` |
| `NoRegisteredMulticast` | `*bool` | Optional | Whether to disable storm control on registered multicast traffic<br><br>**Default**: `false` |
| `NoUnknownUnicast` | `*bool` | Optional | Whether to disable storm control on unknown unicast traffic<br><br>**Default**: `false` |
| `Percentage` | `*int` | Optional | Bandwidth-percentage, configures the storm control level as a percentage of the available bandwidth<br><br>**Default**: `80`<br><br>**Constraints**: `>= 0`, `<= 100` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchPortLocalUsageStormControl := models.SwitchPortLocalUsageStormControl{
        DisablePort:           models.ToPointer(false),
        NoBroadcast:           models.ToPointer(false),
        NoMulticast:           models.ToPointer(false),
        NoRegisteredMulticast: models.ToPointer(false),
        NoUnknownUnicast:      models.ToPointer(false),
        Percentage:            models.ToPointer(80),
    }

}
```

