
# Module Stat Item Errors Items

Error condition reported for a device module

## Structure

`ModuleStatItemErrorsItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Feature` | `*string` | Optional | Affected feature associated with this module error condition |
| `MinimumVersion` | `*string` | Optional | Minimum software version associated with this module error |
| `Reason` | `*string` | Optional | Human-readable reason for the module error condition |
| `Since` | `int` | Required | Epoch timestamp when the module error condition began |
| `Type` | `string` | Required | Module error type code reported by the device |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    moduleStatItemErrorsItems := models.ModuleStatItemErrorsItems{
        Feature:              models.ToPointer("Mist-Management"),
        MinimumVersion:       models.ToPointer("128T-6.0.0-1"),
        Reason:               models.ToPointer("reason8"),
        Since:                1657497600,
        Type:                 "FW_UPGRADE_REQUIRED_BY_FEATURE",
    }

}
```

