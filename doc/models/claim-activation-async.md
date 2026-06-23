
# Claim Activation Async

Request to schedule an asynchronous inventory claim

*This model accepts additional fields of type interface{}.*

## Structure

`ClaimActivationAsync`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Code` | `string` | Required | Activation code to claim |
| `DeviceType` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Optional | enum: `ap`, `gateway`, `switch`<br><br>**Default**: `"ap"` |
| `Type` | [`models.ClaimTypeAsyncEnum`](../../doc/models/claim-type-async-enum.md) | Required | Claim scope for async inventory claiming. enum: `all`, `inventory` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    claimActivationAsync := models.ClaimActivationAsync{
        Code:                 "code8",
        DeviceType:           models.ToPointer(models.DeviceTypeDefaultApEnum_AP),
        Type:                 models.ClaimTypeAsyncEnum_ALL,
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

