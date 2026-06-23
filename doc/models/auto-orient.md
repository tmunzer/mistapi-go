
# Auto Orient

Request options for validating or starting AP auto-orientation

*This model accepts additional fields of type interface{}.*

## Structure

`AutoOrient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dryrun` | `*bool` | Optional | Set to `true` to perform an invalid AP check and provide an estimated run time without enqueuing the run into the auto orient service. |
| `ForceCollection` | `*bool` | Optional | If `force_collection`==`false`, the API attempts to start auto orientation with existing BLE data.<br>If `force_collection`==`true`, the API attempts to start BLE orchestration.<br><br>**Default**: `false` |
| `Macs` | `[]string` | Optional | List of device macs |
| `Override` | `*bool` | Optional | Set to `true` to run auto orient even if there are invalid APs in the selected APs. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    autoOrient := models.AutoOrient{
        Dryrun:               models.ToPointer(false),
        ForceCollection:      models.ToPointer(false),
        Macs:                 []string{
            "macs5",
            "macs6",
        },
        Override:             models.ToPointer(false),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

