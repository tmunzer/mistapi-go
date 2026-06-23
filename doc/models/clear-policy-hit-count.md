
# Clear Policy Hit Count

Request body for clearing hit counters on an application policy

*This model accepts additional fields of type interface{}.*

## Structure

`ClearPolicyHitCount`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PolicyName` | `string` | Required | Application policy name whose hit counters should be cleared |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    clearPolicyHitCount := models.ClearPolicyHitCount{
        PolicyName:           "policy_name4",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

