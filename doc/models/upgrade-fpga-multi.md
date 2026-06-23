
# Upgrade Fpga Multi

FPGA upgrade request for multiple devices

*This model accepts additional fields of type interface{}.*

## Structure

`UpgradeFpgaMulti`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceIds` | `[]uuid.UUID` | Optional | Device IDs selected for a multi-device FPGA upgrade |
| `Models` | `[]string` | Optional | Device models selected for a multi-device FPGA upgrade |
| `Reboot` | `*bool` | Optional | Whether to restart the selected devices immediately after the upgrade completes<br><br>**Default**: `false` |
| `Version` | `*string` | Optional | FPGA version to install on the selected devices |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    upgradeFpgaMulti := models.UpgradeFpgaMulti{
        DeviceIds:            []uuid.UUID{
            uuid.MustParse("0000094f-0000-0000-0000-000000000000"),
        },
        Models:               []string{
            "models0",
        },
        Reboot:               models.ToPointer(false),
        Version:              models.ToPointer("REV37"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

