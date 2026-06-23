
# Use Auto Ap Values

Request to accept or reject cached AP placement or orientation values on a map

*This model accepts additional fields of type interface{}.*

## Structure

`UseAutoApValues`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Accept` | `*bool` | Optional | Whether to accept cached values; false rejects them for the selected APs<br><br>**Default**: `false` |
| `For` | [`*models.UseAutoApValuesForEnum`](../../doc/models/use-auto-ap-values-for-enum.md) | Optional | The selector to choose auto placement or auto orientation. enum: `orientation`, `placement`<br><br>**Default**: `"placement"` |
| `Macs` | `[]string` | Optional | AP MAC addresses to accept or reject. If omitted, the API applies the decision to the full map. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    useAutoApValues := models.UseAutoApValues{
        Accept:               models.ToPointer(false),
        For:                  models.ToPointer(models.UseAutoApValuesForEnum_PLACEMENT),
        Macs:                 []string{
            "macs5",
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

