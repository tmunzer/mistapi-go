
# Switch Port Config Overwrite

Switch port configuration overrides

## Structure

`SwitchPortConfigOverwrite`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Description` | `*string` | Optional | Administrative description applied to the switch port override |
| `Disabled` | `*bool` | Optional | Whether the port is disabled<br><br>**Default**: `false` |
| `Duplex` | [`*models.SwitchPortUsageDuplexOverwriteEnum`](../../doc/models/switch-port-usage-duplex-overwrite-enum.md) | Optional | Link connection mode. enum: `auto`, `full`, `half`<br><br>**Default**: `"auto"` |
| `MacLimit` | [`*models.SwitchPortUsageMacLimitOverwrite`](../../doc/models/containers/switch-port-usage-mac-limit-overwrite.md) | Optional | Max number of MAC addresses, default is 0 for unlimited, otherwise range is 1 to 16383 (upper bound constrained by platform) |
| `PoeDisabled` | `*bool` | Optional | Whether PoE capabilities are disabled for a port<br><br>**Default**: `false` |
| `PoeKeepStateWhenReboot` | `*bool` | Optional | Whether Perpetual PoE is enabled; keeps PoE state across reboots<br><br>**Default**: `false` |
| `PortNetwork` | `*string` | Optional | Native network/vlan for untagged traffic |
| `Speed` | [`*models.SwitchPortUsageSpeedOverwriteEnum`](../../doc/models/switch-port-usage-speed-overwrite-enum.md) | Optional | Port Speed, default is auto to automatically negotiate speed enum: `100m`, `10m`, `1g`, `2.5g`, `5g`, `10g`, `25g`, `40g`, `100g`,`auto`<br><br>**Default**: `"auto"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchPortConfigOverwrite := models.SwitchPortConfigOverwrite{
        Description:            models.ToPointer("description4"),
        Disabled:               models.ToPointer(false),
        Duplex:                 models.ToPointer(models.SwitchPortUsageDuplexOverwriteEnum_AUTO),
        MacLimit:               models.ToPointer(models.SwitchPortUsageMacLimitOverwriteContainer.FromNumber(96)),
        PoeDisabled:            models.ToPointer(false),
        PoeKeepStateWhenReboot: models.ToPointer(false),
        Speed:                  models.ToPointer(models.SwitchPortUsageSpeedOverwriteEnum_AUTO),
    }

}
```

