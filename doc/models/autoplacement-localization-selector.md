
# Autoplacement Localization Selector

Request body to apply or clear cached autoplacement or auto-orientation values for a map or subset of APs

*This model accepts additional fields of type interface{}.*

## Structure

`AutoplacementLocalizationSelector`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `For` | [`*models.UseAutoApValuesForEnum`](../../doc/models/use-auto-ap-values-for-enum.md) | Optional | The selector to choose auto placement or auto orientation. enum: `orientation`, `placement`<br><br>**Default**: `"placement"` |
| `Macs` | `[]string` | Optional | List of AP MAC addresses to apply the action to. If omitted, the action applies to all APs on the map |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    autoplacementLocalizationSelector := models.AutoplacementLocalizationSelector{
        For:                  models.ToPointer(models.UseAutoApValuesForEnum_PLACEMENT),
        Macs:                 []string{
            "macs9",
            "macs8",
            "macs7",
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

