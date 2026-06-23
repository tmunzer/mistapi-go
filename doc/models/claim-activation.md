
# Claim Activation

Request to claim organization licenses or activation codes

*This model accepts additional fields of type interface{}.*

## Structure

`ClaimActivation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Code` | `string` | Required | Activation or license claim code to redeem |
| `DeviceType` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Optional | enum: `ap`, `gateway`, `switch`<br><br>**Default**: `"ap"` |
| `Type` | [`models.ClaimTypeEnum`](../../doc/models/claim-type-enum.md) | Required | what to claim. enum: `all`, `inventory`, `license`<br><br>**Default**: `"all"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    claimActivation := models.ClaimActivation{
        Code:                 "code4",
        DeviceType:           models.ToPointer(models.DeviceTypeDefaultApEnum_AP),
        Type:                 models.ClaimTypeEnum_ALL,
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

