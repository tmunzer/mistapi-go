
# Utils Reset Radio Config

Request to reset AP radio settings to RRM-managed values

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsResetRadioConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bands` | `[]string` | Required | Radio bands to reset, such as `24`, `5`, or `6` |
| `Force` | `*bool` | Optional | Whether to reset radios that are currently disabled; default `false` honors radios intentionally disabled by the user<br><br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsResetRadioConfig := models.UtilsResetRadioConfig{
        Bands:                []string{
            "bands0",
            "bands1",
        },
        Force:                models.ToPointer(false),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

