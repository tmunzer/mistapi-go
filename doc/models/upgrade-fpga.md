
# Upgrade Fpga

FPGA upgrade request for a single device

*This model accepts additional fields of type interface{}.*

## Structure

`UpgradeFpga`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Reboot` | `*bool` | Optional | Whether to restart the device immediately after the upgrade completes<br><br>**Default**: `false` |
| `Version` | `*string` | Optional | FPGA version to install on the device |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    upgradeFpga := models.UpgradeFpga{
        Reboot:               models.ToPointer(false),
        Version:              models.ToPointer("REV37"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

