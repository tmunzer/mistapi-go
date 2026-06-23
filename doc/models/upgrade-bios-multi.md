
# Upgrade Bios Multi

BIOS upgrade request for multiple devices

*This model accepts additional fields of type interface{}.*

## Structure

`UpgradeBiosMulti`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceIds` | `[]uuid.UUID` | Optional | Device IDs selected for a multi-device BIOS upgrade |
| `Models` | `[]string` | Optional | Device models selected for a multi-device BIOS upgrade |
| `Reboot` | `*bool` | Optional | Whether to restart the selected devices immediately after the upgrade completes<br><br>**Default**: `false` |
| `Version` | `*string` | Optional | BIOS version to install on the selected devices |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    upgradeBiosMulti := models.UpgradeBiosMulti{
        DeviceIds:            []uuid.UUID{
            uuid.MustParse("00001cc1-0000-0000-0000-000000000000"),
            uuid.MustParse("00001cc2-0000-0000-0000-000000000000"),
            uuid.MustParse("00001cc3-0000-0000-0000-000000000000"),
        },
        Models:               []string{
            "models8",
            "models9",
            "models0",
        },
        Reboot:               models.ToPointer(false),
        Version:              models.ToPointer("CDEN_P_EX1_00.15.01.00"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

